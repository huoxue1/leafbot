package cqhttp_http_driver

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/guonaihong/gout"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	client           *gout.Client
	selfId           int64
	responses        sync.Map
	lock             sync.Mutex
	disConnectHandle func(selfId int64)
	postHost         string
	postPort         int
}

// GetSelfId
/**
 * @Description: 获取bot的账号
 * @receiver b
 * @return int64
 * example
 */
func (b *Bot) GetSelfId() int64 {
	return b.selfId
}

// Do
/**
 * @Description: 执行一个api的调用
 * @receiver b
 * @param i
 * example
 */
func (b *Bot) Do(i interface{}) {
	data := i.(UseApi)
	var resp []byte
	err := b.client.POST(fmt.Sprintf("http://%v:%v/%v", b.postHost, b.postPort, data.Action)).SetJSON(data.Params).BindBody(&resp).Do()
	if err != nil {
		log.Errorln("调用api出现错误", err.Error())
		return
	}
	b.responses.Store(data.Echo, resp)
}

// GetResponse
/**
 * @Description: 获取一个api调用的响应
 * @receiver b
 * @param echo api调用的唯一标识
 * @return []byte
 * @return error
 * example
 */
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

func (b *Bot) wsClose() {
	// todo
}
