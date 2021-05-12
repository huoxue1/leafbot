package plugins

import (
	"fmt"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/cqCode"
)

// UseCreateQrCode 生成二维码的插件
func UseCreateQrCode() {
	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		switch len(args) {
		case 0:
			{
				bot.Send(event, "参数不足")
			}
		case 1:
			{
				bot.Send(event, cqCode.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=0&e=L&p=15&url=%s", args[0]), map[string]interface{}{"cache": 0, "c": 3}))
			}
		case 2:
			{
				bot.Send(event, cqCode.Image(fmt.Sprintf("https://api.isoyu.com/qr/?m=%v&e=L&p=15&url=%s", args[1], args[0]), map[string]interface{}{"cache": 0, "c": 3}))
			}
		}
	}, "/creatQrCode", []string{"生成二维码"}, nil, 10, false)
}
