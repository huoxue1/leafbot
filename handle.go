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
	SetBlock(IsBlock bool) noticeHandle
	AddHandle(func(event Event, bot *Bot, args []string))
}

type RequestInt interface {
	AddRule(rule Rule) requestHandle
	SetWeight(weight int) requestHandle
	SetBlock(IsBlock bool) requestHandle
	AddHandle(func(event Event, bot *Bot, args []string))
}

type MessageInt interface {
	AddRule(rule Rule) messageHandle
	SetWeight(weight int) messageHandle
	SetBlock(IsBlock bool) messageHandle
	AddHandle(func(event Event, bot *Bot, args []string))
}

type PretreatmentInt interface {
	AddRule(rule Rule) PretreatmentHandle
	SetWeight(weight int) PretreatmentHandle
	AddHandle(func(event Event, bot *Bot, args []string))
}

type MetaInt interface {
	AddRule(rule Rule) MetaInt
	SetWeight(weight int) MetaInt
	AddHandle(func(event Event, bot *Bot, args []string))
}

// AddRule
/**
 * @Description:
 * @receiver m
 * @param rule
 * @return MetaInt
 */
func (m metaHandle) AddRule(rule Rule) MetaInt {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver m
 * @param weight
 * @return MetaInt
 */
func (m metaHandle) SetWeight(weight int) MetaInt {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver m
 * @param f
 */
func (m metaHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}

// AddAllies
/**
 * @Description:
 * @receiver c
 * @param allies
 * @return commandHandle
 */
func (c commandHandle) AddAllies(allies string) commandHandle {
	panic("implement me")
}

// AddRule
/**
 * @Description:
 * @receiver c
 * @param rule
 * @return commandHandle
 */
func (c commandHandle) AddRule(rule Rule) commandHandle {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver c
 * @param weight
 * @return commandHandle
 */
func (c commandHandle) SetWeight(weight int) commandHandle {
	panic("implement me")
}

// SetBlock
/**
 * @Description:
 * @receiver c
 * @param IsBlock
 * @return commandHandle
 */
func (c commandHandle) SetBlock(IsBlock bool) commandHandle {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver c
 * @param f
 */
func (c commandHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}

// AddRule
/**
 * @Description:
 * @receiver n
 * @param rule
 * @return noticeHandle
 */
func (n noticeHandle) AddRule(rule Rule) noticeHandle {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver n
 * @param weight
 * @return noticeHandle
 */
func (n noticeHandle) SetWeight(weight int) noticeHandle {
	panic("implement me")
}

// SetBlock
/**
 * @Description:
 * @receiver n
 * @param IsBlock
 * @return noticeHandle
 */
func (n noticeHandle) SetBlock(IsBlock bool) noticeHandle {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver n
 * @param f
 */
func (n noticeHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}

// AddRule
/**
 * @Description:
 * @receiver r
 * @param rule
 * @return requestHandle
 */
func (r requestHandle) AddRule(rule Rule) requestHandle {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver r
 * @param weight
 * @return requestHandle
 */
func (r requestHandle) SetWeight(weight int) requestHandle {
	panic("implement me")
}

// SetBlock
/**
 * @Description:
 * @receiver r
 * @param IsBlock
 * @return requestHandle
 */
func (r requestHandle) SetBlock(IsBlock bool) requestHandle {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver r
 * @param f
 */
func (r requestHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}

// AddRule
/**
 * @Description:
 * @receiver m
 * @param rule
 * @return messageHandle
 */
func (m messageHandle) AddRule(rule Rule) messageHandle {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver m
 * @param weight
 * @return messageHandle
 */
func (m messageHandle) SetWeight(weight int) messageHandle {
	panic("implement me")
}

// SetBlock
/**
 * @Description:
 * @receiver m
 * @param IsBlock
 * @return messageHandle
 */
func (m messageHandle) SetBlock(IsBlock bool) messageHandle {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver m
 * @param f
 */
func (m messageHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}

// AddRule
/**
 * @Description:
 * @receiver p
 * @param rule
 * @return PretreatmentHandle
 */
func (p PretreatmentHandle) AddRule(rule Rule) PretreatmentHandle {
	panic("implement me")
}

// SetWeight
/**
 * @Description:
 * @receiver p
 * @param weight
 * @return PretreatmentHandle
 */
func (p PretreatmentHandle) SetWeight(weight int) PretreatmentHandle {
	panic("implement me")
}

// AddHandle
/**
 * @Description:
 * @receiver p
 * @param f
 */
func (p PretreatmentHandle) AddHandle(f func(event Event, bot *Bot, args []string)) {
	panic("implement me")
}
