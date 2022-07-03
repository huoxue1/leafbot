package leafbot

import (
	_ "embed"
	"path"

	//nolint:gci
	"time"

	//nolint:gci
	rotates "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/leafbot/utils"
)

// InitBots
/*
   @Description:
*/
func InitBots(config Config) {
	LoadConfig(&config)
	w, err := rotates.New(path.Join("logs", "%Y-%m-%d.log"), rotates.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("rotates init err: %v", err)
		panic(err)
	}
	f := &utils.LogFormat{
		TimeStampFormat: "2006-01-02 15:04:05",
		LogContent:      "[%time%] [%file%] [%lvl%] : %msg% \n",
		LogTruncate:     true,
	}

	levels := utils.GetLogLevel(GetConfig().LogLevel)
	hook := utils.NewLogHook(f, levels, w)
	log.AddHook(hook)
	level, err := log.ParseLevel(GetConfig().LogLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
	// log.Infoln("\n" + label)
	go eventMain()
}
