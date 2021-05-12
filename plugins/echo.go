package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func UseEchoHandle() {

	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(args[0]))
	}, "/echo", nil, nil, 1, false)
}
