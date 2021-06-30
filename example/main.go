package main

import (
	"flag" //nolint:gci
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/3343780376/leafBot/plugins"
	"github.com/3343780376/leafBot/plugins/autoReply"
	"github.com/3343780376/leafBot/plugins/blacklist"
	"github.com/3343780376/leafBot/plugins/groupWelcome"
	"github.com/3343780376/leafBot/plugins/manager"
	_ "github.com/3343780376/leafBot/plugins/poke"
	"github.com/3343780376/leafBot/plugins/searchImage"
	_ "github.com/3343780376/leafBot/plugins/weibo"
	"os" //nolint:gci
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		go leafBot.InitWindow()
	}
	//为bot添加weather响应器，命令为 ”/天气“ ,allies为命令别名，
	//参数格式为一个字符串数组，rule为一个结构体，响应前会先判断所以rules为true，weight为权重，block为是否阻断
	manager.InitBanPlugin()
	leafBot.OnCommand("/天气").
		SetWeight(10).
		SetPluginName("天气").
		SetBlock(false).
		AddHandle(Weather)
	plugins.Ocr()
	groupWelcome.WelcomeInit()
	leafBot.InitPluginManager()
	searchImage.InitImage()
	plugins.UseCreateQrCode()      //加载生成二维码插件
	plugins.UseDayImage()          // 加载每日一图插件
	plugins.UseEchoHandle()        // 加载echo插件
	plugins.UseMusicHandle()       // 加载音乐插件
	plugins.UseSetuHandle()        // 加载涩图插件
	plugins.UseTranslateHandle()   // 加载翻译插件
	plugins.UseFlashImage(0)       // 加载闪照破解插件
	plugins.UseFlashImageToGroup() //加载闪照破解后发到对应群的插件

	blacklist.InitBlackList("./config/blackList.json") //加载黑名单插件
	_ = autoReply.Load("./config/data.json")
	//加载自动回复插件
}

func main() {
	dir, _ := os.Getwd() // 获取当前路径

	leafBot.LoadConfig(dir+"/config/config.json", leafBot.JSON)

	var port int
	if len(os.Args) > 1 {
		flag.IntVar(&port, "port", leafBot.DefaultConfig.Port, "端口")
		flag.Parse()
		if port != leafBot.DefaultConfig.Port {
			leafBot.DefaultConfig.Port = port
		}

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
		message.Text(args[0]+"的天气为"+m[args[0]]))
}
