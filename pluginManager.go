package leafBot

import (
	"encoding/base64" //nolint:gci
	"github.com/3343780376/leafBot/message"
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"io/ioutil" //nolint:gci
	"strconv"
)

// GetHandleList
/**
 * @Description: 获取所有的未被禁用handle
 * @return map[string][]string
 */

type BaseHandle struct {
	ID         string
	Name       string
	Enable     bool
	IsAdmin    bool
	HandleType string
}

func InitPluginManager() {
	OnCommand("/ban_plugin").SetPluginName("禁用插件").
		AddAllies("禁用插件").AddRule(OnlySuperUser).SetWeight(10).SetBlock(false).AddHandle(func(event Event, bot *Bot, args []string) {
		if len(args) < 0 {
			bot.Send(event, message.Text("参数不够"))
			return
		}
		BanPluginByID(args[0])
		bot.Send(event, message.Text("禁用插件成功"))
	})

	OnCommand("/use_plugin").SetPluginName("启用插件").
		AddAllies("启用插件").AddRule(OnlySuperUser).SetWeight(10).SetBlock(false).AddHandle(func(event Event, bot *Bot, args []string) {
		if len(args) < 0 {
			bot.Send(event, message.Text("参数不够"))
			return
		}
		StartPluginByID(args[0])
		bot.Send(event, message.Text("启用插件成功"))
	})

	OnCommand("/get_plugins").SetPluginName("获取插件列表").
		AddAllies("插件列表").AddRule(OnlySuperUser).SetWeight(10).SetBlock(false).AddHandle(func(event Event, bot *Bot, args []string) {
		handleList := GetHandleList()
		//for s, handles := range handleList {
		//	msg += s+"\n"
		//	for _, handle := range handles {
		//		msg += handle.Id+"\t\t\t"+handle.Name+"\t\t\t"+strconv.FormatBool(handle.Enable)+"\n"
		//	}
		//}
		draw(handleList)
		srcByte, err := ioutil.ReadFile("./config/plugin.png")
		if err != nil {
			log.Fatal(err)
		}

		res := base64.StdEncoding.EncodeToString(srcByte)

		bot.Send(event, message.Image("base64://"+res))
	})
}

func draw(data map[string][]BaseHandle) {
	context := gg.NewContext(900, 100*(8+pluginNum))
	context.SetRGB255(255, 255, 0)
	context.DrawRectangle(0, 0, 900, float64(100*(pluginNum+8)))
	//weibo, err := getData()
	context.Fill()
	if err := context.LoadFontFace("./config/NotoSansBold.ttf", 40); err != nil {
		log.Debugln(err)
	}
	context.SetRGB255(0, 0, 0)

	//for i := 0; i < limit; i++ {
	//	fmt.Println(weibo.Data[i].Name)
	//	context.DrawString(strconv.Itoa(i+1)+"："+weibo.Data[i].Name, 0, float64(100*(i+1)))
	//}
	context.DrawString("插件Id", 100, 100)
	context.DrawString("插件名", 400, 100)
	context.DrawString("是否启用", 700, 100)
	n := 2
	for s, handles := range data {
		context.SetRGB255(255, 0, 0)
		context.DrawString(s, 0, float64(100*n))
		n++
		context.SetRGB255(0, 0, 0)
		for _, handle := range handles {
			context.DrawString(handle.ID, 100, float64(100*n))
			context.DrawString(handle.Name, 300, float64(100*n))
			context.DrawString(strconv.FormatBool(handle.Enable), 600, float64(100*n))
			n++
		}
	}
	err := context.SavePNG("./config/plugin.png")
	if err != nil {
		log.Debugln("图片保存失败")
	}
}

