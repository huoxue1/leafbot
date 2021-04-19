package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/3343780376/leafBot"
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

	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		//if len(args)<1 {
		//	bot.Send(event,"请输入正确的参数")
		//	return
		//}
		switch len(args) {
		case 0:
			{
				nextEvent := bot.GetNextEvent(10, event.UserId)
				tran, err := translate(nextEvent.Message, "AUTO")
				if err != nil {
					bot.Send(event, "翻译失败："+err.Error())
					return
				}
				message := ""
				for _, result := range tran.TranslateResult {
					for _, s := range result {
						message += s.Tgt + "\n"
					}
				}
				bot.Send(event, "翻译结果为：\n"+message)
				return
			}
		}

	}, "/ts", []string{"翻译"}, nil, 1, false)
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
