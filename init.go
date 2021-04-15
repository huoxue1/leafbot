package leafBot

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Bot struct {
	Name string `json:"name"`

	SelfId int         `json:"self_id"`
	Client *connection `json:"con"`
}

type Config struct {
	Bots []*Bot `json:"bots"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var (
	config = new(Config)
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	},
	)
}

func LoadConfig(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		log.Panicln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	data, _ := io.ReadAll(file)
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Panicln("加载配置文件失败" + err.Error())
	}
	log.Infoln("已加载配置：" + string(data))
}

func InitBots() {
	go eventMain()
	http.HandleFunc("/cqhttp/ws", EventHandle)
	for _, bot := range config.Bots {
		run(bot)
		UseEchoHandle(bot.Name)
	}
	if err := http.ListenAndServe(config.Host+":"+strconv.Itoa(config.Port), nil); err != nil {
		log.Panicln("监听端口失败，端口可能被占用")
	}
}
