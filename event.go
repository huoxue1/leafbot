package leafBot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huoxue1/leafBot/message" //nolint:gci
	log "github.com/sirupsen/logrus"
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	c = make(chan Event, 10)
)

var ENABLE = true // 是否启用gui

func init() {
	if runtime.GOOS != "windows" {
		ENABLE = false
	}
}

var (
	MessageChan = make(chan Event, 10)
	//NoticeChan  = make(chan Event, 10)
	//Request     = make(chan Event, 10)
)

var (
	ConnectHandles      ConnectChain
	DisConnectHandles   DisConnectChain
	MessageHandles      MessageChain
	RequestHandles      RequestChain
	NoticeHandles       NoticeChain
	CommandHandles      CommandChain
	MetaHandles         MetaChain
	PretreatmentHandles PretreatmentChain

	sessions sync.Map
)

// sessionHandle
/*
   @Description: 遍历所有的session发现是否需要将消息传到对应的handle里面
   @param event Event
*/
func sessionHandle(event Event) bool {
	DISUSE := false
	sessions.Range(func(key, value interface{}) bool {
		s := value.(session)
		rule := checkRule(event, s.rules, &State{})
		if s.rules == nil {
			rule = true
		}
		if rule {
			s.queue <- event
			DISUSE = true
		}
		return true
	})

	return DISUSE
}

type (
	session struct {
		id    int
		queue chan Event
		rules []Rule
	}
)

// eventMain
/*
   @Description: 事件总处理器，所有的事件都从这里开始处理
*/
func eventMain() {
	sort.Sort(&MessageHandles)
	sort.Sort(&RequestHandles)
	sort.Sort(&NoticeHandles)
	sort.Sort(&CommandHandles)

	if len(DefaultConfig.CommandStart) == 0 {
		DefaultConfig.CommandStart = append(DefaultConfig.CommandStart, "")
	}

	for _, handle := range PretreatmentHandles {
		log.Infoln("已加载预处理器响应器：" + getFunctionName(handle.handle, '/'))
	}
	for _, handle := range CommandHandles {
		if handle.command == "" && handle.regexMatcher != "" {
			log.Infoln("已加载regex响应器：" + handle.Name)
		} else {
			log.Infoln("已加载command响应器：" + handle.command)
		}
	}
	for _, handle := range MessageHandles {
		log.Infoln("已加载message响应器：" + getFunctionName(handle.handle, '/'))
	}
	for _, handle := range RequestHandles {
		log.Infoln("已加载request响应器：" + getFunctionName(handle.handle, '/'))
	}
	for _, handle := range NoticeHandles {
		log.Infoln("已加载notice响应器：" + getFunctionName(handle.handle, '/'))
	}
	for _, handle := range MetaHandles {
		log.Infoln("已加载meta响应器：" + getFunctionName(handle.handle, '/'))
	}

	go func() {
		for {
			data, ok := <-eventChan
			if !ok {
				continue
			}
			log.Debugln(string(data))
			var event Event
			err := json.Unmarshal(data, &event)
			if err != nil {
				log.Debugln("反向解析json失败" + err.Error() + "\n" + string(data))
			}
			go viewsMessage(event)
		}
	}()
}

// GetOneEvent
/*
   @Description: 向session队列里面添加一个对象，等待用户的响应，设置超时时间
   @receiver b
   @param rules ...Rule
   @return Event  Event
   @return error  error
*/
func (b *Bot) GetOneEvent(rules ...Rule) (Event, error) {
	s := session{
		id:    int(time.Now().Unix() + rand.Int63n(10000)),
		queue: make(chan Event),
		rules: rules,
	}
	sessions.Store(s.id, s)
	defer sessions.Delete(s.id)
	select {
	case event := <-s.queue:
		return event, nil
	case <-time.After(time.Minute):
		return Event{}, errors.New("等待下一条信息超时")
	}

}

// GetMoreEvent
/*
   @Description: 获取一个通道不断从用户获取消息
   @receiver b
   @param rules ...Rule
   @return int  int 对应session在队列中的编号，后面关闭需要该编号
   @return chan  Event  事件通道
*/
func (b *Bot) GetMoreEvent(rules ...Rule) (int, chan Event) {
	s := session{
		id:    int(time.Now().Unix() + rand.Int63n(10000)),
		queue: make(chan Event),
		rules: rules,
	}
	sessions.Store(s.id, s)
	return s.id, s.queue
}

// CloseMessageChan
/*
   @Description: 关闭session，即从等待队列中删除
   @receiver b
   @param id int
*/
func (b *Bot) CloseMessageChan(id int) {
	sessions.Delete(id)
}

