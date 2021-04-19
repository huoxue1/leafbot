package plugins

import "github.com/3343780376/leafBot"

func UseEchoHandle() {

	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, args[0], false)
	}, "/echo", nil, nil, 1, false)
}
