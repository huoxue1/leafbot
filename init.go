package leafBot

import (
	_ "embed"
	"encoding/json"
	"github.com/hjson/hjson-go" //nolint:gci
	"github.com/huoxue1/leafBot/utils"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path"
	//nolint:gci
	"time"
)

//go:embed config/label.txt
var label string

// init
/*
   @Description:
*/
func init() {

	err := initConfig(YAML)
	if err != nil {
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
	}

	levels := utils.GetLogLevel(DefaultConfig.LogLevel)
	hook = utils.NewLogHook(f, levels, w)
	log.AddHook(hook)
	level, err := log.ParseLevel(DefaultConfig.LogLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
	log.Infoln("\n" + label)
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
func initConfig(fileType string) error {
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
	utils.SetConfig(DefaultConfig.EnablePlaywright)
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
	go InitWindow()
	go eventMain()
	if DefaultConfig.EnablePlaywright {
		go utils.PwInit()
	}
}
