/**
2 * @Author :goujiangshan
3 * @DATA :  16:22
4 */

package leafBot

import ( //nolint:gci

	"embed"
	"github.com/3343780376/leafBot/message"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/huoxue1/lorca"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http" //nolint:gci
	"os"
	"os/signal"
	"sort"
	"strconv"
)

//go:embed gui/static
var static embed.FS

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWindow() {
	defer func() {
		err := recover()
		log.Infoln(err)
	}()
	go func() {
		ui, err := lorca.New("http://127.0.0.1:3000/static/gui/static/html/default.html", "", 800, 600)
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c) //nolint:govet
			for {
				log.Infoln(<-c)
				ui.Close()
			}
		}()
		defer ui.Close()
		if err != nil {
			log.Panic(err)
		}
		<-ui.Done()
		os.Exit(3)
	}()
	if DefaultConfig.LogLevel != "debug" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	engine := gin.New()

	engine.StaticFS("/static", http.FS(static))
	//engine.StaticFile("/", "./gui/view/static/html/default.html")
	//engine.LoadHTMLGlob("./gui/view/html/*.html")
	//engine.GET("/", func(context *gin.Context) {
	//	context.HTML(200, "default.html", nil)
	//})

	engine.POST("/get_config", GetConfig)
	engine.POST("/get_group_list", GetGroupList)
	engine.POST("/get_friend_list", GetFriendList)

	engine.POST("/update_plugin_states", func(context *gin.Context) {
		id := context.PostForm("id")
		status, err := strconv.ParseBool(context.PostForm("status"))
		if err != nil {
			log.Errorln("改变插件状态出错" + err.Error())
		}
		if status {
			StartPluginByID(id)
		} else {
			BanPluginByID(id)
		}
		context.JSON(200, nil)
	})

	engine.POST("/get_plugins", func(context *gin.Context) {
		list := GetHandleList()
		var pluginList []BaseHandle

		for _, handles := range list {
			pluginList = append(pluginList, handles...)
		}
		sort.SliceStable(pluginList, func(i, j int) bool {
			id1, _ := strconv.Atoi(pluginList[i].ID)
			id2, _ := strconv.Atoi(pluginList[j].ID)
			return id1 < id2
		})
		context.JSON(200, pluginList)
	})

	engine.GET("/get_log", func(context *gin.Context) {
		conn, err := upGrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Errorln(err)
			return
		}
		go func() {
			for {
				event := <-hook.LogChan
				err = conn.WriteMessage(websocket.TextMessage, []byte(event))
				if err != nil {
					log.Debugln("前端日志消息发送失败" + err.Error())
					continue
				}
			}
		}()
	})

	engine.POST("/send_msg", CallApi)
	engine.GET("/data", data)
	if err := engine.Run(":3000"); err != nil {
		log.Debugln(err.Error())
	}

}

func data(ctx *gin.Context) {
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Infoln(err)
		return
	}

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Debugln("接收消息失败" + err.Error())
				break
			}
			if string(message) == "ping" {
				message = []byte("pong")
			}
			//写入ws数据
		}
	}()

	go func() {
		for {
			event := <-MessageChan
			log.Debugln("已向前端发送信息")
			err = conn.WriteJSON(&event)
			if err != nil {
				log.Debugln("消息发送失败" + err.Error())
				continue
			}
		}

	}()

}

func GetConfig(ctx *gin.Context) {
	var bots []int
	for _, bot := range DefaultConfig.Bots {
		bots = append(bots, bot.SelfId)
	}
	ctx.JSON(200, bots)
}

func GetGroupList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := GetBotById(selfID)
	list := bot.GetGroupList()
	ctx.JSON(200, list)
}

func GetFriendList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := GetBotById(selfID)
	list := bot.GetFriendList()
	ctx.JSON(200, list)
}

func CallApi(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	id, err := strconv.Atoi(ctx.PostForm("id"))
	message1 := ctx.PostForm("message")
	messageType := ctx.PostForm("message_type")
	if err != nil {
		ctx.JSON(404, nil)
	}
	bot := GetBotById(selfID)
	msgID := bot.SendMsg(messageType, id, id, message.ParseMessageFromString(message1))
	ctx.JSON(200, msgID)
}
