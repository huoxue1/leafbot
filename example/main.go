package main

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_positive_ws_driver"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	leafBot.NewPlugin("测试").OnCommand("测试").AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		event.Send(message.Text("测试"))
	})
}

func main() {
	// 创建一个驱动
	driver := cqhttp_positive_ws_driver.NewDriver()
	// 注册驱动
	leafBot.LoadDriver(driver)
	// 初始化Bot
	leafBot.InitBots()
	// 运行驱动
	driver.Run()
	select {}
}
