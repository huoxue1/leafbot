package leafBot

// GetHandleList
/**
 * @Description: 获取所有的未被禁用handle
 * @return map[string][]string
 */

type baseHandle struct {
	Id      string
	Name    string
	Enable  bool
	IsAdmin bool
}

func GetHandleList() map[string][]baseHandle {
	var (
		list     = make(map[string][]baseHandle)
		preList  []baseHandle
		cmdList  []baseHandle
		msgList  []baseHandle
		reqList  []baseHandle
		notList  []baseHandle
		metaList []baseHandle
	)

	for _, handle := range PretreatmentHandles {
		if !handle.Enable {
			continue
		}
		preList = append(preList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["预处理响应器"] = preList

	for _, handle := range CommandHandles {
		if !handle.Enable {
			continue
		}
		cmdList = append(cmdList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["command响应器"] = cmdList

	for _, handle := range MessageHandles {
		if !handle.Enable {
			continue
		}
		msgList = append(msgList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["message响应器"] = msgList

	for _, handle := range RequestHandles {
		if !handle.Enable {
			continue
		}
		reqList = append(reqList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["request响应器"] = reqList

	for _, handle := range NoticeHandles {
		if !handle.Enable {
			continue
		}
		notList = append(notList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["notice响应器"] = notList

	for _, handle := range MetaHandles {
		if !handle.Enable {
			continue
		}
		metaList = append(metaList, baseHandle{
			Id:      handle.Id,
			Name:    handle.Name,
			Enable:  handle.Enable,
			IsAdmin: handle.IsAdmin,
		})
	}
	list["meta响应器"] = metaList
	return list
}

// BanPlugin
/**
 * @Description: 禁用某个插件
 * @param name  插件名
 * @return int  禁用的插件个数
 */
func BanPlugin(name string) int {
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
