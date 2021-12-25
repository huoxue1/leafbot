package main

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_reverse_ws_driver"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	plugin := leafBot.NewPlugin("echo")
	plugin.SetHelp(map[string]string{"echo": "echo the params"})
	plugin.OnCommand("echo").AddHandle(func(event leafBot.Event, bot leafBot.Api, state *leafBot.State) {
		event.Send(message.Text(state.Args[:]))
	})
}

func main() {
	driver := cqhttp_reverse_ws_driver.NewDriver()
	leafBot.LoadDriver(driver)
	leafBot.InitBots()
	driver.Run()
}
