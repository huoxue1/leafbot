package plugins

import (
	"encoding/base64"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/3343780376/leafBot/utils"
)

func init() {
	leafBot.OnCommand("/help").
		AddAllies("帮助").
		AddRule(leafBot.OnlyToMe).
		SetPluginName("帮助").
		SetBlock(false).
		SetCD("default", 0).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				bot.Send(event, message.Text("downloading image ......"))
				screen, err := utils.GetPWScreen("https://huoxue1.github.io/leafBot/Features")
				if err != nil {
					bot.Send(event, message.Text("获取帮助文档失败"+err.Error()))
					return
				}
				bot.Send(event, message.Image("base64://"+base64.StdEncoding.EncodeToString(screen)))
			})
}
