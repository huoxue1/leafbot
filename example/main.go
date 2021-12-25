package main

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_reverse_ws_driver"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	leafBot.NewPlugin("测试").OnCommand("测试", leafBot.Option{
		PluginName: "测试",
		Weight:     0,
		Block:      false,
		Allies:     nil,
		Rules: []leafBot.Rule{func(ctx *leafBot.Context) bool {
			return true
		}},
		CD: leafBot.CoolDown{},
	}).AddHandle(func(ctx *leafBot.Context) {
		ctx.Send(message.Text("123"))
	})
}

func main() {
	// 创建一个驱动
	driver := cqhttp_reverse_ws_driver.NewDriver()
	// 注册驱动
	leafBot.LoadDriver(driver)
	// 初始化Bot
	leafBot.InitBots()
	// 运行驱动
	driver.Run()
	select {}
}
