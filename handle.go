package leafbot

import (
	"strconv"
	"sync"
)

type handlers []Matcher
type Action func(ctx *Context)

var (
	pluginNum   = 1
	matcherChin = handlers{}
	plugins     []*Plugin
	lock        sync.Mutex
)

const (
	COMMAND    = "command"
	MESSAGE    = "message"
	REQUEST    = "request"
	REGEX      = "regex"
	NOTICE     = "notice"
	META       = "meta"
	CONNECT    = "connect"
	DISCONNECT = "dis_connect"
)

//State
// @Description: sdk处理消息后将内容传递给plugin
//
type State struct {
	Args        []string
	Cmd         string
	Allies      []string
	RegexResult []string
	Data        map[string]interface{}
}

// 插件的基本接口，实现Handler接口即可成为一个插件
type (
	Matcher interface {
		MatcherSet
		Enabled() bool
		GetHandler() Action
		GetRules() []Rule
		GetWeight() int
		IsBlock() bool
		GetDisAbleGroup() []int64
		GetType() string
		GetPluginType() string
	}

	MatcherSet interface {
		AddRule(rule Rule) Matcher
		SetWeight(weight int) Matcher
		SetBlock(block bool) Matcher
		SetAllies(allies []string) Matcher

		Handle(action Action)
	}

	CommandMatcher interface {
		Matcher
		GetCommand() string
		GetAlias() []string
		GetRegexMatcher() string
	}

	PluginInt interface {
		GetHelp() map[string]string
		GetName() string
		GetMather() []Matcher
	}

	basePlugin interface {
		OnCommand(command string, options ...Option) Matcher
		OnMessage(messageType string, options ...Option) Matcher
		OnRequest(requestType string, options ...Option) Matcher
		OnNotice(noticeType string, options ...Option) Matcher
		OnMeta(options ...Option) Matcher
		OnRegex(regexMatcher string, options ...Option) Matcher
	}
)

type (
	Option struct {
		Weight int
		Block  bool
		Allies []string
		Rules  []Rule
	}
)

type (
	defaultMatcher struct {
		ID     string
		Enable bool
		// 对应MessageType,RequestType,NoticeType
		HandleType string

		// 插件的类型
		PluginType   string
		disableGroup []int64
		handle       func(ctx *Context)
		command      string
		allies       []string
		rules        []Rule
		weight       int
		block        bool
		regexMatcher string
	}

	Plugin struct {
		Name     string
		Help     map[string]string
		Matchers []Matcher
	}
)

func (d *defaultMatcher) Enabled() bool {
	return d.Enable
}

//NewPlugin
/**
 * @Description: 新建一个插件
 * @param name 插件名
 * @return *Plugin 返回一个插件类型对象的指针
 */
func NewPlugin(name string) *Plugin {
	p := new(Plugin)
	p.Name = name
	return p
}

func (p *Plugin) OnRegex(regexMatcher string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.regexMatcher = regexMatcher
	d.PluginType = REGEX
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnCommand(command string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.command = command
	d.PluginType = COMMAND
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnMessage(messageType string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = messageType
	d.PluginType = MESSAGE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnRequest(requestType string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = requestType
	d.PluginType = REQUEST
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnNotice(noticeType string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = noticeType
	d.PluginType = NOTICE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnMeta(options ...Option) Matcher {
	d := new(defaultMatcher)
	d.PluginType = META
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	plugins = append(plugins, p)
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) GetHelp() map[string]string {
	return p.Help
}

func (p *Plugin) GetName() string {
	return p.Name
}

func (p *Plugin) GetMather() []Matcher {
	return p.Matchers
}

func (d *defaultMatcher) AddRule(rule Rule) Matcher {
	d.rules = append(d.rules, rule)
	return d
}

func (d *defaultMatcher) SetWeight(weight int) Matcher {
	d.weight = weight
	return d
}

func (d *defaultMatcher) SetBlock(block bool) Matcher {
	d.block = block
	return d
}

func (d *defaultMatcher) SetAllies(allies []string) Matcher {
	d.allies = allies
	return d
}

func (d *defaultMatcher) Handle(action Action) {
	d.handle = action
	lock.Lock()
	d.ID = strconv.Itoa(pluginNum)
	lock.Unlock()
	matcherChin = append(matcherChin, d)
}

func (d *defaultMatcher) GetCommand() string {
	return d.command
}

func (d *defaultMatcher) GetAlias() []string {
	return d.allies
}

func (d *defaultMatcher) GetRegexMatcher() string {
	return d.regexMatcher
}

func (d *defaultMatcher) GetHandler() Action {
	return d.handle
}

func (d *defaultMatcher) GetRules() []Rule {
	return d.rules
}

func (d *defaultMatcher) GetWeight() int {
	return d.weight
}

func (d *defaultMatcher) IsBlock() bool {
	return d.block
}

func (d *defaultMatcher) GetDisAbleGroup() []int64 {
	return d.disableGroup
}

func (d *defaultMatcher) GetType() string {
	return d.HandleType
}

func (d *defaultMatcher) GetPluginType() string {
	return d.PluginType
}

func (h handlers) forEach(pluginType string, handle func(matcher Matcher) bool) {
	for _, handler := range h {
		if handler.GetPluginType() == pluginType {
			b := handle(handler)
			if !b {
				break
			}
		}
	}
}

func (h handlers) Len() int {
	return len(h)
}
func (h handlers) Less(i, j int) bool {
	return h[i].GetWeight() < h[j].GetWeight()
}
func (h handlers) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
