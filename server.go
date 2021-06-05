package leafBot

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket" //nolint:gci
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	//nolint:gci
)

type connection struct {
	SelfID   int
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

// readData
/*
   @Description:
   @receiver con
*/
func (con *connection) readData() {
	go func() {
		for {
			_, data, err := con.wsSocket.ReadMessage()
			log.Debugln(string(data))
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
	for {
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
	log.Infoln(fmt.Sprintf("bot%v已断开链接", con.SelfID))
	for _, handle := range DisConnectHandles {
		handle.handle(con.SelfID)
	}
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
	role := r.Header.Get("X-Client-Role")
	host := r.Header.Get("Host")

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wscon := &connection{
		SelfID:    selfId,
		wsSocket:  conn,
		InChan:    make(chan []byte, 10),
		OutChan:   make(chan interface{}, 10),
		isClosed:  false,
		closeChan: make(chan byte),
	}

	// 将所有bot实例的client对象初始化
	for _, bot := range DefaultConfig.Bots {
		if bot.SelfId == selfId {
			log.Infoln("bot：" + strconv.Itoa(bot.SelfId) + "已上线")
			bot.Client = wscon
		}
	}
	// 处理所有的connect事件
	for _, handle := range ConnectHandles {
		handle.handle(Connect{SelfID: selfId, ClientRole: role, Host: host}, GetBotById(selfId))
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
