package leafBot

import (
	"strconv"
	"strings"
	"sync"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var plugins []*Plugin

type (
	messageChain      []*messageHandle
	requestChain      []*requestHandle
	noticeChain       []*noticeHandle
	commandChain      []*commandHandle
	metaChain         []*metaHandle
	pretreatmentChain []*PretreatmentHandle

	connectChain    []*connectHandle
	disConnectChain []*disConnectHandle
)

type (
	PretreatmentHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context) bool
		rules        []Rule
		weight       int
	}
	messageHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context)
		messageType  string
		rules        []Rule
		weight       int
	}

	requestHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context)
		requestType  string
		rules        []Rule
		weight       int
	}

	noticeHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context)
		noticeType   string
		rules        []Rule
		weight       int
	}
	commandHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context)
		command      string
		allies       []string
		rules        []Rule
		weight       int
		block        bool
		cd           CoolDown
		lastUseTime  int64
		regexMatcher string
	}

	metaHandle struct {
		BaseHandle
		disableGroup []int
		handle       func(ctx *Context)
		rules        []Rule
		weight       int
	}
	connectHandle struct {
		BaseHandle
		handle func(connect Connect, bot API)
	}
	disConnectHandle struct {
		BaseHandle
		handle func(selfId int64)
	}
)

type (
	NoticeInt interface {
		SetPluginName(name string) *noticeHandle
		AddRule(rule Rule) *noticeHandle
		SetWeight(weight int) *noticeHandle
		AddHandle(func(ctx *Context))
	}

	RequestInt interface {
		SetPluginName(name string) *requestHandle
		AddRule(rule Rule) *requestHandle
		SetWeight(weight int) *requestHandle
		AddHandle(func(ctx *Context))
	}

	MessageInt interface {
		SetPluginName(name string) *messageHandle
		AddRule(rule Rule) *messageHandle
		SetWeight(weight int) *messageHandle
		AddHandle(func(ctx *Context))
	}

	PretreatmentInt interface {
		SetPluginName(name string) *PretreatmentHandle
		AddRule(rule Rule) *PretreatmentHandle
		SetWeight(weight int) *PretreatmentHandle
		AddHandle(func(ctx *Context) bool)
	}

	MetaInt interface {
		SetPluginName(name string) *metaHandle
		AddRule(rule Rule) *metaHandle
		SetWeight(weight int) *metaHandle
		AddHandle(func(ctx *Context))
	}

	ConnectInt interface {
		SetPluginName(name string) *connectHandle
		AddHandle(func(connect Connect, bot API))
	}

	DisConnectInt interface {
		SetPluginName(name string) *disConnectHandle
		AddHandle(func(selfId int64))
	}

	CommandInt interface {
		SetPluginName(name string) *commandHandle
		AddAllies(allies string) *commandHandle
		AddRule(rule Rule) *commandHandle
		SetWeight(weight int) *commandHandle
		SetBlock(IsBlock bool) *commandHandle
		AddHandle(f func(ctx *Context))
		SetCD(types string, long int) *commandHandle
	}
)

type (
	Option struct {
		PluginName string
		Weight     int
		Block      bool
		Allies     []string
		Rules      []Rule
		CD         CoolDown
	}
)

//Plugin
// @Description:
//
type Plugin struct {
	Name    string
	Helps   []map[string]string
	handles []BaseHandle
}

var (
	pluginNum = 0
	lock      sync.Mutex
)

// M
/**
/*	通过id从插件列表获取插件
/*	所有的插件队列都实现了该接口
**/
type M interface {
	get(id string) (interface{}, bool)
}

type (

	//State
	// @Description: sdk处理消息后将内容传递给plugin
	//
	State struct {
		Args        []string
		Cmd         string
		Allies      []string
		RegexResult []string
		Data        map[string]interface{}
	}

	CoolDown struct {
		Types string `json:"types"`
		Long  int    `json:"long"`
	}

	Connect struct {
		SelfID     int64
		Host       string
		ClientRole string
	}
)

