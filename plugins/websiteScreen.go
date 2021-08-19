package plugins

import (
	"encoding/base64"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/3343780376/leafBot/utils"
)

func WebSiteScreenInit() {
	leafBot.OnCommand(">website").AddAllies("网页截图").SetPluginName("网页长截图").SetCD("default", 0).SetBlock(false).SetWeight(10).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		if len(state.Args) < 1 {
			bot.Send(event, message.Text("参数不足，详情参考帮助菜单"))
			return
		}
		data, err := utils.GetPWScreen(state.Args[0])
		if err != nil {
			bot.Send(event, message.Text("获取截图错误"+err.Error()))
			return
		}
		bot.Send(event, message.Image("base64://"+base64.StdEncoding.EncodeToString(data)))
	})
}
