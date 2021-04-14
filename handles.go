package leafBot

func UseEchoHandle(name string) {

	AddCommandHandle(func(event Event, args []string) {
		bot := GetBot("commit")
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, args[0], false)

	}, ".echo", nil, nil, 10, false)
}
