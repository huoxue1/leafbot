package weibo

import (
	"encoding/json"
	"github.com/3343780376/leafBot"
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func init() {
	leafBot.OnCommand("/weibo").
		AddAllies("热搜").
		SetWeight(10).
		SetBlock(false).
		SetPluginName("微博热搜").
		AddHandle(weiBoHandle)
}

func weiBoHandle(event leafBot.Event, bot *leafBot.Bot, args []string) {
	draw(10)
}

func draw(limit int) {
	context := gg.NewContext(300, 20*limit)
	weibo, err := getData()
	if err != nil {
		return
	}
	for i := 0; i < limit; i++ {
		context.DrawString(weibo.Data[i].Name, 0, float64(20*i))
	}
	err = context.SavePNG("weibo.png")
	if err != nil {
		log.Debugln("图片保存失败")
	}
}

func getData() (Weibo, error) {
	resp, err := http.Get("https://api.hmister.cn/weibo/")
	if err != nil {
		return Weibo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weibo{}, err
	}
	weibo := Weibo{}
	err = json.Unmarshal(data, &weibo)
	return weibo, err
}
