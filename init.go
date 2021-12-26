package leafbot

import (
	_ "embed"
	"io"
	"os"
	"path"

	//nolint:gci
	"time"

	//nolint:gci
	rotates "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/huoxue1/leafbot/utils"
)

//go:embed config/label.txt
var label string

// init
/*
   @Description:
*/
func init() {
	err := initConfig()
	if err != nil {
		log.Errorln(err.Error())
		log.Infoln("配置文件加载失败或者不存在")
		log.Infoln("将生成默认配置文件")
		LoadConfig()
	}
	w, err := rotates.New(path.Join("logs", "%Y-%m-%d.log"), rotates.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("rotates init err: %v", err)
		panic(err)
	}
	f := &utils.LogFormat{
		TimeStampFormat: "2006-01-02 15:04:05",
		LogContent:      "[%time%] [%file%] [%lvl%] : %msg% \n",
		LogTruncate:     defaultConfig.LogTruncate,
	}

	levels := utils.GetLogLevel(defaultConfig.LogLevel)
	hook = utils.NewLogHook(f, levels, w)
	log.AddHook(hook)
	level, err := log.ParseLevel(defaultConfig.LogLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
	log.Infoln("\n" + label)

	go InitWindow()
}

const (
	JSON  = "json"
	HJSON = "hjson"
	YAML  = "yaml"
)

// InitConfig
/*
   @Description:
   @param path string
   @param fileType string
*/
func initConfig() error {
	file, err := os.OpenFile("./config/config.yml", os.O_RDWR, 0777)
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

	err = yaml.Unmarshal(data, defaultConfig)

	if err != nil {
		log.Errorln(err)
		return err
	}
	utils.SetConfig(defaultConfig.EnablePlaywright)
	//hook.AddLevel(utils.GetLogLevel(defaultConfig.LogLevel)...)
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
	if defaultConfig.EnablePlaywright {
		go utils.PwInit()
	}
}
