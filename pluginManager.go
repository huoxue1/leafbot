package leafBot

import log "github.com/sirupsen/logrus"

// GetHandleList
/**
 * @Description: 获取所有的未被禁用handle
 * @return map[string][]string
 */

type BaseHandle struct {
	Id         string
	Name       string
	Enable     bool
	IsAdmin    bool
	HandleType string
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
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["预处理响应器"] = preList

	for _, handle := range CommandHandles {

		cmdList = append(cmdList, BaseHandle{
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["command响应器"] = cmdList

	for _, handle := range MessageHandles {

		msgList = append(msgList, BaseHandle{
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["message响应器"] = msgList

	for _, handle := range RequestHandles {

		reqList = append(reqList, BaseHandle{
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["request响应器"] = reqList

	for _, handle := range NoticeHandles {

		notList = append(notList, BaseHandle{
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["notice响应器"] = notList

	for _, handle := range MetaHandles {

		metaList = append(metaList, BaseHandle{
			Id:         handle.Id,
			Name:       handle.Name,
			Enable:     handle.Enable,
			IsAdmin:    handle.IsAdmin,
			HandleType: handle.HandleType,
		})
	}
	list["meta响应器"] = metaList
	return list
}

// BanPluginById
/**
 * @Description: 根据id禁用插件
 * @param id
 */
func BanPluginById(id string) {
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

func StartPluginById(id string) {
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
