package leafBot

func UseEchoHandle() {

	AddCommandHandle(func(event Event, bot *Bot, args []string) {
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, args[0], false)
	}, "/echo", nil, nil, 10, false)
}
