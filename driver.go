package leafBot

import "fmt"

// Driver
// @Description: 实现了该接口即可注册为leafBot的driver
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
	// AddContentHandle
	// @Description: 为驱动注册链接事件
	//
	AddContentHandle()
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
	// SetAddress
	// @Description: 设置driver的运行地址
	// @param string2
	//
	SetAddress(string2 string)
}

// Conn
// @Description: 所有bot对象都应该实现该接口
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

// LoadDriver
/**
 * @Description: 为leafBot注册一个驱动
 * @param driver2 实现了Driver接口的驱动
 * example
 */
func LoadDriver(driver2 Driver) {
	driver2.SetAddress(fmt.Sprintf(":%v", DefaultConfig.Port))
	driver = driver2
}
