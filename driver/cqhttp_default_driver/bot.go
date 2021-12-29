package cqhttp_default_driver

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/api"
	"github.com/tidwall/gjson"
)

type userAPi struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

func (u userAPi) Get(s string) gjson.Result {
	data, _ := json.Marshal(u.Params)
	parse := gjson.Parse(string(data))
	return parse.Get(s)
}

type Bot struct {
	responses sync.Map
	CQBot     *coolq.CQBot
	call      *api.Caller
}

func (b *Bot) Do(i interface{}) {
	data := i.(userAPi)
	call := b.call.Call(data.Action, data)
	resp, _ := json.Marshal(call)
	b.responses.Store(data.Echo, resp)
}

func (b *Bot) GetResponse(echo string) ([]byte, error) {
	defer func() {
		b.responses.Delete(echo)
	}()

	for i := 0; i < 120; i++ {
		value, ok := b.responses.LoadAndDelete(echo)
		if ok {
			return value.([]byte), nil
		}
		time.Sleep(500)
	}

	return nil, errors.New("get response time out")
}

func (b *Bot) GetSelfId() int64 {
	return b.CQBot.Client.Uin
}
