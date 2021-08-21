package help

import (
	_ "embed"
	"encoding/base64"
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
)

//go:embed help.png
var image []byte

func init() {
	leafBot.OnCommand("help").
		AddAllies("帮助").
		AddRule(leafBot.OnlyToMe).
		SetWeight(12).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				bot.Send(event, append(message.Message{}, message.Image("base64://"+base64.StdEncoding.EncodeToString(image)), message.At(int64(event.UserId))))
			})
}
