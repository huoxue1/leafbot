package cqhttp_http_driver

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/guonaihong/gout"
	log "github.com/sirupsen/logrus"
)

// Driver
// @Description:
//
type Driver struct {
	Name    string
	webHook []struct {
		postHost string
		postPort int
		selfID   int64
	}
	token            string
	listenHost       string
	listenPort       int
	bots             sync.Map
	eventChan        chan []byte
	connectHandle    func(selfId int64, host string, clientRole string)
	disConnectHandle func(selfId int64)
}

func (d *Driver) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	d.eventChan <- data
	writer.WriteHeader(200)
}

//SetToken
/**
 * @Description:
 * @receiver d
 * @param token
 */
func (d *Driver) SetToken(token string) {
	d.token = token
}

//Run
/**
 * @Description:
 * @receiver d
 */
func (d *Driver) Run() {
	for _, s := range d.webHook {
		b := new(Bot)
		b.selfId = s.selfID
		b.postHost = s.postHost
		b.postPort = s.postPort
		b.responses = sync.Map{}
		b.disConnectHandle = d.disConnectHandle
		b.client = gout.NewWithOpt()
		b.token = d.token
		d.bots.Store(s.selfID, b)
	}

	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", d.listenHost, d.listenPort), d); err != nil {
		log.Errorln("监听webhook失败" + err.Error())
	}
}

//GetEvent
/**
 * @Description: 获取事件信息通道
 * @receiver d
 * @return chan
 */
func (d *Driver) GetEvent() chan []byte {
	return d.eventChan
}

//GetBot
/**
 * @Description: 获取一个bot对象
 * @receiver d
 * @param i
 * @return interface{}
 */
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
 * @Description: 注册一个bot断开时的钩子
 * @receiver d
 * @param f
 * example
 */
func (d *Driver) OnDisConnect(f func(selfId int64)) {
	d.disConnectHandle = f
}

// GetBots
/**
 * @Description: 获取一个bot对象
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

//SetConfig
/**
 * @Description: 设置配置信息
 * @receiver d
 * @param config
 */
func (d *Driver) SetConfig(config map[string]interface{}) {
	if host, ok := config["listen_host"]; ok {
		d.listenHost = host.(string)
	}
	if port, ok := config["listen_port"]; ok {
		d.listenPort = port.(int)
	}
}

//AddWebHook
/**
 * @Description: 添加一个webHook
 * @receiver d
 * @param selfID
 * @param postHost
 * @param postPort
 */
func (d *Driver) AddWebHook(selfID int64, postHost string, postPort int) {
	d.webHook = append(d.webHook, struct {
		postHost string
		postPort int
		selfID   int64
	}{postHost: postHost, postPort: postPort, selfID: selfID})
}

//NewDriver
/**
 * @Description: 创建一个cqhttp的http通信方式驱动
 * @return *Driver
 */
func NewDriver() *Driver {
	d := new(Driver)
	d.Name = "cqhttp"
	d.bots = sync.Map{}
	d.eventChan = make(chan []byte)
	return d
}
