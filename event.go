// Package leafbot
// @Description:
package leafbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	//nolint:gci

	"github.com/huoxue1/leafbot/message"
)

//var ENABLE = false // 是否启用gui

var (
// MessageChan = make(chan Event, 10)
// NoticeChan  = make(chan Event, 10)
// Request     = make(chan Event, 10)
)

var (
	sessions sync.Map
)

// sessionHandle
/*
   @Description: 遍历所有的session发现是否需要将消息传到对应的handle里面
   @param event Event
*/
func sessionHandle(ctx *Context) bool {
	DISUSE := false
	sessions.Range(func(key, value interface{}) bool {
		s := value.(session)
		rule := checkRule(s.rules, ctx)
		if s.rules == nil {
			rule = true
		}
		if rule {
			s.queue <- *ctx.Event
			DISUSE = true
		}
		return true
	})

	return DISUSE
}

type (
	session struct {
		id    string
		queue chan Event
		rules []Rule
	}
)

// eventMain
/*
   @Description: 事件总处理器，所有的事件都从这里开始处理
*/
func eventMain() {
	sort.Sort(matcherChin)

	if len(defaultConfig.CommandStart) == 0 {
		defaultConfig.CommandStart = append(defaultConfig.CommandStart, "")
	}

	for _, plugin := range plugins {
		log.Infoln("已加载插件 ==》 " + plugin.Name)
	}

	e := driver.GetEvent()

	go func() {
		for {
			data, ok := <-e
			if !ok {
				continue
			}
			log.Debugln(string(data))
			var event Event
			err := json.Unmarshal(data, &event)
			if err != nil {
				log.Debugln("反向解析json失败" + err.Error() + "\n" + string(data))
				if gjson.GetBytes(data, "post_type").String() == "message" && !gjson.GetBytes(data, "message").IsArray() {
					log.Errorln("检测到onebot端采用了string上报，建议更改未array上报")
				}
			}
			go viewsMessage(event)
		}
	}()
}

// GetQuestion
/**
 * @Description: 向当前用户发送一个问题，并获取答案
 * @receiver ctx
 * @param question
 * @return string
 * @return error
 */
func (ctx *Context) GetQuestion(question string) (string, error) {
	ctx.Send(message.Text(question))
	event, err := ctx.GetOneEvent(func(ctx1 *Context) bool {
		if ctx1.Event.GroupId == ctx.Event.GroupId && ctx1.Event.UserId == ctx.Event.UserId {
			return true
		}
		return false
	}, func(ctx *Context) bool {
		if ctx.Event.MessageType != "message" {
			return false
		}
		return true
	})
	return event.RawMessage, err
}

// GetOneEvent
/*
  @Description: 向session队列里面添加一个对象，等待用户的响应，设置超时时间
  @receiver b
  @param rules ...Rule
  @return Event  Event
  @return error  error
*/
func (ctx *Context) GetOneEvent(rules ...Rule) (Event, error) {
	s := session{
		id:    uuid.NewV4().String(),
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
func (ctx *Context) GetMoreEvent(rules ...Rule) (chan Event, func()) {
	s := session{
		id:    uuid.NewV4().String(),
		queue: make(chan Event),
		rules: rules,
	}
	sessions.Store(s.id, s)
	return s.queue, func() {
		ctx.closeMessageChan(s.id)
		close(s.queue)
	}
}

// CloseMessageChan
/*
  @Description: 关闭session，即从等待队列中删除
  @receiver b
  @param id int
*/
func (ctx *Context) closeMessageChan(id string) {
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
	bot := GetBotById(event.SelfId)
	state := new(State)
	ctx := new(Context)
	ctx.Event = &event
	ctx.Bot = bot
	ctx.State = state
	ctx.SelfID = int64(event.SelfId)
	ctx.UserID = event.UserId
	ctx.GroupID = event.GroupId
	data1, _ := json.Marshal(&event)
	log.Debugln(string(data1))
	//// 执行所有预处理handle
	//for _, handle := range PretreatmentHandles {
	//	rule := checkRule(handle.rules, ctx)
	//	// 执行rule判断
	//	if !rule || !handle.Enable {
	//		continue
	//	}
	//	// 判断所在群是否禁用
	//	for _, group := range handle.disableGroup {
	//		if event.GroupId == group {
	//			return
	//		}
	//	}
	//	// 执行handle
	//	b := handle.handle(ctx)
	//	// 判断是否被block
	//	if !b {
	//		return
	//	}
	//}
	log.Debugln("预处理执行完毕")
	switch event.PostType {
	case "message":
		// if ENABLE {
		//	MessageChan <- event
		// }
		processMessageHandle(ctx)

	case "notice":
		processNoticeHandle(ctx)
	case "request":
		log.Infoln(fmt.Sprintf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType, event.GroupId, event.UserId))
		processRequestEventHandle(ctx)

	case "meta_event":
		log.Debugln(fmt.Sprintf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%d",
			event.PostType, event.MetaEventType, event.Interval))
		processMetaEventHandle(ctx)
	}
}

