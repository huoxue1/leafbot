package leafBot

import "strconv"

type (
	Rule func(Event, *Bot) bool
)

func OnlyToMe(event Event, bot *Bot) bool {
	if event.MessageType == "private" {
		return true
	}
	msg := event.GetMsg()
	for _, segment := range msg {
		if segment.Type == "at" && segment.Data["qq"] == strconv.Itoa(event.SelfId) {
			return true
		}
	}

	return false
}

func OnlySuperUser(event Event, bot *Bot) bool {
	if event.UserId == DefaultConfig.Admin {
		return true
	}
	for _, user := range DefaultConfig.SuperUser {
		if event.UserId == user {
			return true
		}
	}
	return false
}

func OnlyGroupMessage(event Event, bot *Bot) bool {

	return event.MessageType == "group"
}
