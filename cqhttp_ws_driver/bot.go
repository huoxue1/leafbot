package cqhttp_ws_driver

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Bot
// @Description: bot实例对象
//
type Bot struct {
	selfId           int64
	conn             *websocket.Conn
	responses        sync.Map
	lock             sync.Mutex
	disConnectHandle func(selfId int64)
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
	err := b.conn.WriteJSON(i)
	if err != nil {
		b.wsClose()
		return
	}
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

	c := make(chan []byte, 1)
	b.responses.Store(echo, c)
	after := time.After(60 * 100000)
	select {
	case data := <-c:
		return data, nil
	case <-after:
		return nil, errors.New("")
	}
}

func (b *Bot) wsClose() {
	_ = b.conn.Close()
	b.lock.Lock()
	b.disConnectHandle(b.selfId)
	defer b.lock.Unlock()
}
