package leafBot

import (
	"encoding/json"
	"github.com/hjson/hjson-go"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	yaml "gopkg.in/yaml.v3"
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
	Bots     []*Bot `json:"bots"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
}

var (
	config = new(Config)
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	},
	)
}

const (
	JSON  = "json"
	HJSON = "hjson"
	YAML  = "yaml"
)

func LoadConfig(path string, fileType string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		file, err = os.OpenFile("config.json", os.O_RDWR, 0777)
		if err != nil {
			log.Panicln(err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Debugln(err.Error())
		}
	}(file)
	data, _ := io.ReadAll(file)
	switch fileType {
	case JSON:
		{
			err = json.Unmarshal(data, config)
		}
	case HJSON:
		{
			err = hjson.Unmarshal(data, config)
		}
	case YAML:
		{
			err = yaml.Unmarshal(data, config)
		}
	}
	if err != nil {
		log.Panicln("加载配置文件失败")
	}

	log.SetLevel(GetLogLevel(config.LogLevel))
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

	}
	if err := http.ListenAndServe(config.Host+":"+strconv.Itoa(config.Port), nil); err != nil {
		log.Panicln("监听端口失败，端口可能被占用")
	}
}

func GetLogLevel(level string) log.Level {
	switch level {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	default:
		return log.InfoLevel
	}
}
