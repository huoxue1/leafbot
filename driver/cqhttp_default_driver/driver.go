package cqhttp_default_driver

import (
	"github.com/Mrs4s/go-cqhttp/cmd/gocq"
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/api"
	"github.com/Mrs4s/go-cqhttp/modules/servers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/huoxue1/leafbot/message"
)

type Driver struct {
	coolq.CQBot
	EventChan chan []byte
	bot       *Bot
}

func (d *Driver) Run() {
	servers.RegisterCustom("leafBot", func(bot *coolq.CQBot) {
		b := new(Bot)
		b.CQBot = bot
		b.call = api.NewCaller(bot)
		d.bot = b
		bot.OnEventPush(func(e *coolq.Event) {
			data := e.JSONString()
			result := gjson.Parse(data)
			if result.Get("message").Exists() {
				m := message.ParseMessageFromString(result.Get("message").String())
				data, _ = sjson.Set(data, "message", m)
			}
			d.EventChan <- []byte(data)
		})
	})
	gocq.Main()
}

func (d *Driver) GetEvent() chan []byte {
	return d.EventChan
}

func (d *Driver) OnConnect(f func(selfId int64, host string, clientRole string)) {
}

func (d *Driver) OnDisConnect(f func(selfId int64)) {

}

func (d *Driver) GetBot(i int64) interface{} {
	return d.bot
}

func (d *Driver) GetBots() map[int64]interface{} {
	return map[int64]interface{}{d.CQBot.Client.Uin: d.bot}
}

func (d *Driver) SetConfig(config map[string]interface{}) {

}

func (d *Driver) AddWebHook(selfID int64, postHost string, postPort int) {

}

func (d *Driver) SetToken(token string) {

}

func NewDriver() *Driver {
	d := new(Driver)
	d.EventChan = make(chan []byte)
	return d
}
