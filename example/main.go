package main

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/plugins"
	"os"
)

func init() {
	// 为bot添加weather响应器，命令为 ”/天气“ ,allies为命令别名，
	//参数格式为一个字符串数组，rule为一个结构体，响应前会先判断所以rules为true，weight为权重，block为是否阻断
	plugins.UseDayImage()
	plugins.UseEchoHandle()
	plugins.UseMusicHandle()
	plugins.UseSetuHandle()
	plugins.UseTranslateHandle()
	plugins.UseFlashImage(3343780376)
	leafBot.AddCommandHandle(Weather, "/天气", nil, nil, 10, false)
}

func main() {
	dir, _ := os.Getwd() // 获取当前路径
	if len(os.Args) > 1 {
		leafBot.LoadConfig(os.Args[1])
	} else {
		leafBot.LoadConfig(dir + "/example/config.json")
	}

	//拼接配置文件路径，并且加载配置文件
	leafBot.InitBots() //初始化Bot
}

/*
	event: bot的event，里面包含了事件的所有字段
	bot: 触发事件的bot指针
	args ： 命令的参数，为一个数组
*/
func Weather(event leafBot.Event, bot *leafBot.Bot, args []string) {
	m := map[string]string{"北京": "晴", "山东": "下雨"}
	// 调用发送消息的api，会根据messageType自动回复
	bot.SendMsg(event.MessageType, event.UserId, event.GroupId,
		args[0]+"的天气为"+m[args[0]],
		false)
}
