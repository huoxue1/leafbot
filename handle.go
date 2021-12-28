package leafbot

import (
	"strconv"
	"strings"
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

		OnStart(start string, options ...Option) Matcher
		OnEnd(end string, options ...Option) Matcher
		OnFullMatch(content string, options ...Option) Matcher
		OnFullMatchGroup(content string, options ...Option) Matcher

		OnConnect(options ...Option) Matcher
		OnDisConnect(options ...Option) Matcher
	}
)

type (
	// Option
	/*
	 * 用于传递参数
	 */
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
	// Plugin
	/*
	 * 用于记录插件的基本信息
	 */
	Plugin struct {
		Name     string
		Help     map[string]string
		Matchers []Matcher
	}
)

// OnStart
/**
 * @Description: 匹配消息开头
 * @receiver p
 * @param start
 * @param options
 * @return Matcher
 */
func (p *Plugin) OnStart(start string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = ""
	d.PluginType = MESSAGE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	d.rules = append(d.rules, func(ctx *Context) bool {
		if strings.HasPrefix(ctx.Event.Message.ExtractPlainText(), start) {
			return true
		}
		return false
	})
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnEnd
/**
 * @Description: 匹配消息结尾
 * @receiver p
 * @param end
 * @param options
 * @return Matcher
 */
func (p *Plugin) OnEnd(end string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = ""
	d.PluginType = MESSAGE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	d.rules = append(d.rules, func(ctx *Context) bool {
		if strings.HasSuffix(ctx.Event.Message.ExtractPlainText(), end) {
			return true
		}
		return false
	})
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnFullMatch(content string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = ""
	d.PluginType = MESSAGE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	d.rules = append(d.rules, func(ctx *Context) bool {
		if ctx.Event.Message.ExtractPlainText() == content {
			return true
		}
		return false
	})
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnFullMatchGroup(content string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.HandleType = ""
	d.PluginType = MESSAGE
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	d.rules = append(d.rules, func(ctx *Context) bool {
		if ctx.Event.Message.ExtractPlainText() == content && ctx.Event.MessageType == "group" {
			return true
		}
		return false
	})
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnConnect(options ...Option) Matcher {
	d := new(defaultMatcher)
	d.PluginType = CONNECT
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	p.Matchers = append(p.Matchers, d)
	return d
}

func (p *Plugin) OnDisConnect(options ...Option) Matcher {
	d := new(defaultMatcher)
	d.PluginType = DISCONNECT
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	p.Matchers = append(p.Matchers, d)
	return d
}

func (d *defaultMatcher) Enabled() bool {
	return d.Enable
}

// NewPlugin
/**
 * @Description: 新建一个插件
 * @param name 插件名
 * @return *Plugin 返回一个插件类型对象的指针
 */
func NewPlugin(name string) *Plugin {
	p := new(Plugin)
	p.Name = name
	plugins = append(plugins, p)
	return p
}

// OnRegex
/**
 * @Description:
 * @receiver p
 * @param regexMatcher
 * @param options
 * @return Matcher
 */
func (p *Plugin) OnRegex(regexMatcher string, options ...Option) Matcher {
	d := new(defaultMatcher)
	d.regexMatcher = regexMatcher
	d.PluginType = COMMAND
	d.Enable = true
	if len(options) > 0 {
		d.allies = options[0].Allies
		d.weight = options[0].Weight
		d.block = options[0].Block
		d.rules = options[0].Rules
	}
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnCommand
/**
 * @Description:
 * @receiver p
 * @param command
 * @param options
 * @return Matcher
 */
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
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnMessage
/**
 * @Description:
 * @receiver p
 * @param messageType
 * @param options
 * @return Matcher
 */
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
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnRequest
/**
 * @Description:
 * @receiver p
 * @param requestType
 * @param options
 * @return Matcher
 */
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
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnNotice
/**
 * @Description:
 * @receiver p
 * @param noticeType
 * @param options
 * @return Matcher
 */
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
	p.Matchers = append(p.Matchers, d)
	return d
}

// OnMeta
/**
 * @Description:
 * @receiver p
 * @param options
 * @return Matcher
 */
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
	p.Matchers = append(p.Matchers, d)
	return d
}

// GetHelp
/**
 * @Description:
 * @receiver p
 * @return map[string]string
 */
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
