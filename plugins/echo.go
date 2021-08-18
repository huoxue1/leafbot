package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func UseEchoHandle() {

	leafBot.
		OnCommand("/echo").
		SetPluginName("echo").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state leafBot.State) {
				bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(state.Args[0]))
			})

}
