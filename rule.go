package leafBot

type (
	Rule func(Event, *Bot) bool
)
