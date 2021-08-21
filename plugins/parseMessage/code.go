package parsemessage

import "github.com/huoxue1/leafBot"

// Init
/**
 * @Description:
 * example
 */
func Init() {
	leafBot.OnMessage("").SetWeight(10).SetPluginName("特殊消息解析").AddRule(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) bool {
		if event.Message[0].Type == "reply" {
			for _, messageSegment := range event.Message {
				if messageSegment.Type == "text" && messageSegment.Data["text"] == "解析" {
					return true
				}
			}
		}
		return false
	}).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {

	})
}
