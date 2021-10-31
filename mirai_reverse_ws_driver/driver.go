package mirai_reverse_ws_driver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Driver struct {
	Name      string
	host      string
	port      int
	conn      *websocket.Conn
	eventChan chan []byte

	sessionKey string
	verifyKey  string
}

// ws协议
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func (d *Driver) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrade.Upgrade(writer, request, nil)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	d.conn = conn
}

func (d *Driver) Run() {
	http.Handle("/"+d.Name+"/ws", d)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", d.host, d.port), nil); err != nil {
		log.Panicln(err.Error())
	}
}

func (d *Driver) GetEvent() chan []byte {
	panic("implement me")
}

func (d *Driver) OnConnect(f func(selfId int64, host string, clientRole string)) {
	panic("implement me")
}

func (d *Driver) OnDisConnect(f func(selfId int64)) {
	panic("implement me")
}

func (d *Driver) GetBot(i int64) interface{} {
	panic("implement me")
}

func (d *Driver) GetBots() map[int64]interface{} {
	panic("implement me")
}

func (d *Driver) SetConfig(config map[string]interface{}) {
	panic("implement me")
}

func (d *Driver) AddWebHook(selfID int64, postHost string, postPort int) {
	panic("implement me")
}

func (d *Driver) SetToken(token string) {
	panic("implement me")
}
