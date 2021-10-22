package cqhttp_positive_ws_driver

import (
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Driver struct {
	Name             string
	address          string
	port             int
	bots             sync.Map
	eventChan        chan []byte
	connectHandle    func(selfId int64, host string, clientRole string)
	disConnectHandle func(selfId int64)
}

func (d *Driver) Run() {
	u := url.URL{Scheme: "ws", Host: d.address + ":" + strconv.Itoa(d.port)}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return
	}
	_, data, err := conn.ReadMessage()
	if err != nil {
		return
	}
	selfId := gjson.GetBytes(data, "self_id").Int()
	role := ""
	host := d.address

	b := new(Bot)
	b.conn = conn
	b.selfId = selfId
	b.responses = sync.Map{}

	_, ok := d.bots.Load(selfId)
	if ok {
		d.bots.LoadOrStore(selfId, b)
	} else {
		d.bots.Store(selfId, b)
	}

	d.connectHandle(selfId, host, role)
	b.disConnectHandle = d.disConnectHandle
	log.Infoln(fmt.Sprintf("the bot %v is connected", selfId))
	go func() {
		defer func() {
			i := recover()
			if i != nil {
				log.Errorln("ws链接读取出现错误")
				log.Errorln(i)
				d.disConnectHandle(selfId)
			}
		}()
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				b.wsClose()
			}

			echo := gjson.GetBytes(data, "echo")
			if echo.Exists() {
				b.responses.Store(echo.String(), data)
			} else {
				d.eventChan <- data
			}
		}
	}()
}

func (d *Driver) GetEvent() chan []byte {
	return d.eventChan
}

func (d *Driver) GetBot(i int64) interface{} {
	load, ok := d.bots.Load(i)
	if ok {
		return load
	}

	return nil
}

// OnConnect
/**
 * @Description:
 * @receiver d
 * @param f
 * example
 */
func (d *Driver) OnConnect(f func(selfId int64, host string, clientRole string)) {
	d.connectHandle = f
}

// OnDisConnect
/**
 * @Description:
 * @receiver d
 * @param f
 * example
 */
func (d *Driver) OnDisConnect(f func(selfId int64)) {
	d.disConnectHandle = f
}

// GetBots
/**
 * @Description:
 * @receiver d
 * @return map[int64]interface{}
 * example
 */
func (d *Driver) GetBots() map[int64]interface{} {
	m := make(map[int64]interface{})
	d.bots.Range(func(key, value interface{}) bool {
		m[key.(int64)] = value
		return true
	})

	return m
}

// SetAddress
/**
 * @Description:
 * @receiver d
 * @param string2
 * example
 */
func (d *Driver) SetAddress(string2 string) {
	d.address = string2
}

func (d *Driver) SetPort(port int) {
	d.port = port
}

func NewDriver() *Driver {
	d := new(Driver)
	d.Name = "cqhttp"
	d.bots = sync.Map{}
	d.eventChan = make(chan []byte)
	return d
}
