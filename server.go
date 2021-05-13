package leafBot

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type connection struct {
	wsSocket *websocket.Conn  // 底层websocket
	InChan   chan []byte      // 读队列
	OutChan  chan interface{} // 写队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

// GetBot
/*
   @Description:
   @param name string
   @return *Bot
*/
func GetBot(name string) *Bot {
	for _, bot := range DefaultConfig.Bots {
		if bot.Name == name {
			return bot
		}
	}
	return nil
}

// readData
/*
   @Description:
   @receiver con
*/
func (con *connection) readData() {
	go func() {
		for {
			time.Sleep(10)
			_, data, err := con.wsSocket.ReadMessage()
			if err != nil {
				con.wsClose()
			}
			select {
			case con.InChan <- data:
			case <-con.closeChan:
				return
			}

		}
	}()

}

// FilterEventOrResponse
/*
   @Description:
   @receiver con
*/
func (con *connection) FilterEventOrResponse() {
	for true {
		time.Sleep(10)
		data := <-con.InChan
		//fmt.Println(string(data))
		err := json.Unmarshal(data, new(Event))
		if err != nil {
			apiResChan <- data
		} else {
			eventChan <- data
		}
	}

}

// writeData
/*
   @Description:
   @receiver con
*/
func (con *connection) writeData() {
	go func() {
		for {
			time.Sleep(10)
			select {
			case data := <-con.OutChan:
				if err := con.wsSocket.WriteJSON(data); err != nil {
					con.wsClose()
				}
			case <-con.closeChan:
				return
			}

		}
	}()

}

// wsClose
/*
   @Description:
   @receiver con
*/
func (con *connection) wsClose() {
	log.Debugln("链接已关闭")
	_ = con.wsSocket.Close()

	con.mutex.Lock()
	defer con.mutex.Unlock()
	if !con.isClosed {
		con.isClosed = true
		close(con.closeChan)
	}
}

// EventHandle
/*
   @Description:
   @param w http.ResponseWriter
   @param r *http.Request
*/
func EventHandle(w http.ResponseWriter, r *http.Request) {
	selfId, err := strconv.Atoi(r.Header.Get("X-Self-ID"))

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wscon := &connection{
		wsSocket:  conn,
		InChan:    make(chan []byte, 10),
		OutChan:   make(chan interface{}, 10),
		isClosed:  false,
		closeChan: make(chan byte),
	}
	for _, bot := range DefaultConfig.Bots {
		if bot.SelfId == selfId {
			log.Infoln("bot：" + bot.Name + "已上线")
			bot.Client = wscon
		}
	}
	go wscon.readData()
	go wscon.writeData()
	go wscon.FilterEventOrResponse()
}

// getResponse
/*
   @Description:
   @param r *response
   @param echo string
   @return []byte
*/
func getResponse(r *response, echo string) []byte {
	for true {
		data := <-apiResChan
		err := json.Unmarshal(data, r)
		if err != nil {
			return []byte{}
		}
		if r.Echo == "" {
			continue
		}
		if r.Echo == echo {
			return data
		} else {
			apiResChan <- data
		}

	}
	return []byte{}
}
