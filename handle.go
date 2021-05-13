package leafBot

type (
	PretreatmentHandle struct {
		handle func(event Event, bot *Bot) bool
		rules  []Rule
		weight int
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
)

type CommandInt interface {
	AddAllies(allies string) commandHandle
	AddRule(rule Rule) commandHandle
	SetWeight(weight int) commandHandle
	SetBlock(IsBlock bool) commandHandle
	AddHandle(func(event Event, bot *Bot, args []string))
}

type NoticeInt interface {
	AddRule(rule Rule) noticeHandle
	SetWeight(weight int) noticeHandle
	AddHandle(func(event Event, bot *Bot))
}

type RequestInt interface {
	AddRule(rule Rule) requestHandle
	SetWeight(weight int) requestHandle
	AddHandle(func(event Event, bot *Bot))
}

type MessageInt interface {
	AddRule(rule Rule) messageHandle
	SetWeight(weight int) messageHandle
	AddHandle(func(event Event, bot *Bot))
}

type PretreatmentInt interface {
	AddRule(rule Rule) PretreatmentHandle
	SetWeight(weight int) PretreatmentHandle
	AddHandle(func(event Event, bot *Bot) bool)
}

type MetaInt interface {
	AddRule(rule Rule) MetaInt
	SetWeight(weight int) MetaInt
	AddHandle(func(event Event, bot *Bot))
}

// OnCommand
/**
 * @Description: command触发handle
 * @param command
 * @return commandHandle
 */
func OnCommand(command string) *commandHandle {
	return &commandHandle{command: command}
}

func OnNotice(noticeType string) *noticeHandle {
	return &noticeHandle{noticeType: noticeType}
}

func OnMessage(messageType string) *messageHandle {
	return &messageHandle{messageType: messageType}
}

func OnRequest(requestType string) *requestHandle {
	return &requestHandle{requestType: requestType}
}

func OnMeta() *metaHandle {
	return &metaHandle{}
}

func OnPretreatment() *PretreatmentHandle {
	return &PretreatmentHandle{}
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
	m.handle = f
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
func (c *commandHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	c.handle = f
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
	n.handle = f
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
	r.handle = f
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
func (m *messageHandle) AddHandle(f func(event Event, bot *Bot)) {
	m.handle = f
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
	p.handle = f
	PretreatmentHandles = append(PretreatmentHandles, p)
}