// processNoticeHandle
/*
   @Description: Notice事件的handle处理
   @param event Event
*/
func processNoticeHandle(ctx *Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("notice事件处理器出现错误")
			log.Errorln(err)
		}
	}()
	log.Infoln(fmt.Sprintf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
		ctx.Event.NoticeType, ctx.Event.GroupId, ctx.Event.UserId))

	matcherChin.forEach(NOTICE, func(matcher Matcher) bool {
		rule := checkRule(matcher.GetRules(), ctx)
		if !rule {
			return true
		}
		if matcher.GetType() == ctx.Event.NoticeType || matcher.GetType() == "" {
			go func() {
				defer func() {
					err := recover()
					if err != nil {
						log.Errorln(getFunctionName(matcher.GetHandler(), '/') + "发生了不可挽回的错误！")
						log.Errorln(err)
					}
				}()
				matcher.GetHandler()(ctx)
			}()
		}
		if matcher.IsBlock() {
			return false
		}
		return true
	})
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
func checkRule(rules []Rule, ctx *Context) bool {
	for _, rule := range rules {
		check := rule(ctx)
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

//func checkCD(handle *commandHandle) bool {
//	if handle.cd.Types == "" || handle.cd.Types == "default" {
//		if int(time.Now().Unix()-handle.lastUseTime) >= handle.cd.Long {
//			return true
//		}
//	} else if handle.cd.Types == "rand" {
//		rand.Seed(time.Now().UnixNano())
//		cd := rand.Intn(handle.cd.Long)
//		if int(time.Now().Unix()-handle.lastUseTime) >= cd {
//			return true
//		}
//	}
//	return false
//}

func doHandle(handle Matcher, ctx *Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln(getFunctionName(handle.GetHandler(), '/') + "发生不可挽回的错误")
			log.Errorln(err)
		}
	}()
	handle.GetHandler()(ctx)
}

