// Package groupWelcome /**
package groupWelcome

import "github.com/huoxue1/leafBot"

// WelcomeInit
/**
 * @Description:
 */
func WelcomeInit() {
	leafBot.OnNotice(leafBot.NoticeTypeApi.GroupIncrease).SetWeight(10).SetPluginName("入群欢迎").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		for _, s := range leafBot.DefaultConfig.Plugins.Welcome {
			if s.GroupId == event.GroupId {
				bot.SendGroupMsg(event.GroupId, s.Message)
			}
		}
	})
}
