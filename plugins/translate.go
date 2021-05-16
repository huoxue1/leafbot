package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Tran struct {
	Type            string `json:"type"`
	ErrorCode       int    `json:"errorCode"`
	ElapsedTime     int    `json:"elapsedTime"`
	TranslateResult [][]struct {
		Src string `json:"src"`
		Tgt string `json:"tgt"`
	} `json:"translateResult"`
}

func UseTranslateHandle() {

	leafBot.OnCommand("/ts").
		AddAllies("翻译").
		SetWeight(10).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				//if len(args)<1 {
				//	bot.Send(event,"请输入正确的参数")
				//	return
				//}
				switch len(args) {
				case 0:
					{
						nextEvent, err := bot.GetOneEvent(func(event1 leafBot.Event, bot2 *leafBot.Bot) bool {
							if event1.UserId == event.UserId && event1.GroupId == event.GroupId {
								return true
							}
							return false
						})
						if err != nil {
							return
						}
						tran, err := translate(nextEvent.Message, "AUTO")
						if err != nil {
							bot.Send(event, message.Text("翻译失败："+err.Error()))
							return
						}
						message1 := ""
						for _, result := range tran.TranslateResult {
							for _, s := range result {
								message1 += s.Tgt + "\n"
							}
						}
						bot.Send(event, message.Text("翻译结果为：\n"+message1))
						return
					}
				}

			})

}

func translate(text string, types string) (Tran, error) {
	resp, err := http.Get(fmt.Sprintf("http://fanyi.youdao.com/translate?&doctype=json&type=%s&i=%s", types, text))
	if err != nil {
		return Tran{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Debugln(err)
		}
	}(resp.Body)
	data, _ := io.ReadAll(resp.Body)
	tran := Tran{}
	err = json.Unmarshal(data, &tran)
	return tran, err
}
