// Package leafBot
// @Description:
package leafBot

import "strconv"

//Rule
/*
	rule类型
*/
type Rule func(Event, Api, *State) bool

//MustReply
/**
 * @Description:
 * @param event
 * @param api
 * @param state
 * @return bool
 * example
 */
func MustReply(event Event, api Api, state *State) bool {
	for _, segment := range event.Message {
		if segment.Type == "reply" {
			state.Data["reply_id"] = segment.Data["id"]
			id, err := strconv.Atoi(segment.Data["id"])
			if err != nil {
				return false
			}
			state.Data["reply_msg"] = api.(OneBotApi).GetMsg(int32(id))
			return true
		}
	}
	return false
}

//OnlyToMe
/**
 * @Description: 添加了该rule的插件需要在群里艾特或者私聊才会进行响应
 * @param event  leafBot event
 * @param bot    bot实例对象
 * @return bool  返回是否验证通过该rule
 * example
 */
func OnlyToMe(event Event, _ Api, state *State) bool {
	if b, ok := state.Data["only_tome"]; ok {
		return b.(bool)
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
func OnlySuperUser(event Event, _ Api, _ *State) bool {
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

//OnlyGroupMessage
/**
 * @Description:
 * @param event
 * @param _
 * @return bool
 * example
 */
func OnlyGroupMessage(event Event, _ Api) bool {
	return event.MessageType == "group"
}
