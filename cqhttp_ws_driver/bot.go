package cqhttp_ws_driver

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Bot struct {
	selfId    int64
	conn      *websocket.Conn
	responses sync.Map
	lock      sync.Mutex
}

func (b *Bot) GetSelfId() int64 {
	return b.selfId
}

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
	defer b.lock.Unlock()

}
