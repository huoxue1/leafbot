package blacklist

import (
	"encoding/json"
	"fmt"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
)

type blackList struct {
	Users  []int `json:"users"`
	Groups []int `json:"groups"`
}

var (
	BlackList = blackList{}
)

func InitBlackList(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		log.Errorln("加载黑名单文件失败" + err.Error())
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Errorln("读取黑名单数据失败" + err.Error())
		return
	}
	err = json.Unmarshal(data, &BlackList)
	if err != nil {
		log.Errorln("格式化json失败" + err.Error())
		return
	}

	leafBot.OnPretreatment().SetPluginName("黑名单预处理").SetWeight(10).AddHandle(func(event leafBot.Event, bot *leafBot.Bot) bool {
		for _, user := range BlackList.Users {
			if user == event.UserId {
				return false
			}
		}
		for _, group := range BlackList.Groups {
			if event.GroupId == group {
				return false
			}
		}
		return true
	})

	leafBot.OnCommand("/add_blackList_user").
		AddRule(leafBot.OnlySuperUser). // 设置仅可管理员用户触发
		SetWeight(10).
		SetPluginName("添加黑名单用户").
		SetBlock(false).
		AddAllies("添加黑名单用户").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				datas := strings.Split(args[0], ",")
				for _, s := range datas {
					data, _ := strconv.Atoi(s)
					BlackList.Users = append(BlackList.Users, data)
				}
				content, _ := json.Marshal(&BlackList)
				_, err := file.Write(content)
				if err != nil {
					bot.Send(event, message.Text(err.Error()))
					return
				}
				bot.Send(event, message.Text("添加黑名单成功"))

			})

	leafBot.OnCommand("/add_blackList_group").
		AddRule(leafBot.OnlySuperUser).SetWeight(10).
		SetBlock(false).
		SetPluginName("添加黑名单群").
		AddAllies("添加黑名单群").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				datas := strings.Split(args[0], ",")
				for _, s := range datas {
					data, _ := strconv.Atoi(s)
					BlackList.Groups = append(BlackList.Groups, data)
				}
				content, _ := json.Marshal(&BlackList)
				_, err := file.Write(content)
				if err != nil {
					bot.Send(event, message.Text(err.Error()))
					return
				}
				bot.Send(event, message.Text("添加黑名单成功"))

			})

	leafBot.OnCommand("/get_blackList").SetPluginName("获取黑名单列表").AddRule(leafBot.OnlySuperUser).AddAllies("获取黑名单").SetBlock(false).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		msg := "黑名单：\n用户\n"
		for _, user := range BlackList.Users {
			msg += fmt.Sprintf("\t%d\n", user)
		}
		msg += "\n群\n"
		for _, group := range BlackList.Groups {
			msg += fmt.Sprintf("\t%d\n", group)
		}

		bot.Send(event, message.Text(msg))

	})

}
