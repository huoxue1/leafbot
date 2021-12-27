package main

import (
	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/cqhttp_reverse_ws_driver"
	"github.com/huoxue1/leafbot/message"
)

func init() {
	leafbot.NewPlugin("测试").OnCommand("测试", leafbot.Option{
		Weight: 0,
		Block:  false,
		Allies: nil,
		Rules: []leafbot.Rule{func(ctx *leafbot.Context) bool {
			return true
		}},
	}).Handle(func(ctx *leafbot.Context) {
		ctx.Send(message.Text("123"))
	})
}

func main() {
	// 创建一个驱动
	driver := cqhttp_reverse_ws_driver.NewDriver()
	// 注册驱动
	leafbot.LoadDriver(driver)
	// 初始化Bot
	leafbot.InitBots()
	// 运行驱动
	driver.Run()
	select {}
}
