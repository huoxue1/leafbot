package plugins

import (
	"encoding/json"
	"github.com/3343780376/leafBot"
	"io"
	"net/http"
	"strconv"
)

type dayPicture struct {
	Status int `json:"status"`
	Bing   struct {
		Url       string `json:"url"`
		Copyright string `json:"copyright"`
	} `json:"bing"`
}

func UseDayImage() {
	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		if len(args) == 0 {
			image, err := getDayImage(0)
			if err != nil {
				return
			}
			bot.SendMsg(event.MessageType, event.UserId, event.GroupId, image.Bing.Copyright+"\n[CQ:image,file="+image.Bing.Url+"]", false)
		} else {
			day, _ := strconv.Atoi(args[0])
			image, err := getDayImage(day)
			if err != nil {
				return
			}
			bot.SendMsg(event.MessageType, event.UserId, event.GroupId, image.Bing.Copyright+"\n[CQ:image,file="+image.Bing.Url+"]", false)
		}
	}, "/dayPic", []string{"一图"}, nil, 10, false)
}

func getDayImage(day int) (dayPicture, error) {
	resp, err := http.Get("https://api.no0a.cn/api/bing/" + strconv.Itoa(day))
	if err != nil {
		return dayPicture{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	data, err := io.ReadAll(resp.Body)
	picture := dayPicture{}
	err = json.Unmarshal(data, &picture)
	return picture, err
}
