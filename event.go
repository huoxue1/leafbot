package leafBot

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type (
	MessageChain []*messageHandle
	RequestChain []*requestHandle
	NoticeChain  []*noticeHandle
	CommandChain []*commandHandle
	MetaChain    []*metaHandle
)

var (
	c = make(chan Event, 10)
)

var (
	MessageHandles MessageChain
	RequestHandles RequestChain
	NoticeHandles  NoticeChain
	CommandHandles CommandChain
	MetaHandles    MetaChain

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

	messageHandle struct {
		handle      func(event Event, bot *Bot)
		messageType string
		rules       []Rule
		weight      int
	}

	requestHandle struct {
		handle      func(event Event, bot *Bot)
		requestType string
		rules       []Rule
		weight      int
	}

	noticeHandle struct {
		handle     func(event Event, bot *Bot)
		noticeType string
		rules      []Rule
		weight     int
	}
	commandHandle struct {
		handle  func(event Event, bot *Bot, args []string)
		command string
		allies  []string
		rules   []Rule
		weight  int
		block   bool
	}
	metaHandle struct {
		handle func(event Event, bot *Bot)
		rules  []Rule
		weight int
	}

	Rule struct {
		RuleCheck func(Event, ...interface{}) bool
		Dates     []interface{}
	}
)

func (m MessageChain) Len() int {
	return len(m)
}
func (m MessageChain) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}
func (m MessageChain) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MetaChain) Len() int {
	return len(m)
}
func (m MetaChain) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}
func (m MetaChain) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (r RequestChain) Len() int {
	return len(r)
}
func (r RequestChain) Less(i, j int) bool {
	return r[i].weight < r[j].weight
}
func (r RequestChain) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (n NoticeChain) Len() int {
	return len(n)
}
func (n NoticeChain) Less(i, j int) bool {
	return n[i].weight < n[j].weight
}
func (n NoticeChain) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (c CommandChain) Len() int {
	return len(c)
}
func (c CommandChain) Less(i, j int) bool {
	return c[i].weight < c[j].weight
}
func (c CommandChain) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// AddMessageHandle
/*
   @Description:
   @param messageType string
   @param rules []Rule
   @param handles ...func(event Event, bot *Bot)
*/
func AddMessageHandle(messageType string, rules []Rule, handles ...func(event Event, bot *Bot)) {
	for _, handle := range handles {
		MessageHandles = append(MessageHandles, &messageHandle{
			handle:      handle,
			messageType: messageType,
			rules:       rules,
		})
	}

}

// AddNoticeHandle
/*
   @Description:
   @param noticeType string
   @param rules []Rule
   @param weight int
   @param handles ...func(event Event, bot *Bot)
*/
func AddNoticeHandle(noticeType string, rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
	for _, handle := range handles {
		NoticeHandles = append(NoticeHandles, &noticeHandle{
			handle:     handle,
			noticeType: noticeType,
			rules:      rules,
			weight:     weight,
		})
	}

}

// AddRequestHandle
/*
   @Description:
   @param requestType string
   @param rules []Rule
   @param weight int
   @param handles ...func(event Event, bot *Bot)
*/
func AddRequestHandle(requestType string, rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
	for _, handle := range handles {
		RequestHandles = append(RequestHandles, &requestHandle{
			handle:      handle,
			requestType: requestType,
			rules:       rules,
			weight:      weight,
		})
	}

}

// AddMetaHandles
/*
   @Description:
   @param rules []Rule
   @param weight int
   @param handles ...func(event Event, bot *Bot)
*/
func AddMetaHandles(rules []Rule, weight int, handles ...func(event Event, bot *Bot)) {
	for _, handle := range handles {
		MetaHandles = append(MetaHandles, &metaHandle{
			handle: handle,
			rules:  rules,
			weight: weight,
		})
	}
}

// AddCommandHandle
/*
   @Description:
   @param handle func(event Event, bot *Bot, args []string)
   @param command string
   @param allies []string
   @param rules []Rule
   @param weight int
   @param block bool
*/
func AddCommandHandle(handle func(event Event, bot *Bot, args []string), command string, allies []string, rules []Rule, weight int, block bool) {
	CommandHandles = append(CommandHandles, &commandHandle{
		handle:  handle,
		command: command,
		allies:  allies,
		rules:   rules,
		weight:  weight,
		block:   block,
	})
}

// eventMain
/*
   @Description:
*/
func eventMain() {
	sort.Sort(&MessageHandles)
	sort.Sort(&RequestHandles)
	sort.Sort(&NoticeHandles)
	sort.Sort(&CommandHandles)
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
	go func() {
		for true {
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
			go sessionHandle(event)
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
   @Description:
   @param event Event
*/
func viewsMessage(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Infoln(err)
		}
	}()

	switch event.PostType {
	case "message":
		c <- event
		go processMessageHandle()

	case "notice":
		go processNoticeHandle(event)
	case "request":
		log.Infoln(fmt.Sprintf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType, event.GroupId, event.UserId))
		go processRequestEventHandle(event)

	case "meta_event":
		log.Infoln(fmt.Sprintf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%d",
			event.PostType, event.MetaEventType, event.Interval))
		go processMetaEventHandle(event)
	}
}

// processNoticeHandle
/*
   @Description:
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
		if !rule {
			continue
		}
		if v.noticeType == "" {
			v.noticeType = event.NoticeType
		}
		if v.noticeType == event.NoticeType {
			go v.handle(event, GetBotById(event.SelfId))
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
		check := rule.RuleCheck(event, rule.Dates)
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
		rule := checkRule(event, handle.rules)
		if handle.rules == nil {
			rule = true
		}
		if !rule {
			continue
		}
		commands := strings.Split(event.Message, " ")
		if len(commands) < 1 {
			continue
		}
		if commands[0] == handle.command {
			a = 1
			go handle.handle(event, GetBotById(event.SelfId), commands[1:])
			log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
				"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, event.Message, handle.command))
			if handle.block {
				return
			}
		}
		for _, ally := range handle.allies {
			if ally == commands[0] {
				a = 1
				go handle.handle(event, GetBotById(event.SelfId), commands[1:])
				log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command\n\t\t\t\t\t触发了：%v", event.MessageType, event.GroupId, event.UserId, event.Message, handle.command))
				if !handle.block {
					return
				}
			}
		}
	}
	if a == 1 {
		return
	}
	log.Infoln(fmt.Sprintf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType, event.GroupId, event.UserId, event.Message))
	for _, handle := range MessageHandles {
		if handle.messageType != "" && handle.messageType != event.MessageType {
			continue
		}
		rule := checkRule(event, handle.rules)
		if !rule {
			continue
		}
		go handle.handle(event, GetBotById(event.SelfId))
	}
}

// processRequestEventHandle
/*
   @Description:
   @param event Event
*/
func processRequestEventHandle(event Event) {
	for _, handle := range RequestHandles {
		rule := checkRule(event, handle.rules)
		if handle.rules == nil {
			rule = true
		}
		if !rule {
			continue
		}
		go handle.handle(event, GetBotById(event.SelfId))
	}
}

// processMetaEventHandle
/*
   @Description:
   @param event Event
*/
func processMetaEventHandle(event Event) {

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
