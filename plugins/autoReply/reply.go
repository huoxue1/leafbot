package autoReply

import (
	_ "embed"
	"fmt"
	"github.com/3343780376/leafBot"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//go:embed data.json
var data []byte

func Load(filePath string) error {
	defer func() {
		err := recover()
		if err != nil {

		}
	}()
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		log.Infoln("自动回复json文件解析失败，将使用默认配置")
	} else {
		data, err = io.ReadAll(file)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.Infoln("已成功加载词库:" + filePath)
	}

	content := gjson.ParseBytes(data)

	leafBot.OnMessage("").
		SetPluginName("自动回复").
		AddRule(func(event leafBot.Event, bot *leafBot.Bot) bool {

			if leafBot.DefaultConfig.Plugins.EnableReplyTome {
				if event.MessageType == "private" {
					return true
				}
				msg := event.GetMsg()
				for _, segment := range msg {
					if segment.Type == "at" && segment.Data["qq"] == strconv.Itoa(event.SelfId) {
						return true
					}
				}

				return false
			} else {
				return true
			}
		}).SetWeight(10).AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		all := strings.ReplaceAll(event.Message.CQString(), fmt.Sprintf("[CQ:at,qq=%d]", event.SelfId), "")
		result := content.Get(strings.TrimSpace(all))
		if result.String() == "" {
			return
		} else {
			switch result.Type {
			case 10:
				{
					bot.Send(event, result.String())
				}
			default:
				{
					i := len(result.Array())
					if i <= 0 {
						return
					}
					r := result.Array()[rand.Intn(i)]
					bot.Send(event, r.String())
				}
			}
		}
	})
	return nil
}
