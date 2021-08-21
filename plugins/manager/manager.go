package manager

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/message"
	"strconv"
)

func init() {
	Init()
}
func Init() {
	leafBot.OnRegex(`^升为管理.*?qq=(\d+)`).
		SetPluginName("群管系统-设置管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupAdmin(event.GroupId, ID, true)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"升为了管理！"))
			})

	leafBot.OnRegex(`^取消管理.*?qq=(\d+)`).
		SetPluginName("群管系统-取消管理").
		SetBlock(false).
		AddRule(leafBot.OnlySuperUser).
		SetWeight(10).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
				ID, err := strconv.Atoi(state.RegexResult[1])
				if err != nil {
					bot.Send(event, message.Text("发生未知错误"+err.Error()))
					return
				}
				bot.SetGroupAdmin(event.GroupId, ID, false)
				nickName := bot.GetGroupMemberInfo(event.GroupId, ID, true).NickName
				bot.Send(event, message.Text(nickName+"升为了管理！"))
			},
		)

}
