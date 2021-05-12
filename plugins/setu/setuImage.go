package main

import (
	"github.com/everpcpc/pixiv"
)

var (
	app *pixiv.AppPixivAPI
)

func init() {
	//_, err := pixiv.Login("3343780376@qq.com", "1743224847gou")
	//if err != nil {
	//	log.Errorln("pixiv登录失败"+err.Error())
	//}
	app = pixiv.NewApp()
}

func getImage() {
	result, _ := app.SearchIllust("r18", "partial_match_for_tags", "date_desc", "", "", 10)
	for _, illust := range result.Illusts {
		app.Download(illust.ID, "./a.png")
	}
}

func main() {
	getImage()
}
