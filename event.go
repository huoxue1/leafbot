package leafBot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/3343780376/leafBot/message" //nolint:gci
	log "github.com/sirupsen/logrus"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	c = make(chan Event, 10)
)

var ISGUI = true

func init() {
	if runtime.GOOS != "windows" {
		ISGUI = false
	}
}

// var
/**
 * @Description: 通向前端的通道
 * @return unc
 */
var (
	MessageChan = make(chan Event, 10)
	NoticeChan  = make(chan Event, 10)
	Request     = make(chan Event, 10)
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
   @Description:
   @param event Event
*/
func sessionHandle(event Event) {
	sessions.Range(func(key, value interface{}) bool {
		s := value.(session)
		rule := checkRule(event, s.rules)
		if s.rules == nil {
			rule = true
		}
		if rule {
			s.queue <- event
		}
		return true
	})
}

type (
	session struct {
		id    int
		queue chan Event
		rules []Rule
	}
)

// AddMessageHandle
/*
   @Description:
   @param messageType string
   @param rules []Rule
   @param handles ...func(event Event, bot *Bot)
*/
//func AddMessageHandle(messageType string, rules []Rule, handles ...func(event Event, bot *Bot)) {
//	for _, handle := range handles {
//		MessageHandles = append(MessageHandles, &messageHandle{
//			handle:      handle,
//			messageType: messageType,
//			rules:       rules,
//		})
//	}
//
//}
//
//func AddPretreatmentHandle(rules []Rule, weight int, handles ...func(event Event, bot *Bot) bool) {
//	for _, handle := range handles {
//		PretreatmentHandles = append(PretreatmentHandles, &PretreatmentHandle{
//			handle: handle,
//			rules:  rules,
//			weight: weight,
//		})
//	}
//}
//
//// AddNoticeHandle
///*
//   @Description:
//   @param noticeType string
//   @param rules []Rule
//   @param weight int
//   @param handles ...func(event Event, bot *Bot)
//*/
//func AddNoticeHandle(noticeType string, rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
//	for _, handle := range handles {
//		NoticeHandles = append(NoticeHandles, &noticeHandle{
//			handle:     handle,
//			noticeType: noticeType,
//			rules:      rules,
//			weight:     weight,
//		})
//	}
//}
//
//// AddRequestHandle
///*
//   @Description:
//   @param requestType string
//   @param rules []Rule
//   @param weight int
//   @param handles ...func(event Event, bot *Bot)
//*/
//func AddRequestHandle(requestType string, rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
//	for _, handle := range handles {
//		RequestHandles = append(RequestHandles, &requestHandle{
//			handle:      handle,
//			requestType: requestType,
//			rules:       rules,
//			weight:      weight,
//		})
//	}
//}
//
//// AddMetaHandles
///*
//   @Description:
//   @param rules []Rule
//   @param weight int
//   @param handles ...func(event Event, bot *Bot)
//*/
//func AddMetaHandles(rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
//	for _, handle := range handles {
//		MetaHandles = append(MetaHandles, &metaHandle{
//			handle: handle,
//			rules:  rules,
//			weight: weight,
//		})
//	}
//}
//
//// AddCommandHandle
///*
//   @Description:
//   @param handle func(event Event, bot *Bot, args []string)
//   @param command string
//   @param allies []string
//   @param rules []Rule
//   @param weight int
//   @param block bool
//*/
//func AddCommandHandle(handle func(event Event, bot *Bot, args []string), command string, allies []string, rules []Rule, weight int, block bool) {
//	CommandHandles = append(CommandHandles, &commandHandle{
//		handle:  handle,
//		command: command,
//		allies:  allies,
//		rules:   rules,
//		weight:  weight,
//		block:   block,
//	})
//}

// eventMain
/*
   @Description:
*/
func eventMain() {
	sort.Sort(&MessageHandles)
	sort.Sort(&RequestHandles)
	sort.Sort(&NoticeHandles)
	sort.Sort(&CommandHandles)

	for _, handle := range PretreatmentHandles {
		log.Infoln("已加载预处理器响应器：" + getFunctionName(handle.handle, '/'))
	}
	for _, handle := range CommandHandles {
		log.Infoln("已加载command响应器：" + handle.command)
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
   @Description:
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
   @Description:
   @receiver b
   @param rules ...Rule
   @return int  int
   @return chan  Event
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
   @Description:
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
			log.Infoln(err)
		}
	}()

	// 执行所有预处理handle
	for _, handle := range PretreatmentHandles {
		bot := GetBotById(event.SelfId)
		rule := checkRule(event, handle.rules)
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
		if ISGUI {
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
		log.Infoln(err)
	}()
	log.Infoln(fmt.Sprintf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
		event.NoticeType, event.GroupId, event.UserId))

	for _, v := range NoticeHandles {
		rule := checkRule(event, v.rules)
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
/*
   @Description:
   @param event Event
   @param rules []Rule
   @return bool
*/
func checkRule(event Event, rules []Rule) bool {
	for _, rule := range rules {
		check := rule(event, GetBotById(event.SelfId))
		if !check {
			return false
		}
	}
	return true
}

// getFunctionName
/*
   @Description:
   @param i interface{}
   @param seps ...rune
   @return string
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

// processMessageHandle
/*
   @Description:
*/
func processMessageHandle() {
	defer func() {
		err := recover()
		if err != nil {
			log.Infoln(err)
		}
	}()
	event := <-c
	a := 0
	log.Debugln(len(CommandHandles))
	for _, handle := range CommandHandles {

		disable := true
		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				disable = false
			}
		}
		if !disable {
			continue
		}

		rule := checkRule(event, handle.rules)
		if handle.rules == nil {
			rule = true
		}
		if !rule || !handle.Enable {
			continue
		}
		if event.Message[0].Type != "text" {
			continue
		}

		commands := strings.Split(event.Message[0].Data["text"], " ")
		if len(commands) < 1 {
			continue
		}
		if commands[0] == handle.command {
			a = 1

			go func(handle2 *commandHandle) {
				defer func() {
					err := recover()
					if err != nil {
						log.Errorln(handle2.Name + "发生不可挽回的错误")
						log.Errorln(err)
					}
				}()
				handle2.handle(event, GetBotById(event.SelfId), commands[1:])
			}(handle)
			log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
				"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, event.Message, handle.command))
			if handle.block {
				return
			}
		}
		for _, ally := range handle.allies {
			if ally == commands[0] {
				a = 1
				go func(handle2 *commandHandle) {
					defer func() {
						err := recover()
						if err != nil {
							log.Errorln(handle2.Name + "发生不可挽回的错误")
							log.Errorln(err)
						}
					}()
					handle2.handle(event, GetBotById(event.SelfId), commands[1:])
				}(handle)
				log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, event.Message, handle.command))
				if handle.block {
					return
				}
			}
		}
	}
	if a == 1 {
		return
	}

	go sessionHandle(event)

	log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType, event.GroupId, event.UserId, event.Message))
	for _, handle := range MessageHandles {
		if handle.messageType != "" && handle.messageType != event.MessageType {
			continue
		}

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules)
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
			handle2.handle(event, GetBotById(event.SelfId))
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
			log.Infoln(err)
		}
	}()
	for _, handle := range RequestHandles {

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules)
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
/*
   @Description:
   @param event Event
*/
func processMetaEventHandle(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Infoln(err)
		}
	}()
	for _, handle := range MetaHandles {

		for _, group := range handle.disableGroup {
			if event.GroupId == group {
				return
			}
		}

		rule := checkRule(event, handle.rules)
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
