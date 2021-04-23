package main

import (
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
)

//func main() {
//	ph("你好", "❓")
//}
func ph(text1, text2 string) {

	text1Length := spacing(text1, 90)
	text2Length := spacing(text2, 90)
	context := gg.NewContext(text1Length+text2Length+110, 90+20)
	context.SetRGB255(20, 20, 20)
	context.Clear()
	context.SetRGB255(255, 255, 255)
	if err := context.LoadFontFace("./plugins/ph/NotoSansBold.ttf", 90); err != nil {
		log.Debugln(err)
	}
	context.DrawString(text1, 0, 90)
	context.SetRGB255(247, 151, 29)
	context.DrawRectangle(float64(text1Length+55), 10, float64(text2Length+55), 90)
	context.Fill()
	context.SetRGB255(20, 20, 20)
	context.DrawString(text2, float64(text1Length)+55, 90)
	context.SavePNG("ph.png")
}

func spacing(text string, fontSize int) int {
	space := 0
	for _, i := range text {
		if i <= '\u9fff' && i >= '\u4e00' {
			space += fontSize
		} else {
			space += fontSize / 2
		}
	}
	return space
}