// NewPlugin
/**
 * @Description:
 * @param name
 * @return *Plugin
 * example
 */
func NewPlugin(name string) *Plugin {
	plugin := &Plugin{Name: name}
	plugins = append(plugins, plugin)
	return plugin
}

type Options struct {
}

func (p *Plugin) SetHelp(help map[string]string) {
	p.Helps = append(p.Helps, help)
}

func (p *Plugin) OnRegex(regex string, options ...Option) *commandHandle {
	if len(options) > 0 {
		return &commandHandle{
			regexMatcher: regex,
			allies:       options[0].Allies,
			weight:       options[0].Weight,
			block:        options[0].Block,
			rules:        options[0].Rules,
			cd:           options[0].CD,
		}
	}
	return &commandHandle{
		regexMatcher: regex,
	}
}

func (p *Plugin) OnFullMatch(match string, options ...Option) *messageHandle {
	rule := func(ctx *Context) bool {
		return ctx.Event.RawMessage == match
	}
	if len(options) > 0 {
		return &messageHandle{
			weight: options[0].Weight,
			rules:  append(options[0].Rules, rule),
		}
	}
	return &messageHandle{rules: []Rule{rule}}
}

func (p *Plugin) OnFullMatchGroup(matchs []string, options ...Option) *messageHandle {
	rule := func(ctx *Context) bool {
		for _, match := range matchs {
			if ctx.Event.RawMessage == match {
				return true
			}
		}
		return false
	}
	if len(options) > 0 {
		return &messageHandle{
			weight: options[0].Weight,
			rules:  append(options[0].Rules, rule),
		}
	}
	return &messageHandle{rules: []Rule{rule}}
}

func (p *Plugin) OnConnect() *connectHandle {
	return &connectHandle{}
}

func (p *Plugin) OnDisConnect() *disConnectHandle {
	return &disConnectHandle{}
}
func (p *Plugin) OnCommand(command string, options ...Option) *commandHandle {
	if len(options) > 0 {
		return &commandHandle{
			command: command,
			allies:  options[0].Allies,
			weight:  options[0].Weight,
			block:   options[0].Block,
			rules:   options[0].Rules,
			cd:      options[0].CD,
		}
	}
	return &commandHandle{command: command}
}
func (p *Plugin) OnNotice(noticeType string, options ...Option) *noticeHandle {
	if len(options) > 0 {
		return &noticeHandle{
			weight: options[0].Weight,
			rules:  options[0].Rules,
		}
	}
	return &noticeHandle{noticeType: noticeType}
}

func (p *Plugin) OnMessage(messageType string, options ...Option) *messageHandle {
	if len(options) > 0 {
		return &messageHandle{
			weight:      options[0].Weight,
			rules:       options[0].Rules,
			messageType: messageType,
		}
	}
	return &messageHandle{messageType: messageType}
}

func (p *Plugin) OnRequest(requestType string, options ...Option) *requestHandle {
	if len(options) > 0 {
		return &requestHandle{
			weight:      options[0].Weight,
			rules:       options[0].Rules,
			requestType: requestType,
		}
	}
	return &requestHandle{requestType: requestType}
}

func (p *Plugin) OnMeta() *metaHandle {
	return &metaHandle{}
}

func (p *Plugin) OnPretreatment() *PretreatmentHandle {
	return &PretreatmentHandle{}
}

