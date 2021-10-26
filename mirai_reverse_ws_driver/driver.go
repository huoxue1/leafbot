package mirai_reverse_ws_driver

import "github.com/gorilla/websocket"

type Driver struct {
	host      string
	port      int
	conn      *websocket.Conn
	eventChan chan []byte

	sessionKey string
	verifyKey  string
}

func (d Driver) Run() {
	panic("implement me")
}

func (d Driver) GetEvent() chan []byte {
	panic("implement me")
}

func (d Driver) OnConnect(f func(selfId int64, host string, clientRole string)) {
	panic("implement me")
}

func (d Driver) OnDisConnect(f func(selfId int64)) {
	panic("implement me")
}

func (d Driver) GetBot(i int64) interface{} {
	panic("implement me")
}

func (d Driver) GetBots() map[int64]interface{} {
	panic("implement me")
}

func (d Driver) SetConfig(config map[string]interface{}) {
	panic("implement me")
}

func (d Driver) AddWebHook(selfID int64, postHost string, postPort int) {
	panic("implement me")
}

func (d Driver) SetToken(token string) {
	panic("implement me")
}
