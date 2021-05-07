package leafBot

import (
	"encoding/json"
	"fmt"
	"github.com/hjson/hjson-go"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	yaml "gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Bot struct {
	Name string `json:"name"`

	SelfId int         `json:"self_id"`
	Client *connection `json:"con"`
}

type Config struct {
	Bots     []*Bot `json:"bots"`
	Admin    int    `json:"admin"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
}

var (
	DefaultConfig = new(Config)
)

// init
/*
   @Description:
*/
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
	err := initConfig(path, fileType)
	if err != nil {
		log.Infoln("配置文件加载失败或者不存在")
		log.Infoln("将启用默认配置文件")
	} else {
		return
	}
	input := ""
	log.Infoln("请输入机器人账号,多个账号用逗号进行分割：")
	_, err = fmt.Scanln(&input)
	if err != nil {
		log.Panicln(err)
	}
	selfIds := strings.Split(input, ",")
	for i, id := range selfIds {
		b := new(Bot)
		b.Name = fmt.Sprintf("bot%d", i)
		b.SelfId, err = strconv.Atoi(id)
		if err != nil {
			log.Panicln("输入的账号有错\n" + err.Error())
		}
		DefaultConfig.Bots = append(DefaultConfig.Bots, b)
	}
	writeGoConfig(DefaultConfig.Bots[0].SelfId)
	DefaultConfig.Admin = 0
	DefaultConfig.Host = "127.0.0.1"
	DefaultConfig.Port = 8080
	DefaultConfig.LogLevel = "info"
	config, err := json.MarshalIndent(&DefaultConfig, "", "  ")
	if err != nil {
		log.Infoln("json反向序列号失败")
		return
	}
	file, err := os.OpenFile("config.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("打开config.json文件失败\n" + err.Error())
		return
	}
	_, err = file.WriteString(string(config))
	if err != nil {
		log.Infoln("写入配置到文件失败")
	}
	log.Infoln("成功写入默认配置到config.json")
}

// InitConfig
/*
   @Description:
   @param path string
   @param fileType string
*/
func initConfig(path string, fileType string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		file, err = os.OpenFile("config.json", os.O_RDWR, 0777)
		if err != nil {
			return err
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
			err = json.Unmarshal(data, DefaultConfig)
		}
	case HJSON:
		{
			err = hjson.Unmarshal(data, DefaultConfig)
		}
	case YAML:
		{
			err = yaml.Unmarshal(data, DefaultConfig)
		}
	}
	if err != nil {
		log.Errorln(err)
		return err
	}

	log.SetLevel(GetLogLevel(DefaultConfig.LogLevel))
	log.Infoln("已加载配置：" + string(data))
	return err
}

// InitBots
/*
   @Description:
*/
func InitBots() {
	go eventMain()

	http.HandleFunc("/cqhttp/ws", EventHandle)
	for _, bot := range DefaultConfig.Bots {
		run(bot)

	}
	if err := http.ListenAndServe(DefaultConfig.Host+":"+strconv.Itoa(DefaultConfig.Port), nil); err != nil {
		log.Panicln("监听端口失败，端口可能被占用")
	}
}

// GetLogLevel
/*
   @Description:
   @param level string
   @return log.Level
*/
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
