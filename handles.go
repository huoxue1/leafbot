package leafBot

import (
	"encoding/json"
	"github.com/3343780376/leafBot/message"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type dayPicture struct {
	Status int `json:"status"`
	Bing   struct {
		Url       string `json:"url"`
		Copyright string `json:"copyright"`
	} `json:"bing"`
}

type Music struct {
	Result struct {
		Songs []struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				Id        int           `json:"id"`
				Name      string        `json:"name"`
				PicUrl    interface{}   `json:"picUrl"`
				Alias     []interface{} `json:"alias"`
				AlbumSize int           `json:"albumSize"`
				PicId     int           `json:"picId"`
				Img1V1Url string        `json:"img1v1Url"`
				Img1V1    int           `json:"img1v1"`
				Trans     interface{}   `json:"trans"`
			} `json:"artists"`
			Album struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Artist struct {
					Id        int           `json:"id"`
					Name      string        `json:"name"`
					PicUrl    interface{}   `json:"picUrl"`
					Alias     []interface{} `json:"alias"`
					AlbumSize int           `json:"albumSize"`
					PicId     int           `json:"picId"`
					Img1V1Url string        `json:"img1v1Url"`
					Img1V1    int           `json:"img1v1"`
					Trans     interface{}   `json:"trans"`
				} `json:"artist"`
				PublishTime int64    `json:"publishTime"`
				Size        int      `json:"size"`
				CopyrightId int      `json:"copyrightId"`
				Status      int      `json:"status"`
				PicId       int64    `json:"picId"`
				Mark        int      `json:"mark"`
				Alia        []string `json:"alia,omitempty"`
			} `json:"album"`
			Duration    int         `json:"duration"`
			CopyrightId int         `json:"copyrightId"`
			Status      int         `json:"status"`
			Alias       []string    `json:"alias"`
			Rtype       int         `json:"rtype"`
			Ftype       int         `json:"ftype"`
			Mvid        int         `json:"mvid"`
			Fee         int         `json:"fee"`
			RUrl        interface{} `json:"rUrl"`
			Mark        int         `json:"mark"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

func UseEchoHandle() {

	AddCommandHandle(func(event Event, bot *Bot, args []string) {
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, args[0], false)
	}, "/echo", nil, nil, 1, false)
}

func UseDayImage() {
	AddCommandHandle(func(event Event, bot *Bot, args []string) {
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

func UseMusicHandle() {
	AddCommandHandle(func(event Event, bot *Bot, args []string) {
		switch len(args) {
		case 0:
			{
				bot.SendMsg(event.MetaEventType, event.UserId, event.GroupId, "请正确输入参数", false)
			}
		case 1:
			{
				if args[0] == "help" {
					bot.SendMsg(event.MetaEventType, event.UserId, event.GroupId, "歌曲查询命令未：\n/music或者查询歌曲\n\n"+
						"第一个参数为搜索信息，第二个参数为返回条数", false)
				} else {
					music, err := searchMusic(args[0], 10, 0)
					if err != nil {
						bot.SendMsg(event.MessageType, event.UserId, event.GroupId, "点歌出现错误，\n"+err.Error(), false)
						return
					}
					m := "搜索结果为：\n\n "
					for _, song := range music.Result.Songs {
						m += "id：" + strconv.Itoa(song.Id) + "\n"
						m += "歌名：" + song.Name + "\n"
						m += "作者：" + song.Artists[0].Name + "\n\n"
					}
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, m, false)
				}
			}
		case 2:
			{
				limit, _ := strconv.Atoi(args[1])
				music, err := searchMusic(args[0], limit, 0)
				if err != nil {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, "点歌出现错误，\n"+err.Error(), false)
					return
				}
				m := "搜索结果为：\n\n "
				for _, song := range music.Result.Songs {
					m += "id：" + strconv.Itoa(song.Id) + "\n"
					m += "歌名：" + song.Name + "\n"
					m += "作者：" + song.Artists[0].Name + "\n\n"
				}
				bot.SendMsg(event.MessageType, event.UserId, event.GroupId, m, false)
			}
		case 3:
			{
				limit, _ := strconv.Atoi(args[1])
				offset, _ := strconv.Atoi(args[2])
				music, err := searchMusic(args[0], limit, offset)
				if err != nil {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, "点歌出现错误，\n"+err.Error(), false)
					return
				}
				m := "搜索结果为：\n\n "
				for _, song := range music.Result.Songs {
					m += "id：" + strconv.Itoa(song.Id) + "\n"
					m += "歌名：" + song.Name + "\n"
					m += "作者：" + song.Artists[0].Name + "\n\n"
				}
				bot.SendMsg(event.MessageType, event.UserId, event.GroupId, m, false)
			}
		}
	}, "/music", []string{"查询歌曲"}, nil, 10, false)

	AddCommandHandle(func(event Event, bot *Bot, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}
		bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Music(163, id), false)
	}, "/orderMusic", []string{"点歌"}, nil, 10, false)
}

func searchMusic(name string, limit int, offset int) (Music, error) {
	values := url.Values{}
	values.Add("csrf_token", "")
	values.Add("hlpretag", "")
	values.Add("hlposttag", "")
	values.Add("s", name)
	values.Add("type", strconv.Itoa(1))
	values.Add("offset", strconv.Itoa(offset))
	values.Add("total", "true")
	values.Add("limit", strconv.Itoa(limit))
	resp, err := http.PostForm("http://music.163.com/api/search/get/web", values)
	//resp, err := http.Get(fmt.Sprintf("http://music.163.com/api/search/get/web?csrf_token=&hlpretag=&hlposttag=&s=%v&type=1&offset=%d&total=true&limit=%d",
	//	name, offset, limit))
	if err != nil {
		return Music{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	data, _ := io.ReadAll(resp.Body)
	log.Debugln(string(data))
	music := Music{}
	err = json.Unmarshal(data, &music)
	if err != nil {
		return Music{}, err
	}
	return music, err
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
