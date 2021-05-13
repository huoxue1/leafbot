package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func UseEchoHandle() {

	leafBot.
		OnCommand("/echo").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(args[0]))
			})

}
