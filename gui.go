package leafbot

import ( //nolint:gci

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http" //nolint:gci
	"os"
	"os/signal"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/huoxue1/lorca"
	"github.com/huoxue1/test3"
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/leafbot/message"
	"github.com/huoxue1/leafbot/utils"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	logConn  *websocket.Conn
	dataCoon *websocket.Conn
	ui       lorca.UI
	engine   *gin.Engine
)

// GetEngine
/**
 * @Description: 获取web的引擎
 * @return *gin.Engine
 * @return error
 * example
 */
func GetEngine() (*gin.Engine, error) {
	if engine == nil {
		return nil, errors.New("engine not init")
	}
	return engine, nil
}

//OpenUi
/**
 * @Description:
 */
func OpenUi() {
	var err error
	ui, err = lorca.New("http://127.0.0.1:3000/static/gui/static/html/default.html", "", 800, 600)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c) //nolint:govet
		for {
			log.Infoln(<-c)
			err := ui.Close()
			if err != nil {
				return
			}
			err = utils.PW.Stop()
			if err != nil {
				return
			}
			err = utils.Browser.Close()
			if err != nil {
				return
			}
		}
	}()
	defer func(ui lorca.UI) {
		err := ui.Close()
		if err != nil {
			fmt.Println("关闭ui失败")
		}
	}(ui)
	if err != nil {
		log.Panic(err)
	}
	<-ui.Done()
	os.Exit(3)
}

// InitWindow
/**
 * @Description:
 * example
 */
func InitWindow() {
	defer func() {
		err := recover()
		log.Errorln("初始化web控制台出现错误")
		log.Errorln(err)
	}()
	if defaultConfig.LogLevel != "debug" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	log.Infoln("web页面：http://127.0.0.1:3000")
	engine = gin.New()
	engine.Use(Cors())
	engine.StaticFS("/dist", http.FS(test3.Dist))
	engine.POST("/get_bots", GetConfig)
	engine.POST("/get_group_list", GetGroupList)
	engine.POST("/get_friend_list", GetFriendList)
	engine.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/dist/dist/default.html")
	})
	engine.POST("/update_plugin_states", func(context *gin.Context) {
		// id := context.PostForm("id")
		// status, err := strconv.ParseBool(context.PostForm("status"))
		// if err != nil {
		//	log.Errorln("改变插件状态出错" + err.Error())
		// }
		// if status {
		//	StartPluginByID(id)
		// } else {
		//	BanPluginByID(id)
		// }
		// todo
		context.JSON(200, nil)
	})

	engine.POST("/get_plugins", func(context *gin.Context) {
		//list := GetHandleList()
		//var pluginList []BaseHandle
		//
		//for _, handles := range list {
		//	pluginList = append(pluginList, handles...)
		//}
		//sort.SliceStable(pluginList, func(i, j int) bool {
		//	id1, _ := strconv.Atoi(pluginList[i].ID)
		//	id2, _ := strconv.Atoi(pluginList[j].ID)
		//	return id1 < id2
		//})
		context.JSON(200, nil)
	})

	engine.GET("/get_log", func(context *gin.Context) {
		conn, err := upGrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Errorln(err)
			return
		}
		if logConn == nil {
			logConn = conn
			hook.EnableLogChan = true
			go func() {
				for {
					event := <-hook.LogChan
					err = logConn.WriteMessage(websocket.TextMessage, []byte(event))
					if err != nil {
						fmt.Println("前端日志消息发送失败" + err.Error())
						hook.EnableLogChan = false
						break
					}
				}
			}()
		}
		logConn = conn
	})

	engine.POST("/get_all_config", getAllConfig)

	engine.POST("/send_msg", CallApi)
	engine.GET("/data", data)
	if err := engine.Run("127.0.0.1:3000"); err != nil {
		log.Debugln(err.Error())
	}
}

func data(ctx *gin.Context) {
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Infoln(err)
		return
	}

	if dataCoon == nil {
		dataCoon = conn
		ENABLE = true

		go func() {
			for {
				event := <-MessageChan
				log.Debugln("已向前端发送信息")
				err = dataCoon.WriteJSON(&event)
				if err != nil {
					fmt.Println("消息发送失败" + err.Error())
					ENABLE = false
					break
				}
			}
		}()
	}
	dataCoon = conn
}

func getAllConfig(ctx *gin.Context) {
	ctx.JSON(200, defaultConfig)
}

//GetConfig
/**
 * @Description:
 * @param ctx
 */
func GetConfig(ctx *gin.Context) {
	bots := make([]int64, 1)
	for i := range driver.GetBots() {
		bots = append(bots, i)
	}
	ctx.JSON(200, bots)
}

//GetGroupList
/**
 * @Description:
 * @param ctx
 */
func GetGroupList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		var data map[string]interface{}
		err := ctx.BindJSON(&data)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		selfID = int(data["self_id"].(float64))
	}

	bot := GetBotById(selfID)
	var resp []interface{}
	list := bot.(OneBotApi).GetGroupList().String()
	err = json.Unmarshal([]byte(list), &resp)
	if err != nil {
		log.Errorln(err.Error())
	}
	ctx.JSON(200, resp)
}

func GetFriendList(ctx *gin.Context) {
	selfID, err := strconv.Atoi(ctx.PostForm("self_id"))
	if err != nil {
		log.Errorln(err.Error())
		var data map[string]interface{}
		err := ctx.BindJSON(&data)
		if err != nil {
			log.Errorln(err.Error())
			log.Errorln("绑定错误")
			return
		}
		selfID = int(data["self_id"].(float64))
	}
	bot := GetBotById(selfID)
	var resp []interface{}
	list := bot.(OneBotApi).GetFriendList().String()
	err = json.Unmarshal([]byte(list), &resp)
	if err != nil {
		log.Errorln(err.Error())
		log.Errorln("解析json错误")
	}
	ctx.JSON(200, resp)
}

func CallApi(ctx *gin.Context) {
	selfID, err := strconv.ParseInt(ctx.PostForm("self_id"), 10, 64)
	if err != nil {
		return
	}
	id, err := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	message1 := ctx.PostForm("message")
	messageType := ctx.PostForm("message_type")
	if err != nil {
		var data map[string]interface{}
		err := ctx.BindJSON(&data)
		if err != nil {
			ctx.JSON(404, nil)
			return
		}
		selfID = data["self_id"].(int64)
		id = data["id"].(int64)
		message1 = data["message"].(string)
		messageType = data["message_type"].(string)
	}
	bot := GetBotById(int(selfID))
	msgID := bot.(OneBotApi).SendMsg(messageType, int(id), int(id), message.ParseMessageFromString(message1))
	ctx.JSON(200, msgID)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session, Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
