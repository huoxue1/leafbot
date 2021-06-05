package plugins

import (
	"encoding/json"
	"errors"
	"github.com/3343780376/leafBot" //nolint:gci
	"github.com/3343780376/leafBot/message"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings" //nolint:gci
)

type MusicQQ struct {
	Code int `json:"code"`
	Data struct {
		Keyword  string        `json:"keyword"`
		Priority int           `json:"priority"`
		Qc       []interface{} `json:"qc"`
		Semantic struct {
			Curnum   int           `json:"curnum"`
			Curpage  int           `json:"curpage"`
			List     []interface{} `json:"list"`
			Totalnum int           `json:"totalnum"`
		} `json:"semantic"`
		Song struct {
			Curnum  int `json:"curnum"`
			Curpage int `json:"curpage"`
			List    []struct {
				Albumid          int    `json:"albumid"`
				Albummid         string `json:"albummid"`
				Albumname        string `json:"albumname"`
				AlbumnameHilight string `json:"albumname_hilight"`
				Alertid          int    `json:"alertid"`
				BelongCD         int    `json:"belongCD"`
				CdIdx            int    `json:"cdIdx"`
				Chinesesinger    int    `json:"chinesesinger"`
				Docid            string `json:"docid"`
				Grp              []struct {
					Albumid          int    `json:"albumid"`
					Albummid         string `json:"albummid"`
					Albumname        string `json:"albumname"`
					AlbumnameHilight string `json:"albumname_hilight"`
					Alertid          int    `json:"alertid"`
					BelongCD         int    `json:"belongCD"`
					CdIdx            int    `json:"cdIdx"`
					Chinesesinger    int    `json:"chinesesinger"`
					Docid            string `json:"docid"`
					Interval         int    `json:"interval"`
					Isonly           int    `json:"isonly"`
					Lyric            string `json:"lyric"`
					LyricHilight     string `json:"lyric_hilight"`
					MediaMid         string `json:"media_mid"`
					Msgid            int    `json:"msgid"`
					NewStatus        int    `json:"newStatus"`
					Nt               int64  `json:"nt"`
					Pay              struct {
						Payalbum      int `json:"payalbum"`
						Payalbumprice int `json:"payalbumprice"`
						Paydownload   int `json:"paydownload"`
						Payinfo       int `json:"payinfo"`
						Payplay       int `json:"payplay"`
						Paytrackmouth int `json:"paytrackmouth"`
						Paytrackprice int `json:"paytrackprice"`
					} `json:"pay"`
					Preview struct {
						Trybegin int `json:"trybegin"`
						Tryend   int `json:"tryend"`
						Trysize  int `json:"trysize"`
					} `json:"preview"`
					Pubtime int `json:"pubtime"`
					Pure    int `json:"pure"`
					Singer  []struct {
						ID          int    `json:"id"`
						Mid         string `json:"mid"`
						Name        string `json:"name"`
						NameHilight string `json:"name_hilight"`
					} `json:"singer"`
					Size128         int    `json:"size128"`
					Size320         int    `json:"size320"`
					Sizeape         int    `json:"sizeape"`
					Sizeflac        int    `json:"sizeflac"`
					Sizeogg         int    `json:"sizeogg"`
					Songid          int    `json:"songid"`
					Songmid         string `json:"songmid"`
					Songname        string `json:"songname"`
					SongnameHilight string `json:"songname_hilight"`
					StrMediaMid     string `json:"strMediaMid"`
					Stream          int    `json:"stream"`
					Switch          int    `json:"switch"`
					T               int    `json:"t"`
					Tag             int    `json:"tag"`
					Type            int    `json:"type"`
					Ver             int    `json:"ver"`
					Vid             string `json:"vid"`
					Format          string `json:"format,omitempty"`
					Songurl         string `json:"songurl,omitempty"`
				} `json:"grp"`
				Interval     int    `json:"interval"`
				Isonly       int    `json:"isonly"`
				Lyric        string `json:"lyric"`
				LyricHilight string `json:"lyric_hilight"`
				MediaMid     string `json:"media_mid"`
				Msgid        int    `json:"msgid"`
				NewStatus    int    `json:"newStatus"`
				Nt           int64  `json:"nt"`
				Pay          struct {
					Payalbum      int `json:"payalbum"`
					Payalbumprice int `json:"payalbumprice"`
					Paydownload   int `json:"paydownload"`
					Payinfo       int `json:"payinfo"`
					Payplay       int `json:"payplay"`
					Paytrackmouth int `json:"paytrackmouth"`
					Paytrackprice int `json:"paytrackprice"`
				} `json:"pay"`
				Preview struct {
					Trybegin int `json:"trybegin"`
					Tryend   int `json:"tryend"`
					Trysize  int `json:"trysize"`
				} `json:"preview"`
				Pubtime int `json:"pubtime"`
				Pure    int `json:"pure"`
				Singer  []struct {
					ID          int    `json:"id"`
					Mid         string `json:"mid"`
					Name        string `json:"name"`
					NameHilight string `json:"name_hilight"`
				} `json:"singer"`
				Size128         int    `json:"size128"`
				Size320         int    `json:"size320"`
				Sizeape         int    `json:"sizeape"`
				Sizeflac        int    `json:"sizeflac"`
				Sizeogg         int    `json:"sizeogg"`
				Songid          int    `json:"songid"`
				Songmid         string `json:"songmid"`
				Songname        string `json:"songname"`
				SongnameHilight string `json:"songname_hilight"`
				StrMediaMid     string `json:"strMediaMid"`
				Stream          int    `json:"stream"`
				Switch          int    `json:"switch"`
				T               int    `json:"t"`
				Tag             int    `json:"tag"`
				Type            int    `json:"type"`
				Ver             int    `json:"ver"`
				Vid             string `json:"vid"`
				Format          string `json:"format,omitempty"`
				Songurl         string `json:"songurl,omitempty"`
			} `json:"list"`
			Totalnum int `json:"totalnum"`
		} `json:"song"`
		Tab       int           `json:"tab"`
		Taglist   []interface{} `json:"taglist"`
		Totaltime int           `json:"totaltime"`
		Zhida     struct {
			Chinesesinger int `json:"chinesesinger"`
			Type          int `json:"type"`
		} `json:"zhida"`
	} `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int    `json:"subcode"`
	Time    int    `json:"time"`
	Tips    string `json:"tips"`
}

type Music163 struct {
	Result struct {
		Songs []struct {
			ID      int    `json:"id"`
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

func UseMusicHandle() {
	defer func() {
		i := recover()
		if i != nil {
		}
	}()
	leafBot.OnCommand("/music").
		SetPluginName("查询歌曲").
		SetWeight(10).
		AddAllies("查询歌曲").
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				switch len(args) {
				case 0:
					{
						bot.SendMsg(event.MetaEventType, event.UserId, event.GroupId, message.Text("请正确输入参数"))
					}
				case 1:
					{
						if args[0] == "help" {
							bot.SendMsg(event.MetaEventType, event.UserId, event.GroupId,
								message.Text("歌曲查询命令未：\n/music或者查询歌曲\n\n"+
									"第一个参数为搜索信息，第二个参数为返回条数"))
						} else {
							music, err := searchMusic(args[0], 10, 0)
							if err != nil {
								bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Text("点歌出现错误，\n"+err.Error()))
								return
							}
							m := "搜索结果为：\n\n "
							for _, song := range music.Result.Songs {
								m += "id：" + strconv.Itoa(song.ID) + "\n"
								m += "歌名：" + song.Name + "\n"
								m += "作者：" + song.Artists[0].Name + "\n\n"
							}
							bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Text(m))
						}
					}
				case 2:
					{
						limit, _ := strconv.Atoi(args[1])
						music, err := searchMusic(args[0], limit, 0)
						if err != nil {
							bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Text("点歌出现错误，\n"+err.Error()))
							return
						}
						m := "搜索结果为：\n\n "
						for _, song := range music.Result.Songs {
							m += "id：" + strconv.Itoa(song.ID) + "\n"
							m += "歌名：" + song.Name + "\n"
							m += "作者：" + song.Artists[0].Name + "\n\n"
						}
						bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Text(m))
					}
				case 3:
					{
						limit, _ := strconv.Atoi(args[1])
						offset, _ := strconv.Atoi(args[2])
						music, err := searchMusic(args[0], limit, offset)
						if err != nil {
							bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Text("点歌出现错误，\n"+err.Error()))
							return
						}
						m := "搜索结果为：\n\n "
						for _, song := range music.Result.Songs {
							m += "id：" + strconv.Itoa(song.ID) + "\n"
							m += "歌名：" + song.Name + "\n"
							m += "作者：" + song.Artists[0].Name + "\n\n"
						}
						bot.SendMsg(event.MessageType, event.UserId, event.GroupId, m)
					}
				}
			})

	leafBot.OnCommand("/orderMusic").
		SetWeight(10).
		SetPluginName("点歌").
		AddAllies("点歌").
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				if len(args) < 2 {
					bot.Send(event, message.Text("请输入正确的参数"))
					return
				}
				if args[0] == "163" {
					if args[1] == "id" {
						id, err := strconv.Atoi(args[2])
						if err != nil {
							return
						}
						bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.Music("163", int64(id)))
					} else {
						name := args[1]
						music, err := searchMusic(name, 1, 0)
						log.Debugln(music)
						if err != nil {
							bot.Send(event, message.Text("搜索歌曲错误\n"+err.Error()))
							return
						}
						bot.Send(event, message.Music("163", int64(music.Result.Songs[0].ID)))

					}
				} else if args[0] == "qq" {
					log.Infoln("已触发qq音乐点歌")
					if args[1] == "id" {
						id, err := strconv.Atoi(args[2])
						if err != nil {
							return
						}
						bot.Send(event, message.Music("qq", int64(id)))
					} else {
						name := args[1]
						music, err := searchMusicByQQ(name, 1, 0)
						if err != nil {
							bot.Send(event, message.Text("搜索歌曲错误\n"+err.Error()))
							return
						}
						bot.Send(event, message.Music("qq", int64(music.Data.Song.List[0].Songid)))
					}
				}

			})
}

func searchMusic(name string, limit int, offset int) (Music163, error) {
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
		return Music163{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	data, _ := io.ReadAll(resp.Body)
	log.Debugln(string(data))
	music := Music163{}
	err = json.Unmarshal(data, &music)
	if err != nil {
		return Music163{}, err
	}
	if len(music.Result.Songs) < 1 {
		return music, errors.New("点歌错误")
	}
	return music, err
}

func searchMusicByQQ(name string, limit int, offset int) (MusicQQ, error) {
	values := url.Values{}
	values.Add("aggar", strconv.Itoa(1))
	values.Add("cr", strconv.Itoa(1))
	values.Add("flag_qc", strconv.Itoa(0))
	values.Add("p", strconv.Itoa(offset))
	values.Add("n", strconv.Itoa(limit))
	values.Add("w", name)

	resp, err := http.Get("https://c.y.qq.com/soso/fcgi-bin/client_search_cp?" + values.Encode())
	if err != nil {
		return MusicQQ{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	data, err := io.ReadAll(resp.Body)
	content := strings.TrimPrefix(string(data), "callback(")
	content = strings.TrimSuffix(content, ")")
	qq := MusicQQ{}
	err = json.Unmarshal([]byte(content), &qq)

	return qq, err
}
