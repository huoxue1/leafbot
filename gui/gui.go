/**
2 * @Author :goujiangshan
3 * @DATA :  16:22
4 */

package gui

import ( //nolint:gci
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/zserge/lorca"
	"net/http" //nolint:gci
	"os"
	"os/signal"
	"sort"
	"strconv"
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	},
	)
}

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
		ui, err := lorca.New("http://127.0.0.1:3000", "", 800, 600)
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
	engine := gin.New()
	gin.SetMode(gin.ReleaseMode)
	engine.StaticFS("/static/", http.Dir("./gui/view/static/"))
	engine.StaticFile("/", "./gui/view/html/index.html")
	engine.LoadHTMLGlob("./gui/view/html/*.html")
	//engine.GET("/", func(context *gin.Context) {
	//	context.HTML(200, "index.html", nil)
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
			leafBot.StartPluginByID(id)
		} else {
			leafBot.BanPluginByID(id)
		}
		context.JSON(200, nil)
	})

	engine.POST("/get_plugins", func(context *gin.Context) {
		list := leafBot.GetHandleList()
		var pluginList []leafBot.BaseHandle

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
			event := <-leafBot.MessageChan
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
	for _, bot := range leafBot.DefaultConfig.Bots {
		bots = append(bots, bot.SelfId)
	}
	ctx.JSON(200, bots)
}

func GetGroupList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := leafBot.GetBotById(selfID)
	list := bot.GetGroupList()
	ctx.JSON(200, list)
}

func GetFriendList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := leafBot.GetBotById(selfID)
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
	bot := leafBot.GetBotById(selfID)
	msgID := bot.SendMsg(messageType, id, id, message.ParseMessageFromString(message1))
	ctx.JSON(200, msgID)
}
