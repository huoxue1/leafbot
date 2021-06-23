package searchImage

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func InitImage() {
	leafBot.OnCommand("搜图").SetPluginName("搜图").SetWeight(10).SetBlock(false).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		images, err := SearchImage(args[0])
		if err != nil {
			bot.Send(event, message.Text("接口报错了"+err.Error()))
			return
		}
		mess := message.Message{}
		for _, image := range images {
			mess = append(mess, message.Image(image))
			//mess = append(mess, message.CustomNode(event.Sender.NickName, int64(event.UserId), "[CQ:image,file="+image+"]"))

		}

		bot.Send(event, mess)
		//bot.SendGroupForwardMsg(event.GroupId, mess)
	})
}

func SearchImage(ketWord string) ([]string, error) {
	resp, err := http.Get("https://pixiv.kurocore.com/illust?keyword=" + ketWord)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var list []string
	doc.Find("div.illust-image a.cover img").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("data-original")
		list = append(list, "http:"+link)
	})
	return list, err
}
