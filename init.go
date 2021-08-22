package leafBot

import (
	"encoding/json"
	"fmt"
	"github.com/hjson/hjson-go" //nolint:gci
	"github.com/huoxue1/leafBot/utils"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	//nolint:gci
	"time"
)

type Bot struct {
	Name string `json:"name"`

	SelfId int         `json:"self_id"`
	Client *connection `json:"con"`
}

type Config struct {
	Bots         []*Bot   `json:"bots"`
	Admin        int      `json:"admin"`
	Host         string   `json:"host"`
	Port         int      `json:"port"`
	LogLevel     string   `json:"log_level"`
	SuperUser    []int    `json:"super_user"`
	CommandStart []string `json:"command_start"`
	Plugins      struct {
		FlashGroupID    int    `json:"flash_group_id"`
		AlApiToken      string `json:"al_api_token"`
		EnableReplyTome bool   `json:"enable_reply_tome"`
		Welcome         []struct {
			GroupId int    `json:"group_id"`
			Message string `json:"message"`
		} `json:"welcome"`
		GithubToken string `json:"github_token"`
	} `json:"plugins"`
}

var (
	DefaultConfig = new(Config)
	hook          *utils.LogHook
)

// init
/*
   @Description:
*/
func init() {
	err := initConfig(JSON)
	if err != nil {
		log.Infoln("配置文件加载失败或者不存在")
		log.Infoln("将启用默认配置文件")
		LoadConfig()
	}
	w, err := rotates.New(path.Join("logs", "%Y-%m-%d.log"), rotates.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("rotates init err: %v", err)
		panic(err)
	}
	f := &utils.LogFormat{
		TimeStampFormat: "2006-01-02 15:04:05",
		LogContent:      "[%time%] [LeafBot] [%lvl%]: %msg% \n",
	}
	levels := utils.GetLogLevel(DefaultConfig.LogLevel)
	hook = utils.NewLogHook(f, levels, w)
	log.AddHook(hook)
	level, err := log.ParseLevel(DefaultConfig.LogLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}

const (
	JSON  = "json"
	HJSON = "hjson"
	YAML  = "yaml"
)

func LoadConfig() {

	input := ""
	log.Infoln("请输入机器人账号")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Panicln(err)
	}

	selfID, err := strconv.Atoi(input)
	if err != nil {
		log.Errorln("输入有误")
	}
	writeGoConfig(selfID)
	DefaultConfig.Admin = 0
	DefaultConfig.Host = "127.0.0.1"
	DefaultConfig.Port = 8080
	DefaultConfig.LogLevel = "info"
	DefaultConfig.CommandStart = []string{"", "/"}
	config, err := json.MarshalIndent(&DefaultConfig, "", "  ")
	if err != nil {
		log.Infoln("json反向序列号失败")
		return
	}
	file, err := os.OpenFile("./config/config.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("打开config.json文件失败\n" + err.Error())
		return
	}
	_, err = file.WriteString(string(config))
	if err != nil {
		log.Infoln("写入配置到文件失败")
	}
	log.Infoln("成功写入默认配置到config.json")
	log.Infoln("程序将在五秒后重启")
	time.Sleep(5000)
	os.Exit(3)
	ui.Close()
}

// InitConfig
/*
   @Description:
   @param path string
   @param fileType string
*/
func initConfig(fileType string) error {
	file, err := os.OpenFile("./config/config.json", os.O_RDWR, 0777)
	if err != nil {
		return err
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

	//hook.AddLevel(utils.GetLogLevel(DefaultConfig.LogLevel)...)
	// log.Infoln("已加载配置：" + string(data))
	//log.SetLevel(log.DebugLevel)
	return err
}

// InitBots
/*
   @Description:
*/
func InitBots() {
	go eventMain()

	http.HandleFunc("/cqhttp/ws", eventHandle)
	for _, bot := range DefaultConfig.Bots {
		run(bot)

	}
	log.Infoln("listening in " + DefaultConfig.Host + "  " + strconv.Itoa(DefaultConfig.Port))
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