// viewsMessage
/*
   @Description: 对所有event进行分类，按照不同的type进行不同的处理
   @param event Event
*/
func viewsMessage(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("消息分类出现不可挽回的错误")
			log.Errorln(err)
		}
	}()

	// 执行所有预处理handle
	for _, handle := range PretreatmentHandles {
		bot := GetBotById(event.SelfId)
		rule := checkRule(event, handle.rules, &State{})
		// 执行rule判断
		if !rule || !handle.Enable {
			continue
		}
		// 判断所在群是否禁用
		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}
		// 执行handle
		b := handle.handle(event, bot)
		// 判断是否被block
		if !b {
			return
		}
	}
	log.Debugln("预处理执行完毕" + event.Message.CQString())
	switch event.PostType {
	case "message":
		c <- event
		if ENABLE {
			MessageChan <- event
		}
		processMessageHandle()

	case "notice":
		processNoticeHandle(event)
	case "request":
		log.Infoln(fmt.Sprintf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType, event.GroupId, event.UserId))
		processRequestEventHandle(event)

	case "meta_event":
		log.Debugln(fmt.Sprintf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%d",
			event.PostType, event.MetaEventType, event.Interval))
		processMetaEventHandle(event)
	}
}

// processNoticeHandle
/*
   @Description: Notice事件的handle处理
   @param event Event
*/
func processNoticeHandle(event Event) {
	defer func() {
		err := recover()
		log.Errorln("notice事件处理器出现错误")
		log.Errorln(err)
	}()
	log.Infoln(fmt.Sprintf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
		event.NoticeType, event.GroupId, event.UserId))

	for _, v := range NoticeHandles {
		rule := checkRule(event, v.rules, &State{})
		if !rule || !v.Enable {
			continue
		}
		if v.noticeType == "" {
			v.noticeType = event.NoticeType
		}
		if v.noticeType == event.NoticeType {
			go func(handle2 *noticeHandle) {
				defer func() {
					err := recover()
					if err != nil {
						log.Errorln(handle2.Name + "发生不可挽回的错误")
						log.Errorln(err)
					}
				}()
				handle2.handle(event, GetBotById(event.SelfId))
			}(v)
		}
	}
}

// checkRule
/**
 * @Description:
 * @param event
 * @param rules
 * @param state
 * @return bool
 * example
 */
func checkRule(event Event, rules []Rule, state *State) bool {
	for _, rule := range rules {
		check := rule(event, GetBotById(event.SelfId), state)
		if !check {
			return false
		}
	}
	return true
}

// getFunctionName
/**
 * @Description:
 * @param i
 * @param seps
 * @return string
 * example
 */
func getFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func checkCD(handle *commandHandle) bool {
	if handle.cd.Types == "" || handle.cd.Types == "default" {
		if int(time.Now().Unix()-handle.lastUseTime) >= handle.cd.Long {
			return true
		}
	} else if handle.cd.Types == "rand" {
		rand.Seed(time.Now().UnixNano())
		cd := rand.Intn(handle.cd.Long)
		if int(time.Now().Unix()-handle.lastUseTime) >= cd {
			return true
		}
	}
	return false
}

func doHandle(handle *commandHandle, event Event, state *State) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln(handle.Name + "发生不可挽回的错误")
			log.Errorln(err)
		}
	}()
	handle.handle(event, GetBotById(event.SelfId), state)
}

func checkOnlyTome(event *Event, state *State) {
	if event.Message[0].Type == "at" && event.Message[0].Data["qq"] == strconv.Itoa(event.SelfId) {
		event.Message = event.Message[1:]
		state.Data["only_tome"] = true
	}
	for _, segment := range event.Message {
		if segment.Type == "at" && segment.Data["qq"] == strconv.Itoa(event.SelfId) {
			state.Data["only_tome"] = true
		}
	}
	if event.MessageType == "private" {
		state.Data["only_tome"] = true
	}
}

/**
 * @Description: 处理message的响应器
 * example
 */
