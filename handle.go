package leafBot

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"sync"
)

var plugins []*Plugin

type (
	MessageChain      []*messageHandle
	RequestChain      []*requestHandle
	NoticeChain       []*noticeHandle
	CommandChain      []*commandHandle
	MetaChain         []*metaHandle
	PretreatmentChain []*PretreatmentHandle

	ConnectChain    []*connectHandle
	DisConnectChain []*disConnectHandle
)

type Plugin struct {
	Name  string
	Helps []map[string]string
}

var (
	pluginNum = 0
	lock      sync.Mutex
)

//M
/*
	通过id从插件列表获取插件
	所有的插件队列都实现了该接口
*/
type M interface {
	get(id string) (interface{}, bool)
}

// State
// @Description: sdk处理消息后将内容传递给plugin
//
type State struct {
	Args        []string
	Cmd         string
	Allies      []string
	RegexResult []string
	Data        map[string]interface{}
}

type (

	/**
	     * @Description: 命令冷却的结构体
		    types: 冷却设置的类型
		    long: cd的长短，若types为rand则表示随机数的最大值
	*/
	coolDown struct {
		Types string `json:"types"`
		Long  int    `json:"long"`
	}

	PretreatmentHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot) bool
		rules        []Rule
		weight       int
	}
	messageHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot, state *State)
		messageType  string
		rules        []Rule
		weight       int
	}

	requestHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot)
		requestType  string
		rules        []Rule
		weight       int
	}

	noticeHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot)
		noticeType   string
		rules        []Rule
		weight       int
	}
	commandHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot, state *State)
		command      string
		allies       []string
		rules        []Rule
		weight       int
		block        bool
		cd           coolDown
		lastUseTime  int64
		regexMatcher string
	}

	metaHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(event Event, bot *Bot)
		rules        []Rule
		weight       int
	}
	connectHandle struct {
		BaseHandle
		handle func(connect Connect, bot *Bot)
	}
	disConnectHandle struct {
		BaseHandle
		handle func(selfId int)
	}

	Connect struct {
		SelfID     int
		Host       string
		ClientRole string
	}
)

func NewPlugin(name string) *Plugin {
	plugin := &Plugin{Name: name}
	plugins = append(plugins, plugin)
	return plugin
}

