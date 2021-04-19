package plugins

import (
	"github.com/3343780376/leafBot"
	"strings"
)

/*
	当获取到闪照信息之后，
	会向提供的qq号进行转发该闪照
*/
func UseFlashImage(userId int) {
	leafBot.AddMessageHandle(leafBot.MessageTypeApi.Group, []leafBot.Rule{{RuleCheck: FlashMessageRule, Dates: nil}},
		func(event leafBot.Event, bot *leafBot.Bot) {
			bot.SendPrivateMsg(userId, strings.Replace(event.Message, "type=flash,", "", -1), false)
		})
}

func FlashMessageRule(event leafBot.Event, i ...interface{}) bool {
	if strings.HasPrefix(event.Message, "[CQ:image,type=flash") && strings.HasSuffix(event.Message, "]") {
		return true
	} else {
		return false
	}
}
