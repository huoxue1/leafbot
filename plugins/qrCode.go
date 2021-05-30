package plugins

import (
	"fmt" //nolint:gci
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

// UseCreateQrCode 生成二维码的插件
func UseCreateQrCode() {

	leafBot.OnCommand("/createQrcode").
		AddAllies("生产二维码").
		SetWeight(10).
		SetPluginName("二维码生成").
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				switch len(args) {
				case 0:
					{
						bot.Send(event, "参数不足")
					}
				case 1:
					{
						bot.Send(event, message.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=0&e=L&p=15&url=%s", args[0])).Add("c", 3).Add("cache", 0))
					}
				case 2:
					{
						bot.Send(event, message.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=%v&e=L&p=15&url=%s", args[1], args[0])).Add("cache", 0).Add("c", 3))
					}
				}
			})

}