func processMessageHandle() {

	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("message事件处理器出现不可挽回的错误")
			log.Errorln(err)
		}
	}()
	// 从队列中取出事件
	event := <-c
	// 判断是否触发命令的flag
	a := 0
	log.Debugln(len(CommandHandles))
	eventData, _ := json.Marshal(event)
	// 执行连续会话的handle，如果返回true说明该消息被连续对话捕捉
	if sessionHandle(event) {
		return
	}

	state := new(State)
	state.Data = make(map[string]interface{})
	checkOnlyTome(&event, state)
	// 遍历所有的command对象
	for _, handle := range CommandHandles {

		// 判断该cmd在该群是否被禁用
		disable := true
		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				disable = false
			}
		}
		if !disable {
			continue
		}

		// 检查rules
		rule := checkRule(event, handle.rules, state)
		if handle.rules == nil {
			rule = true
		}
		// 判断rules和插件是否被禁用
		if !rule || !handle.Enable {
			continue
		}
		//if event.Message[0].Type != "text" {
		//	continue
		//}

		commands := strings.Split(event.Message[0].Data["text"], " ")
		if len(commands) < 1 {
			continue
		}
		for _, start := range DefaultConfig.CommandStart {
			if commands[0] == start+handle.command && handle.command != "" {

				// 检查cd是否达到
				if !checkCD(handle) {
					continue
				}
				a = 1
				handle.lastUseTime = time.Now().Unix()

				state.Args = commands[1:]
				state.Cmd = handle.command
				state.Allies = handle.allies

				doHandle(handle, event, state)

				log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, eventData, handle.command))
				if handle.block {
					return
				}
			}
		}

		// 处理别名匹配
		for _, ally := range handle.allies {
			if ally == commands[0] {

				// 检查cd是否达到
				if !checkCD(handle) {
					continue
				}
				a = 1
				handle.lastUseTime = time.Now().Unix()

				state.Args = commands[1:]
				state.Cmd = handle.command
				state.Allies = handle.allies

				doHandle(handle, event, state)
				log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, eventData, handle.command))
				if handle.block {
					return
				}
			}
		}

		// 处理正则匹配
		if handle.command == "" && handle.regexMatcher != "" {
			compile := regexp.MustCompile(handle.regexMatcher)
			if compile.MatchString(event.Message.CQString()) {

				state.Args = commands[1:]
				state.Cmd = handle.regexMatcher
				state.Allies = handle.allies
				state.RegexResult = compile.FindStringSubmatch(event.Message.CQString())

				doHandle(handle, event, state)
				log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, eventData, handle.regexMatcher))
				if handle.block {
					return
				}
			}
		}
	}

	// 如果出发了command事件则不再触发message事件
	if a == 1 {
		return
	}
	s := new(State)
	s.Data = make(map[string]interface{})
	checkOnlyTome(&event, s)
	log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType, event.GroupId, event.UserId, eventData))
	for _, handle := range MessageHandles {
		if handle.messageType != "" && handle.messageType != event.MessageType {
			continue
		}

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules, s)
		if !rule || !handle.Enable {
			continue
		}
		go func(handle2 *messageHandle) {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(handle2.Name + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			handle2.handle(event, GetBotById(event.SelfId), s)
		}(handle)
	}
}

// processRequestEventHandle
/*
   @Description: Request类型event处理
   @param event Event
*/
func processRequestEventHandle(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln(err)
			log.Errorln("request事件处理器出现不可挽回的错误")
		}
	}()
	for _, handle := range RequestHandles {

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules, &State{})
		if handle.rules == nil {
			rule = true
		}
		if !rule || !handle.Enable {
			continue
		}
		go func(handle2 *requestHandle) {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(handle2.Name + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			handle2.handle(event, GetBotById(event.SelfId))
		}(handle)
	}
}

// processMetaEventHandle
/**
 * @Description:
 * @param event
 * example
 */
func processMetaEventHandle(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("元事件处理器错误")
			log.Errorln(err)
		}
	}()
	for _, handle := range MetaHandles {

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules, &State{})
		if handle.rules == nil {
			rule = true
		}
		if !rule || !handle.Enable {
			continue
		}
		go func(handle2 *metaHandle) {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(handle2.Name + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			handle2.handle(event, GetBotById(event.SelfId))
		}(handle)

	}
}

// GetBotById
/*
   @Description:
   @param id int
   @return *Bot
*/
func GetBotById(id int) *Bot {
	for _, bot := range DefaultConfig.Bots {
		if bot.SelfId == id {
			return bot
		}
	}
	return nil
}

// GetMsg
/**
 * @Description:
 * @receiver e
 * @return message.Message
 */
func (e Event) GetMsg() message.Message {
	return e.Message
}

func (e Event) GetPlainText() string {
	content := ""
	for _, mes := range e.Message {
		if mes.Type == "text" {
			content += mes.Data["text"]
		}
	}
	return content
}

func (e Event) GetImages() []message.MessageSegment {
	var images []message.MessageSegment
	for _, mes := range e.Message {
		if mes.Type == "image" {
			images = append(images, mes)
		}
	}
	return images
}