func OnStartWith(str string) *messageHandle {
	c := &messageHandle{}
	c.AddRule(func(event Event, bot *Bot, state *State) bool {
		if event.Message[0].Type != "text" {
			return false
		}
		if strings.HasPrefix(event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func (p *Plugin) SetHelp(help map[string]string) {
	p.Helps = append(p.Helps, help)
}

func (p *Plugin) OnRegex(regex string) *commandHandle {
	return &commandHandle{regexMatcher: regex}
}

func (p *Plugin) OnConnect() *connectHandle {
	return &connectHandle{}
}

func (p *Plugin) OnDisConnect() *disConnectHandle {
	return &disConnectHandle{}
}
func (p *Plugin) OnCommand(command string) *commandHandle {
	return &commandHandle{command: command}
}
func (p *Plugin) OnNotice(noticeType string) *noticeHandle {
	return &noticeHandle{noticeType: noticeType}
}

func (p *Plugin) OnMessage(messageType string) *messageHandle {
	return &messageHandle{messageType: messageType}
}

func (p *Plugin) OnRequest(requestType string) *requestHandle {
	return &requestHandle{requestType: requestType}
}

func (p *Plugin) OnMeta() *metaHandle {
	return &metaHandle{}
}

func (p *Plugin) OnPretreatment() *PretreatmentHandle {
	return &PretreatmentHandle{}
}

func (p *Plugin) OnStartWith(str string) *messageHandle {
	c := &messageHandle{}
	c.AddRule(func(event Event, bot *Bot, state *State) bool {
		if event.Message[0].Type != "text" {
			return false
		}
		if strings.HasPrefix(event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func (p *Plugin) OnTime(crons string, selfId int, handle func(bot *Bot)) {
	c2 := cron.New()
	_, err := c2.AddFunc(crons, func() {
		handle(GetBotById(selfId))
	})
	if err != nil {
		log.Errorln("运行定时任务出现错误")
	}
	c2.Start()
}

func (p *Plugin) OnEndWith(str string) *messageHandle {
	c := &messageHandle{}
	c.AddRule(func(event Event, bot *Bot, state *State) bool {
		if event.Message[0].Type != "text" {
			return false
		}
		if strings.HasSuffix(event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func (p *Plugin) OnKeyWords(keyword string) *messageHandle {
	m := &messageHandle{}
	m.rules = append(m.rules, func(event Event, bot *Bot, state *State) bool {
		if strings.Contains(event.Message.ExtractPlainText(), keyword) {
			state.Data["key_word"] = keyword
			return true
		}
		return false
	})
	return m
}

func OnRegex(regex string) *commandHandle {

	return &commandHandle{regexMatcher: regex}

}

func OnTime(crons string, selfId int, handle func(bot *Bot)) {
	c2 := cron.New()
	_, err := c2.AddFunc(crons, func() {
		handle(GetBotById(selfId))
	})
	if err != nil {
		log.Errorln("运行定时任务出现错误")
	}
	c2.Start()
}

func OnEndWith(str string) *messageHandle {
	c := &messageHandle{}
	c.AddRule(func(event Event, bot *Bot, state *State) bool {
		if event.Message[0].Type != "text" {
			return false
		}
		if strings.HasSuffix(event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func OnKeyWords(keyword string) *messageHandle {
	m := &messageHandle{}
	m.rules = append(m.rules, func(event Event, bot *Bot, state *State) bool {
		if strings.Contains(event.Message.ExtractPlainText(), keyword) {
			state.Data["key_word"] = keyword
			return true
		}
		return false
	})
	return m
}

// OnConnect
/**
 * @Description: 在bot进行连接使响应
 * @return *connectHandle
 * example
 */
func OnConnect() *connectHandle {
	return &connectHandle{}
}

// OnDisConnect
/**
 * @Description: 在bot断开连接时进行响应
 * @return *disConnectHandle
 * example
 */
func OnDisConnect() *disConnectHandle {
	return &disConnectHandle{}
}

// SetPluginName
/**
 * @Description: 设置插件名称
 * @receiver c
 * @param name  插件名称
 * @return *connectHandle
 * example
 */
func (c *connectHandle) SetPluginName(name string) *connectHandle {
	c.Name = name
	return c
}

func (d *disConnectHandle) SetPluginName(name string) *disConnectHandle {
	d.Name = name
	return d
}

func (m *metaHandle) SetPluginName(name string) *metaHandle {
	m.Name = name
	return m
}

func (c *commandHandle) SetPluginName(name string) *commandHandle {
	c.Name = name
	return c
}

func (n *noticeHandle) SetPluginName(name string) *noticeHandle {
	n.Name = name
	return n
}

func (r *requestHandle) SetPluginName(name string) *requestHandle {
	r.Name = name
	return r
}

func (m *messageHandle) SetPluginName(name string) *messageHandle {
	m.Name = name
	return m
}

func (p *PretreatmentHandle) SetPluginName(name string) *PretreatmentHandle {
	p.Name = name
	return p
}

type ConnectInt interface {
	SetPluginName(name string) *connectHandle
	AddHandle(func(connect Connect, bot *Bot))
}

type DisConnectInt interface {
	SetPluginName(name string) *disConnectHandle
	AddHandle(func(selfId int))
}

type CommandInt interface {
	SetPluginName(name string) *commandHandle
	AddAllies(allies string) *commandHandle
	AddRule(rule Rule) *commandHandle
	SetWeight(weight int) *commandHandle
	SetBlock(IsBlock bool) *commandHandle
	AddHandle(func(event Event, bot *Bot, args []string))
	SetCD(types string, long int) *commandHandle
}

func (c *commandHandle) SetCD(types string, long int) *commandHandle {
	c.cd = coolDown{Types: types, Long: long}
	return c
}

type NoticeInt interface {
	SetPluginName(name string) *noticeHandle
	AddRule(rule Rule) *noticeHandle
	SetWeight(weight int) *noticeHandle
	AddHandle(func(event Event, bot *Bot))
}

type RequestInt interface {
	SetPluginName(name string) *requestHandle
	AddRule(rule Rule) *requestHandle
	SetWeight(weight int) *requestHandle
	AddHandle(func(event Event, bot *Bot))
}

type MessageInt interface {
	SetPluginName(name string) *messageHandle
	AddRule(rule Rule) *messageHandle
	SetWeight(weight int) *messageHandle
	AddHandle(func(event Event, bot *Bot))
}

type PretreatmentInt interface {
	SetPluginName(name string) *PretreatmentHandle
	AddRule(rule Rule) *PretreatmentHandle
	SetWeight(weight int) *PretreatmentHandle
	AddHandle(func(event Event, bot *Bot) bool)
}

type MetaInt interface {
	SetPluginName(name string) *metaHandle
	AddRule(rule Rule) *metaHandle
	SetWeight(weight int) *metaHandle
	AddHandle(func(event Event, bot *Bot))
}

// OnCommand
/**
 * @Description: command触发handle
 * @param 该插件响应的命令
 * @return *commandHandle
 * example
	leafBot.OnCommand("天气").
	SetPluginName("天气插件").
	AddHandle(
		func(event LeafBot.Event,bot *leafBot.Bot,args []string){

		})
*/
func OnCommand(command string) *commandHandle {
	return &commandHandle{command: command}
}

// OnNotice
/**
 * @Description: notice触发handle
 * @param noticeType  notice事件类型
 * @return *noticeHandle
 * example
	响应戳一戳事件的示例
	leafBot.OnNotice(leafBot.NoticeTypeApi.Notify).
		SetPluginName("poke").
		AddRule(
			func(event leafBot.Event, bot *leafBot.Bot) bool {
				if event.SubType != "poke" || event.UserId == event.SelfId || int(event.TargetId) != event.SelfId {
					return false
				}
				return true
			}).SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot) {

			})
*/
func OnNotice(noticeType string) *noticeHandle {
	return &noticeHandle{noticeType: noticeType}
}

// OnMessage
/**
 * @Description: message事件触发handle
 * @param messageType  message事件的子类型
 * @return *messageHandle
 * example
 */
func OnMessage(messageType string) *messageHandle {
	return &messageHandle{messageType: messageType}
}

// OnRequest
/**
 * @Description: request事件触发handle
 * @param requestType  request事件的子类型
 * @return *requestHandle
 * example
 */
func OnRequest(requestType string) *requestHandle {
	return &requestHandle{requestType: requestType}
}

// OnMeta
/**
 * @Description: 元事件触发handle
 * @return *metaHandle
 * example
 */
func OnMeta() *metaHandle {
	return &metaHandle{}
}

// OnPretreatment
/**
 * @Description: 预处理事件
 * @return *PretreatmentHandle
 * example
 */
func OnPretreatment() *PretreatmentHandle {
	return &PretreatmentHandle{}
}

func (d *disConnectHandle) AddHandle(f func(selfId int)) {
	d.HandleType = "disConnect"
	lock.Lock()
	pluginNum++
	d.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	d.Enable = true
	d.handle = f
	if d.Name == "" {
		d.Name = getFunctionName(f, '/')
	}
	DisConnectHandles = append(DisConnectHandles, d)

}

// AddHandle
/**
 * @Description:
 * @receiver c
 * @param f
 */
func (c *connectHandle) AddHandle(f func(connect Connect, bot *Bot)) {
	c.HandleType = "connect"
	lock.Lock()
	pluginNum++
	c.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	c.Enable = true
	c.handle = f
	if c.Name == "" {
		c.Name = getFunctionName(f, '/')
	}
	ConnectHandles = append(ConnectHandles, c)
}

// AddRule
/**
 * @Description:
 * @receiver m
 * @param rule
 * @return metaHandle
 */
func (m *metaHandle) AddRule(rule Rule) *metaHandle {
	m.rules = append(m.rules, rule)
	return m
}

// SetWeight
/**
 * @Description:
 * @receiver m
 * @param weight
 * @return metaHandle
 */
func (m *metaHandle) SetWeight(weight int) *metaHandle {
	m.weight = weight
	return m
}

// AddHandle
/**
 * @Description:
 * @receiver m
 * @param f
 */
func (m *metaHandle) AddHandle(f func(event Event, bot *Bot)) {
	m.HandleType = "meta"
	m.handle = f
	m.Enable = true
	m.disableGroup = []int{}
	lock.Lock()
	pluginNum++
	m.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	if m.Name == "" {
		m.Name = getFunctionName(f, '/')
	}
	MetaHandles = append(MetaHandles, m)
}

// AddAllies
/**
 * @Description:
 * @receiver c
 * @param allies
 * @return commandHandle
 */
func (c *commandHandle) AddAllies(allies string) *commandHandle {
	c.allies = append(c.allies, allies)
	return c
}

// AddRule
/**
 * @Description:
 * @receiver c
 * @param rule
 * @return commandHandle
 */
func (c *commandHandle) AddRule(rule Rule) *commandHandle {
	c.rules = append(c.rules, rule)
	return c
}

// SetWeight
/**
 * @Description:
 * @receiver c
 * @param weight
 * @return commandHandle
 */
func (c *commandHandle) SetWeight(weight int) *commandHandle {
	c.weight = weight
	return c
}

// SetBlock
/**
 * @Description:
 * @receiver c
 * @param IsBlock
 * @return commandHandle
 */
func (c *commandHandle) SetBlock(IsBlock bool) *commandHandle {
	c.block = IsBlock
	return c
}

// AddHandle
/**
 * @Description:
 * @receiver c
 * @param f
 */
func (c *commandHandle) AddHandle(f func(event Event, bot *Bot, state *State)) {
	c.HandleType = "command"
	c.handle = f
	c.Enable = true
	c.disableGroup = []int{}
	lock.Lock()
	pluginNum++
	c.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	if c.Name == "" {
		c.Name = getFunctionName(f, '/')
	}
	CommandHandles = append(CommandHandles, c)
}

// AddRule
/**
 * @Description:
 * @receiver n
 * @param rule
 * @return noticeHandle
 */
func (n *noticeHandle) AddRule(rule Rule) *noticeHandle {
	n.rules = append(n.rules, rule)
	return n
}

// SetWeight
/**
 * @Description:
 * @receiver n
 * @param weight
 * @return noticeHandle
 */
func (n *noticeHandle) SetWeight(weight int) *noticeHandle {
	n.weight = weight
	return n
}

// AddHandle
/**
 * @Description:
 * @receiver n
 * @param f
 */
func (n *noticeHandle) AddHandle(f func(event Event, bot *Bot)) {
	n.HandleType = "notice"
	n.handle = f
	n.Enable = true
	lock.Lock()
	pluginNum++
	n.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	n.disableGroup = []int{}
	if n.Name == "" {
		n.Name = getFunctionName(f, '/')
	}
	NoticeHandles = append(NoticeHandles, n)
}

// AddRule
/**
 * @Description:
 * @receiver r
 * @param rule
 * @return requestHandle
 */
func (r *requestHandle) AddRule(rule Rule) *requestHandle {
	r.rules = append(r.rules, rule)
	return r
}

// SetWeight
/**
 * @Description:
 * @receiver r
 * @param weight
 * @return requestHandle
 */
func (r *requestHandle) SetWeight(weight int) *requestHandle {
	r.weight = weight
	return r
}

// AddHandle
/**
 * @Description:
 * @receiver r
 * @param f
 */
func (r *requestHandle) AddHandle(f func(event Event, bot *Bot)) {
	r.HandleType = "request"
	r.handle = f
	r.Enable = true
	r.disableGroup = []int{}
	lock.Lock()
	pluginNum++
	r.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	if r.Name == "" {
		r.Name = getFunctionName(f, '/')
	}
	RequestHandles = append(RequestHandles, r)
}

// AddRule
/**
 * @Description:
 * @receiver m
 * @param rule
 * @return messageHandle
 */
func (m *messageHandle) AddRule(rule Rule) *messageHandle {
	m.rules = append(m.rules, rule)
	return m
}

// SetWeight
/**
 * @Description:
 * @receiver m
 * @param weight
 * @return messageHandle
 */
func (m *messageHandle) SetWeight(weight int) *messageHandle {
	m.weight = weight
	return m
}

// AddHandle
/**
 * @Description:
 * @receiver m
 * @param f
 */
func (m *messageHandle) AddHandle(f func(event Event, bot *Bot, state *State)) {
	m.HandleType = "message"
	m.handle = f
	m.Enable = true
	m.disableGroup = []int{}
	lock.Lock()
	pluginNum++
	m.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	if m.Name == "" {
		m.Name = getFunctionName(f, '/')
	}
	MessageHandles = append(MessageHandles, m)
}

// AddRule
/**
 * @Description:
 * @receiver p
 * @param rule
 * @return PretreatmentHandle
 */
func (p *PretreatmentHandle) AddRule(rule Rule) *PretreatmentHandle {
	p.rules = append(p.rules, rule)
	return p
}

// SetWeight
/**
 * @Description:
 * @receiver p
 * @param weight
 * @return PretreatmentHandle
 */
func (p *PretreatmentHandle) SetWeight(weight int) *PretreatmentHandle {
	p.weight = weight
	return p
}

// AddHandle
/**
 * @Description:
 * @receiver p
 * @param f
 */
func (p *PretreatmentHandle) AddHandle(f func(event Event, bot *Bot) bool) {
	p.HandleType = "Pretreatment"
	p.handle = f
	p.Enable = true
	p.disableGroup = []int{}
	lock.Lock()
	pluginNum++
	p.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	if p.Name == "" {
		p.Name = getFunctionName(f, '/')
	}
	PretreatmentHandles = append(PretreatmentHandles, p)
}

func (m MessageChain) Len() int {
	return len(m)
}
func (m MessageChain) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}
func (m MessageChain) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (p PretreatmentChain) Len() int {
	return len(p)
}
func (p PretreatmentChain) Less(i, j int) bool {
	return p[i].weight < p[j].weight
}
func (p PretreatmentChain) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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

func (m MessageChain) get(id string) (interface{}, bool) {
	for _, handle := range m {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (n NoticeChain) get(id string) (interface{}, bool) {
	for _, handle := range n {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (r RequestChain) get(id string) (interface{}, bool) {
	for _, handle := range r {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (c CommandChain) get(id string) (interface{}, bool) {
	for _, handle := range c {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (m MetaChain) get(id string) (interface{}, bool) {
	for _, handle := range m {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (p PretreatmentChain) get(id string) (interface{}, bool) {
	for _, handle := range p {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (c ConnectChain) get(id string) (interface{}, bool) {
	for _, handle := range c {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (d DisConnectChain) get(id string) (interface{}, bool) {
	for _, handle := range d {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
