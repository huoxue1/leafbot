/**
2 * @Author :goujiangshan
3 * @DATA :  16:22
4 */

package gui

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/zserge/lorca"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

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
			signal.Notify(c)
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
	engine.LoadHTMLGlob("./gui/view/html/*.html")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	engine.POST("/get_config", GetConfig)
	engine.POST("/get_group_list", GetGroupList)
	engine.POST("/get_friend_list", GetFriendList)

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
	selfId, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := leafBot.GetBotById(selfId)
	list := bot.GetGroupList()
	ctx.JSON(200, list)
}

func GetFriendList(ctx *gin.Context) {
	selfId, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		return
	}
	bot := leafBot.GetBotById(selfId)
	list := bot.GetFriendList()
	ctx.JSON(200, list)
}

func CallApi(ctx *gin.Context) {
	selfId, err := strconv.Atoi(ctx.PostForm("self_id"))
	id, err := strconv.Atoi(ctx.PostForm("id"))
	message1 := ctx.PostForm("message")
	messageType := ctx.PostForm("message_type")
	if err != nil {
		ctx.JSON(404, nil)
	}
	bot := leafBot.GetBotById(selfId)
	msgId := bot.SendMsg(messageType, id, id, message.ParseMessageFromString(message1))
	ctx.JSON(200, msgId)
}
