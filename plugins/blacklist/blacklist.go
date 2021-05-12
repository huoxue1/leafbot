package blacklist

import "github.com/3343780376/leafBot"

type blackList struct {
	users  []int
	groups []int
}

var (
	BlackList = blackList{}
)

func InitBlackList() {
	leafBot.AddPretreatmentHandle(nil, 10, func(event leafBot.Event, bot *leafBot.Bot) bool {
		for _, user := range BlackList.users {
			if user == event.UserId {
				return false
			}
		}
		for _, group := range BlackList.groups {
			if event.GroupId == group {
				return false
			}
		}
		return true
	})
}
