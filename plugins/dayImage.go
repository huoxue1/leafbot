package plugins

import (
	"encoding/json"
	"github.com/3343780376/leafBot" //nolint:gci
	"github.com/3343780376/leafBot/message"
	"io"
	"net/http"
	"strconv"
)

type dayPicture struct {
	Status int `json:"status"`
	Bing   struct {
		URL       string `json:"url"`
		Copyright string `json:"copyright"`
	} `json:"bing"`
}

func UseDayImage() {
	leafBot.OnCommand("/dayPic").
		SetPluginName("每日一图").
		SetWeight(10).
		SetBlock(false).
		AddAllies("一图").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				if len(args) == 0 {
					image, err := getDayImage(0)
					if err != nil {
						return
					}
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, []message.MessageSegment{message.Text(image.Bing.Copyright), message.Image(image.Bing.URL)})
				} else {
					day, _ := strconv.Atoi(args[0])
					image, err := getDayImage(day)
					if err != nil {
						return
					}
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, []message.MessageSegment{message.Text(image.Bing.Copyright), message.Image(image.Bing.URL)})
				}
			})
}

func getDayImage(day int) (dayPicture, error) {
	resp, err := http.Get("https://api.no0a.cn/api/bing/" + strconv.Itoa(day))
	if err != nil {
		return dayPicture{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return dayPicture{}, err
	}
	picture := dayPicture{}
	err = json.Unmarshal(data, &picture)
	return picture, err
}