func GetHandleList() map[string][]BaseHandle {
	var (
		list     = make(map[string][]BaseHandle)
		preList  []BaseHandle
		cmdList  []BaseHandle
		msgList  []BaseHandle
		reqList  []BaseHandle
		notList  []BaseHandle
		metaList []BaseHandle
	)

	for _, handle := range PretreatmentHandles {

		preList = append(preList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["预处理响应器"] = preList

	for _, handle := range CommandHandles {

		cmdList = append(cmdList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["command响应器"] = cmdList

	for _, handle := range MessageHandles {

		msgList = append(msgList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["message响应器"] = msgList

	for _, handle := range RequestHandles {

		reqList = append(reqList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["request响应器"] = reqList

	for _, handle := range NoticeHandles {

		notList = append(notList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["notice响应器"] = notList

	for _, handle := range MetaHandles {

		metaList = append(metaList, BaseHandle{
			ID:         handle.ID,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["meta响应器"] = metaList
	return list
}

// BanPluginByID
/**
 * @Description: 根据id禁用插件
 * @param id
 */
func BanPluginByID(id string) {
	handle, ok := CommandHandles.get(id)
	if ok {
		handle.(*commandHandle).Enable = false
		log.Infoln(handle.(*commandHandle).Name + "插件已被禁用")
	}
	handle, ok = MessageHandles.get(id)
	if ok {
		handle.(*messageHandle).Enable = false
		log.Infoln(handle.(*messageHandle).Name + "插件已被禁用")
	}
	handle, ok = RequestHandles.get(id)
	if ok {
		handle.(*requestHandle).Enable = false
		log.Infoln(handle.(*requestHandle).Name + "插件已被禁用")
	}
	handle, ok = NoticeHandles.get(id)
	if ok {
		handle.(*noticeHandle).Enable = false
		log.Infoln(handle.(*noticeHandle).Name + "插件已被禁用")
	}
	handle, ok = MetaHandles.get(id)
	if ok {
		handle.(*metaHandle).Enable = false
		log.Infoln(handle.(*metaHandle).Name + "插件已被禁用")
	}
	handle, ok = ConnectHandles.get(id)
	if ok {
		handle.(*connectHandle).Enable = false
		log.Infoln(handle.(*connectHandle).Name + "插件已被禁用")
	}
	handle, ok = DisConnectHandles.get(id)
	if ok {
		handle.(*disConnectHandle).Enable = false
		log.Infoln(handle.(*disConnectHandle).Name + "插件已被禁用")
	}

}

// BanPluginByName
/**
 * @Description: 禁用某个插件
 * @param name  插件名
 * @return int  禁用的插件个数
 */
func BanPluginByName(name string) int {
	num := 0
	for _, handle := range PretreatmentHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	for _, handle := range CommandHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	for _, handle := range MessageHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	for _, handle := range RequestHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	for _, handle := range NoticeHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	for _, handle := range MetaHandles {
		if handle.Name == name && handle.Enable {
			handle.Enable = false
			num++
		}
	}
	return num
}

func StartPluginByID(id string) {
	handle, ok := CommandHandles.get(id)
	if ok {
		handle.(*commandHandle).Enable = true
		log.Infoln(handle.(*commandHandle).Name + "插件已被启用")
	}
	handle, ok = MessageHandles.get(id)
	if ok {
		handle.(*messageHandle).Enable = true
		log.Infoln(handle.(*messageHandle).Name + "插件已被启用")
	}
	handle, ok = RequestHandles.get(id)
	if ok {
		handle.(*requestHandle).Enable = true
		log.Infoln(handle.(*requestHandle).Name + "插件已被启用")
	}
	handle, ok = NoticeHandles.get(id)
	if ok {
		handle.(*noticeHandle).Enable = true
		log.Infoln(handle.(*noticeHandle).Name + "插件已被启用")
	}
	handle, ok = MetaHandles.get(id)
	if ok {
		handle.(*metaHandle).Enable = true
		log.Infoln(handle.(*metaHandle).Name + "插件已被启用")
	}
	handle, ok = ConnectHandles.get(id)
	if ok {
		handle.(*connectHandle).Enable = true
		log.Infoln(handle.(*connectHandle).Name + "插件已被启用")
	}
	handle, ok = DisConnectHandles.get(id)
	if ok {
		handle.(*disConnectHandle).Enable = true
		log.Infoln(handle.(*disConnectHandle).Name + "插件已被启用")
	}
}

// StartPlugin
/**
 * @Description: 启用某个插件
 * @param name	插件名
 * @return int  启用插件个数
 */
func StartPlugin(name string) int {
	num := 0
	for _, handle := range PretreatmentHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	for _, handle := range CommandHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	for _, handle := range MessageHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	for _, handle := range RequestHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	for _, handle := range NoticeHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	for _, handle := range MetaHandles {
		if handle.Name == name && !handle.Enable {
			handle.Enable = true
			num++
		}
	}
	return num

}
