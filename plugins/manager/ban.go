package manager

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"strconv"
)

func InitBanPlugin() {
	leafBot.OnCommand("/ban").
		AddRule(leafBot.OnlySuperUser).
		SetBlock(false).
		AddAllies("禁言").
		SetWeight(10).
		SetPluginName("禁言").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				msgs := event.GetMsg()
				var banIds []int
				duration := 10
				for _, msg := range msgs {
					if msg.Type == "text" {
						if msg.Data["text"] == "禁言" {
							continue
						} else {
							long, err := strconv.Atoi(msg.Data["text"])
							if err != nil {
								continue
							}
							duration = long
						}
					}

					if msg.Type == "at" {
						banId, err := strconv.Atoi(msg.Data["qq"])
						if err != nil {
							continue
						}
						banIds = append(banIds, banId)
					}
				}
				if len(banIds) < 1 {
					bot.Send(event, message.Text("请艾特被禁言的成员才能禁言"))
					return
				}
				for _, id := range banIds {
					bot.SetGroupBan(event.GroupId, id, duration*60)
				}
			})
}
