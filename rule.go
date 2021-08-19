package leafBot

import "strconv"

type (
	Rule func(Event, *Bot, *State) bool
)

// OnlyToMe
/**
 * @Description: 添加了该rule的插件需要在群里艾特或者私聊才会进行响应
 * @param event  leafBot event
 * @param bot    bot实例对象
 * @return bool  返回是否验证通过该rule
 * example
 */
func OnlyToMe(event Event, _ *Bot, _ *State) bool {
	if event.MessageType == "private" {
		return true
	}
	msg := event.GetMsg()
	for _, segment := range msg {
		if segment.Type == "at" && segment.Data["qq"] == strconv.Itoa(event.SelfId) {
			return true
		}
	}

	return false
}

// OnlySuperUser
/**
 * @Description: 加了该rule的插件只会对配置文件中配置的管理员用户进行响应
 * @param event  leafBot event
 * @param bot    bot实例对象
 * @return bool  是否通过该rule验证
 * example
 */
func OnlySuperUser(event Event, bot *Bot, _ *State) bool {
	if event.UserId == DefaultConfig.Admin {
		return true
	}
	for _, user := range DefaultConfig.SuperUser {
		if event.UserId == user {
			return true
		}
	}
	return false
}

func OnlyGroupMessage(event Event, bot *Bot) bool {

	return event.MessageType == "group"
}
