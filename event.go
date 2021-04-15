package leafBot

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

type (
	MessageChain []*messageHandle
	RequestChain []*requestHandle
	NoticeChain  []*noticeHandle
	CommandChain []*commandHandle
)

var (
	nextEvent = false
	c         = make(chan Event, 10)
)

var (
	MessageHandles MessageChain
	RequestHandles RequestChain
	NoticeHandles  NoticeChain
	CommandHandles CommandChain
)

type (
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

	Rule struct {
		RuleCheck func(Event, ...interface{}) bool
		dates     []interface{}
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

func AddMessageHandle(handle func(event Event, bot *Bot), messageType string, rules []Rule) {

	MessageHandles = append(MessageHandles, &messageHandle{
		handle:      handle,
		messageType: messageType,
		rules:       rules,
	})
}

func AddNoticeHandle(handle func(event Event, bot *Bot), noticeType string, rules []Rule, weight int) {
	NoticeHandles = append(NoticeHandles, &noticeHandle{
		handle:     handle,
		noticeType: noticeType,
		rules:      rules,
		weight:     weight,
	})
}

func AddRequestHandle(handle func(event Event, bot *Bot), requestType string, rules []Rule, weight int) {
	RequestHandles = append(RequestHandles, &requestHandle{
		handle:      handle,
		requestType: requestType,
		rules:       rules,
		weight:      weight,
	})
}

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

func eventMain() {
	for _, handle := range CommandHandles {
		log.Infoln("已加载command响应器：" + handle.command)
	}
	for _, handle := range MessageHandles {
		log.Infoln("已加载message响应器：" + runtime.FuncForPC(reflect.ValueOf(handle.handle).Pointer()).Name())
	}
	for _, handle := range RequestHandles {
		log.Infoln("已加载request响应器：" + runtime.FuncForPC(reflect.ValueOf(handle.handle).Pointer()).Name())
	}
	for _, handle := range NoticeHandles {
		log.Infoln("已加载notice响应器：" + runtime.FuncForPC(reflect.ValueOf(handle.handle).Pointer()).Name())
	}
	go func() {
		for true {
			data, ok := <-eventChan
			if !ok {
				continue
			}
			var event Event
			err := json.Unmarshal(data, &event)
			if err != nil {
				log.Debugln("反向解析json失败" + err.Error() + "\n" + string(data))
			}
			go viewsMessage(event)
		}
	}()

}

func (b *Bot) GetNextEvent(n, UserId int) Event {
	defer func() {
		nextEvent = false
	}()

	nextEvent = true
	for i := 0; i < n; i++ {
		event := <-c
		if event.UserId != UserId {
			c <- event
			go processMessageHandle()
		} else {
			return event
		}
	}

	return Event{}
}

func viewsMessage(event Event) {
	defer func() {
		err := recover()
		if err != nil {
			log.Error(err)
		}
	}()
	sort.Sort(&MessageHandles)
	sort.Sort(&RequestHandles)
	sort.Sort(&NoticeHandles)
	sort.Sort(&CommandHandles)
	switch event.PostType {
	case "message":
		c <- event
		if !nextEvent {
			go processMessageHandle()
		}
	case "notice":
		go processNoticeHandle(event)
	case "request":
		log.Printf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType, event.GroupId, event.UserId)
		go processRequestEventHandle(event)

	case "meta_event":
		log.Printf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%d",
			event.PostType, event.MetaEventType, event.Interval)
		go processMetaEventHandle(event)
	}
}

func processNoticeHandle(event Event) {
	defer func() {
		err := recover()
		log.Println(err)
	}()
	log.Printf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
		event.NoticeType, event.GroupId, event.UserId)

	for _, v := range NoticeHandles {
		rule := checkRule(event, v.rules)
		if !rule {
			continue
		}
		if v.noticeType == "" {
			v.noticeType = event.NoticeType
		}
		if v.noticeType == event.NoticeType {
			go v.handle(event, getBotById(event.SelfId))
		}
	}
}

func checkRule(event Event, rules []Rule) bool {
	for _, rule := range rules {
		check := rule.RuleCheck(event, rule.dates)
		if !check {
			return false
		}
	}
	return true
}

func processMessageHandle() {
	event := <-c
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
			go handle.handle(event, getBotById(event.SelfId), commands[1:])
			log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
				"\n\t\t\t\t\tthis is a command", event.MessageType, event.GroupId, event.UserId, event.Message)
			if handle.block {
				return
			}
		}
		for _, ally := range handle.allies {
			if ally == commands[0] {
				go handle.handle(event, getBotById(event.SelfId), commands[1:])
				log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s"+
					"\n\t\t\t\t\tthis is a command", event.MessageType, event.GroupId, event.UserId, event.Message)
				if !handle.block {
					return
				}
			}
		}
	}
	log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType, event.GroupId, event.UserId, event.Message)
	for _, handle := range MessageHandles {
		rule := checkRule(event, handle.rules)
		if !rule {
			continue
		}
		go handle.handle(event, getBotById(event.SelfId))
	}
}

func processRequestEventHandle(event Event) {

}

func processMetaEventHandle(event Event) {

}

func getBotById(id int) *Bot {
	for _, bot := range config.Bots {
		if bot.SelfId == id {
			return bot
		}
	}
	return nil
}
