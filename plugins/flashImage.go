package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"strconv"
	"strings"
	"time"
)

/*
	当获取到闪照信息之后，
	会向提供的qq号进行转发该闪照
*/
func UseFlashImage(userId int) {
	leafBot.OnMessage("").AddRule(FlashMessageRule).AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		if userId == 0 {
			userId = leafBot.DefaultConfig.Admin
		}
		mess := message.MessageSegment{}
		if event.MessageType == "group" {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群" + strconv.Itoa(event.GroupId) + "用户" +
				strconv.Itoa(event.UserId) + "所发闪照")
		} else {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊信息" + "用户" +
				strconv.Itoa(event.UserId) + "所发闪照")
		}
		bot.SendPrivateMsg(userId, append(message.ParseMessageFromString(strings.Replace(event.Message, "type=flash,", "", -1)), mess))
	})
}

func UseFlashImageToGroup(groupId int) {

	leafBot.
		OnMessage("").
		AddRule(FlashMessageRule).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot) {
				mess := message.MessageSegment{}
				if event.MessageType == "group" {
					mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群" + strconv.Itoa(event.GroupId) + "用户" +
						strconv.Itoa(event.UserId) + "所发闪照")
				} else {
					mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊信息" + "用户" +
						strconv.Itoa(event.UserId) + "所发闪照")
				}
				bot.SendGroupMsg(groupId, append(message.ParseMessageFromString(strings.Replace(event.Message, "type=flash,", "", -1)), mess))
			})

}

func FlashMessageRule(event leafBot.Event, bot *leafBot.Bot) bool {
	if strings.HasPrefix(event.Message, "[CQ:image,type=flash") && strings.HasSuffix(event.Message, "]") {
		return true
	} else {
		return false
	}
}
