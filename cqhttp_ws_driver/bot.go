package cqhttp_ws_driver

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Bot
// @Description:
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
 * @Description:
 * @receiver b
 * @return int64
 * example
 */
func (b *Bot) GetSelfId() int64 {
	return b.selfId
}

// Do
/**
 * @Description:
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
