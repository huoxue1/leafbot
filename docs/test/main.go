package main

import (
	leafbot "github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_reverse_ws_driver"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	plugin := leafbot.NewPlugin("echo")
	plugin.SetHelp(map[string]string{"echo": "echo the params"})
	plugin.OnCommand("echo").AddHandle(func(ctx *leafbot.Context) {
		ctx.Send(message.Text(ctx.State.Args[:]))
	})
}

func main() {
	driver := cqhttp_reverse_ws_driver.NewDriver()
	leafbot.LoadDriver(driver)
	leafbot.InitBots()
	driver.Run()
}
