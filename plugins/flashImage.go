package plugins

import (
	"github.com/3343780376/leafBot"
	"strconv"
	"strings"
	"time"
)

/*
	当获取到闪照信息之后，
	会向提供的qq号进行转发该闪照
*/
func UseFlashImage(userId int) {
	leafBot.AddMessageHandle("", []leafBot.Rule{{RuleCheck: FlashMessageRule, Dates: nil}},
		func(event leafBot.Event, bot *leafBot.Bot) {
			mess := ""
			if event.MessageType == "group" {
				mess = time.Now().Format("2006-01-02 15:04:05") + "\n来自群" + strconv.Itoa(event.GroupId) + "用户" +
					strconv.Itoa(event.UserId) + "所发闪照"
			} else {
				mess = time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊信息" + "用户" +
					strconv.Itoa(event.UserId) + "所发闪照"
			}
			bot.SendPrivateMsg(userId, mess+strings.Replace(event.Message, "type=flash,", "", -1), false)
		})
}

func FlashMessageRule(event leafBot.Event, i ...interface{}) bool {
	if strings.HasPrefix(event.Message, "[CQ:image,type=flash") && strings.HasSuffix(event.Message, "]") {
		return true
	} else {
		return false
	}
}
