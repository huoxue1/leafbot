package message

import "fmt"

func At(userId int) string {
	return fmt.Sprintf("[CQ:at,qq=%d]", userId)
}

func AtAll() string {
	return "[CQ:at,qq=all]"
}
func Image(file string, data map[string]interface{}) string {
	date := "[CQ:image,file=" + file
	for k, datum := range data {
		date += fmt.Sprintf(","+k+"=%v", datum)
	}

	return date + "]"

}

func Face(id int) string {
	return fmt.Sprintf("[CQ:face,id=%d]", id)
}

func Share(url, title, content, image string) string {
	return fmt.Sprintf("[CQ:share,url=%v,title=%v,image=%v,content=%v]", url, title, image, content)
}

func Music(Type, id int) string {
	return fmt.Sprintf("[CQ:music,type=%v,id=%v]", Type, id)
}

func Poke(userid int) string {
	return fmt.Sprintf("[CQ:poke,qq=%v]", userid)
}

func Gift(userId, id int) string {
	return fmt.Sprintf("[CQ:gift,qq=%v,id=%v]", userId, id)
}

func Tts(text interface{}) string {
	return fmt.Sprintf("[CQ:tts,text=%v]", text)
}
