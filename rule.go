// Package leafBot
// @Description:
package leafBot

import "strconv"

//Rule
/*
	rule类型
*/
type Rule func(ctx *Context) bool

//MustReply
/**
 * @Description:
 * @param event
 * @param api
 * @param state
 * @return bool
 * example
 */
func MustReply(ctx *Context) bool {
	for _, segment := range ctx.Event.Message {
		if segment.Type == "reply" {
			ctx.State.Data["reply_id"] = segment.Data["id"]
			id, err := strconv.Atoi(segment.Data["id"])
			if err != nil {
				return false
			}
			ctx.State.Data["reply_msg"] = ctx.GetMsg(int32(id))
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
func OnlyToMe(ctx *Context) bool {
	if b, ok := ctx.State.Data["only_tome"]; ok {
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
func OnlySuperUser(ctx *Context) bool {
	if ctx.Event.UserId == defaultConfig.Admin {
		return true
	}
	for _, user := range defaultConfig.SuperUser {
		if ctx.Event.UserId == user {
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
func OnlyGroupMessage(ctx *Context) bool {
	return ctx.Event.MessageType == "group"
}
