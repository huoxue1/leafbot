package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

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