func (p *Plugin) OnStartWith(str string, options ...Option) *messageHandle {
	c := &messageHandle{}
	if len(options) > 0 {
		c = &messageHandle{
			weight: options[0].Weight,
			rules:  options[0].Rules,
		}
	}
	c.AddRule(func(ctx *Context) bool {
		if ctx.Event.Message[0].Type != "text" {
			return false
		}
		if strings.HasPrefix(ctx.Event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func (p *Plugin) OnTime(crons string, selfId int, handle func(bot API)) {
	c2 := cron.New()
	_, err := c2.AddFunc(crons, func() {
		handle(GetBotById(selfId))
	})
	if err != nil {
		log.Errorln("运行定时任务出现错误")
	}
	c2.Start()
}

func (p *Plugin) OnEndWith(str string, options ...Option) *messageHandle {
	c := &messageHandle{}
	if len(options) > 0 {
		c = &messageHandle{
			weight: options[0].Weight,
			rules:  options[0].Rules,
		}
	}
	c.AddRule(func(ctx *Context) bool {
		if ctx.Event.Message[0].Type != "text" {
			return false
		}
		if strings.HasSuffix(ctx.Event.Message[0].Data["text"], str) {
			return true
		}
		return false
	})
	return c
}

func (p *Plugin) OnKeyWords(keyword string, options ...Option) *messageHandle {
	m := &messageHandle{}
	if len(options) > 0 {
		m = &messageHandle{
			weight: options[0].Weight,
			rules:  options[0].Rules,
		}
	}
	m.rules = append(m.rules, func(ctx *Context) bool {
		if strings.Contains(ctx.Event.Message.ExtractPlainText(), keyword) {
			ctx.State.Data["key_word"] = keyword
			return true
		}
		return false
	})
	return m
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

func (c *commandHandle) SetCD(types string, long int) *commandHandle {
	c.cd = CoolDown{Types: types, Long: long}
	return c
}

func (d *disConnectHandle) AddHandle(f func(selfId int64)) {
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
func (c *connectHandle) AddHandle(f func(connect Connect, bot API)) {
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
func (m *metaHandle) AddHandle(f func(ctx *Context)) {
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
func (c *commandHandle) AddHandle(f func(ctx *Context)) {
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
func (n *noticeHandle) AddHandle(f func(ctx *Context)) {
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
func (r *requestHandle) AddHandle(f func(ctx *Context)) {
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
func (m *messageHandle) AddHandle(f func(ctx *Context)) {
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
func (p *PretreatmentHandle) AddHandle(f func(ctx *Context) bool) {
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

func (m messageChain) Len() int {
	return len(m)
}
func (m messageChain) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}
func (m messageChain) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (p pretreatmentChain) Len() int {
	return len(p)
}
func (p pretreatmentChain) Less(i, j int) bool {
	return p[i].weight < p[j].weight
}
func (p pretreatmentChain) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (m metaChain) Len() int {
	return len(m)
}
func (m metaChain) Less(i, j int) bool {
	return m[i].weight < m[j].weight
}
func (m metaChain) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (r requestChain) Len() int {
	return len(r)
}
func (r requestChain) Less(i, j int) bool {
	return r[i].weight < r[j].weight
}
func (r requestChain) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (n noticeChain) Len() int {
	return len(n)
}
func (n noticeChain) Less(i, j int) bool {
	return n[i].weight < n[j].weight
}
func (n noticeChain) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (c commandChain) Len() int {
	return len(c)
}
func (c commandChain) Less(i, j int) bool {
	return c[i].weight < c[j].weight
}
func (c commandChain) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (m messageChain) get(id string) (interface{}, bool) {
	for _, handle := range m {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (n noticeChain) get(id string) (interface{}, bool) {
	for _, handle := range n {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (r requestChain) get(id string) (interface{}, bool) {
	for _, handle := range r {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (c commandChain) get(id string) (interface{}, bool) {
	for _, handle := range c {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (m metaChain) get(id string) (interface{}, bool) {
	for _, handle := range m {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (p pretreatmentChain) get(id string) (interface{}, bool) {
	for _, handle := range p {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}

func (c connectChain) get(id string) (interface{}, bool) {
	for _, handle := range c {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
func (d disConnectChain) get(id string) (interface{}, bool) {
	for _, handle := range d {
		if handle.ID == id {
			return handle, true
		}
	}
	return nil, false
}
