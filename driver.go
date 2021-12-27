package leafbot

import (
	log "github.com/sirupsen/logrus"
)

//Driver
// @Description: 驱动器接口
//
type Driver interface {
	// Run
	// @Description: 运行该驱动的接口，该接口应该为阻塞式运行
	//
	Run()
	// GetEvent
	// @Description: 返回一个chan，该chan为事件传递的chan
	// @return chan
	//
	GetEvent() chan []byte

	OnConnect(func(selfId int64, host string, clientRole string))
	OnDisConnect(func(selfId int64))

	// GetBot
	// @Description: 获取一个实现了APi接口的bot
	// @param int64 bot的id
	// @return interface{}
	//
	GetBot(int64) interface{}
	// GetBots
	// @Description: 获取所有bot
	// @return map[int64]interface{}
	//
	GetBots() map[int64]interface{}

	SetConfig(config map[string]interface{})

	AddWebHook(selfID int64, postHost string, postPort int)

	SetToken(token string)
}

//Conn
// @Description:
//
type Conn interface {
	// Do
	// @Description: 执行一个api
	// @param interface{}
	//
	Do(interface{})
	// GetResponse
	// @Description: 获取一次api的执行结果
	// @param echo 标识一次执行的唯一参数
	// @return []byte 响应结果
	// @return error 超时会返回一个error
	//
	GetResponse(echo string) ([]byte, error)
	GetSelfId() int64
}

//  leafBot所注册的驱动
var driver Driver

//LoadDriver
/**
 * @Description: 为leafBot注册一个驱动
 * @param driver2 实现了Driver接口的驱动
 * example
 */
func LoadDriver(driver2 Driver) {
	driver2.SetConfig(map[string]interface{}{
		"host": defaultConfig.Host,
		"port": defaultConfig.Port,

		"listen_host": defaultConfig.ListenHost,
		"listen_port": defaultConfig.ListenPort,
	})
	driver2.SetToken(defaultConfig.Token)
	for _, s := range defaultConfig.WebHook {
		driver2.AddWebHook(s.SelfID, s.PostHost, s.PostPort)
	}
	driver2.OnConnect(func(selfId int64, host string, clientRole string) {
		defer func() {
			err := recover()
			if err != nil {
				log.Errorln("执行连接回调出现不可预料的错误")
				log.Errorln(err)
			}
		}()
		//for _, handle := range ConnectHandles {
		//	go handle.handle(Connect{
		//		SelfID:     selfId,
		//		Host:       host,
		//		ClientRole: clientRole,
		//	}, driver2.GetBot(selfId).(API))
		//}
	})
	driver2.OnDisConnect(func(selfId int64) {
		defer func() {
			err := recover()
			if err != nil {
				log.Errorln("执行连接断开回调出现不可预料的错误")
				log.Errorln(err)
			}
		}()
		//for _, handle := range DisConnectHandles {
		//	go handle.handle(selfId)
		//}
	})
	driver = driver2
}
