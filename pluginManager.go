package leafBot

// GetHandleList
/**
 * @Description: 获取所有的未被禁用handle
 * @return map[string][]string
 */
func GetHandleList() map[string][]string {
	var (
		list     = make(map[string][]string)
		preList  []string
		cmdList  []string
		msgList  []string
		reqList  []string
		notList  []string
		metaList []string
	)

	for _, handle := range PretreatmentHandles {
		if !handle.Enable {
			continue
		}
		preList = append(preList, handle.Name)
	}
	list["预处理响应器"] = preList

	for _, handle := range CommandHandles {
		if !handle.Enable {
			continue
		}
		cmdList = append(cmdList, handle.Name)
	}
	list["command响应器"] = cmdList

	for _, handle := range MessageHandles {
		if !handle.Enable {
			continue
		}
		msgList = append(msgList, handle.Name)
	}
	list["message响应器"] = msgList

	for _, handle := range RequestHandles {
		if !handle.Enable {
			continue
		}
		reqList = append(reqList, handle.Name)
	}
	list["request响应器"] = reqList

	for _, handle := range NoticeHandles {
		if !handle.Enable {
			continue
		}
		notList = append(notList, handle.Name)
	}
	list["notice响应器"] = notList

	for _, handle := range MetaHandles {
		if !handle.Enable {
			continue
		}
		metaList = append(metaList, handle.Name)
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