func checkOnlyTome(event *Event, state *State) {
	if event.Message[0].Type == "at" && event.Message[0].Data["qq"] == strconv.FormatInt(event.SelfId, 10) {
		event.Message = event.Message[1:]
		state.Data["only_tome"] = true
		return
	}
	for _, segment := range event.Message {
		if segment.Type == "at" && segment.Data["qq"] == strconv.FormatInt(event.SelfId, 10) {
			state.Data["only_tome"] = true
			return
		}
	}
	for _, name := range defaultConfig.NickName {
		if event.Message[0].Type == "text" && strings.HasPrefix(event.Message[0].Data["text"], name) {
			state.Data["only_tome"] = true
			text := strings.TrimLeft(event.Message[0].Data["text"], name)
			event.Message[0].Data["text"] = text
			return
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
func processMessageHandle(ctx *Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("message事件处理器出现不可挽回的错误")
			log.Errorln(err)
		}
	}()
	// 从队列中取出事件
	// 判断是否触发命令的flag
	a := 0
	// 执行连续会话的handle，如果返回true说明该消息被连续对话捕捉
	if sessionHandle(ctx) {
		return
	}

	ctx.State.Data = make(map[string]interface{})
	checkOnlyTome(ctx.Event, ctx.State)
	// 遍历所有的command对象
	matcherChin.forEach(COMMAND, func(matcher Matcher) bool {
		// 判断该cmd在该群是否被禁用
		disable := true
		for _, group := range matcher.GetDisAbleGroup() {
			if ctx.Event.GroupId == group {
				disable = false
			}
		}
		if !disable {
			return true
		}
		// 检查rules
		rule := checkRule(matcher.GetRules(), ctx)

		if matcher.GetRules() == nil {
			rule = true
		}

		// 判断rules和插件是否被禁用
		if !rule {
			return true
		}
		// if ctx.Event.Message[0].Type != "text" {
		//	continue
		//}

		commands := strings.Split(ctx.Event.GetPlainText(), " ")
		if len(commands) < 1 {
			return true
		}

		for _, start := range defaultConfig.CommandStart {
			if commands[0] == start+matcher.(CommandMatcher).GetCommand() && matcher.(CommandMatcher).GetCommand() != "" {
				// 检查cd是否达到
				//if !checkCD(handle) {
				//	continue
				//}
				a = 1
				//handle.lastUseTime = time.Now().Unix()

				ctx.State.Args = commands[1:]
				ctx.State.Cmd = matcher.(CommandMatcher).GetCommand()
				ctx.State.Allies = matcher.(CommandMatcher).GetAlias()

				doHandle(matcher, ctx)

				log.Infoln(fmt.Sprintf("触发了：%v", matcher.(CommandMatcher).GetCommand()))
				if matcher.IsBlock() {
					return false
				}
			}
		}

		// 处理别名匹配
		for _, ally := range matcher.(CommandMatcher).GetAlias() {
			if ally == commands[0] {
				// 检查cd是否达到
				//if !checkCD(handle) {
				//	continue
				//}
				a = 1
				//handle.lastUseTime = time.Now().Unix()

				ctx.State.Args = commands[1:]
				ctx.State.Cmd = matcher.(CommandMatcher).GetCommand()
				ctx.State.Allies = matcher.(CommandMatcher).GetAlias()

				doHandle(matcher, ctx)
				log.Infoln(fmt.Sprintf("触发了：%v", matcher.(CommandMatcher).GetCommand()))
				if matcher.IsBlock() {
					return false
				}
			}
		}

		// 处理正则匹配
		if matcher.(CommandMatcher).GetCommand() == "" && matcher.(CommandMatcher).GetRegexMatcher() != "" {
			compile := regexp.MustCompile(matcher.(CommandMatcher).GetRegexMatcher())
			if compile.MatchString(ctx.Event.Message.CQString()) {
				ctx.State.Args = commands[1:]
				ctx.State.Cmd = matcher.(CommandMatcher).GetRegexMatcher()
				ctx.State.Allies = matcher.(CommandMatcher).GetAlias()
				ctx.State.RegexResult = compile.FindStringSubmatch(ctx.Event.Message.CQString())

				doHandle(matcher, ctx)
				log.Infoln(fmt.Sprintf("触发了：%v", getFunctionName(matcher.GetHandler(), '/')))
				if matcher.IsBlock() {
					return false
				}
			}
		}
		return true
	})

	// 如果出发了command事件则不再触发message事件
	if a == 1 {
		return
	}
	s := new(State)
	s.Data = make(map[string]interface{})
	checkOnlyTome(ctx.Event, s)
	log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		ctx.Event.MessageType, ctx.Event.GroupId, ctx.Event.UserId, ctx.RawEvent))

	matcherChin.forEach(MESSAGE, func(matcher Matcher) bool {
		if matcher.GetType() != "" && matcher.GetType() != ctx.Event.MessageType {
			return true
		}
		for _, group := range matcher.GetDisAbleGroup() {
			if ctx.Event.GroupId == group {
				return true
			}
		}
		rule := checkRule(matcher.GetRules(), ctx)
		if !rule {
			return true
		}
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(getFunctionName(matcher.GetHandler(), '/') + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			matcher.GetHandler()(ctx)
		}()

		return !matcher.IsBlock()
	})
}

// processRequestEventHandle
/*
   @Description: Request类型event处理
   @param event Event
*/
func processRequestEventHandle(ctx *Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln(err)
			log.Errorln("request事件处理器出现不可挽回的错误")
		}
	}()
	matcherChin.forEach(REQUEST, func(matcher Matcher) bool {
		if matcher.GetType() != "" && matcher.GetType() != ctx.Event.RequestType {
			return true
		}
		for _, group := range matcher.GetDisAbleGroup() {
			if ctx.Event.GroupId == group {
				return true
			}
		}
		rule := checkRule(matcher.GetRules(), ctx)
		if !rule {
			return true
		}
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(getFunctionName(matcher.GetHandler(), '/') + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			matcher.GetHandler()(ctx)
		}()

		return !matcher.IsBlock()
	})
}

// processMetaEventHandle
/**
 * @Description:
 * @param event
 * example
 */
func processMetaEventHandle(ctx *Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Errorln("元事件处理器错误")
			log.Errorln(err)
		}
	}()
	matcherChin.forEach(META, func(matcher Matcher) bool {
		for _, group := range matcher.GetDisAbleGroup() {
			if ctx.Event.GroupId == group {
				return true
			}
		}
		rule := checkRule(matcher.GetRules(), ctx)
		if !rule {
			return true
		}
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Errorln(getFunctionName(matcher.GetHandler(), '/') + "发生不可挽回的错误")
					log.Errorln(err)
				}
			}()
			matcher.GetHandler()(ctx)
		}()

		return !matcher.IsBlock()
	})
}

// GetBotById
/*
   @Description:
   @param id int
   @return Api
*/
func GetBotById(id int64) API {
	bots := driver.GetBot(id)
	return bots.(API)
}
