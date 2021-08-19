package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func UseSetuHandle() {
	leafBot.OnCommand("/setu").
		AddAllies("来点色图").
		SetWeight(10).
		SetBlock(false).
		SetPluginName("色图").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				if len(state.Args) < 1 {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Image("https://acg.toubiec.cn/random.php"))
				} else if state.Args[0] == "r18" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Image("https://api.pixivweb.com/anime18r.php?return=img"))
				} else if state.Args[0] == "true" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Image("https://api.pixivweb.com/api.php?return=img/json"))
				} else if state.Args[0] == "r18+true" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Image("https://api.pixivweb.com/bw.php?return=img"))
				}
			})

}
