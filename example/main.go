package main

import (
	//nolint:gci
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_ws_driver"
	//nolint:gci
)

func main() {
	// 创建一个驱动
	driver := cqhttp_ws_driver.NewDriver()
	// 注册驱动
	leafBot.LoadDriver(driver)
	//初始化Bot
	leafBot.InitBots()
	// 运行驱动
	driver.Run()
}
