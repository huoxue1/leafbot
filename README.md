<p align="center">
  <a href="https://ishkong.github.io/go-cqhttp-docs/"><img src="https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=2709879415,936942073&fm=26&gp=0.jpg" width="200" height="200" alt="go-cqhttp"></a>
</p>


<div align="center">

# LeafBot

_✨ 基于 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp)，使用[OneBot](https://github.com/howmanybots/onebot)标准的插件 ✨_

</div>

<p align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/golang-v1.16-brightgreen" alt="">
    </a>
  <a href="https://github.com/howmanybots/onebot/blob/master/README.md">
    <img src="https://img.shields.io/badge/OneBot-v11-blue?style=flat&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABABAMAAABYR2ztAAAAIVBMVEUAAAAAAAADAwMHBwceHh4UFBQNDQ0ZGRkoKCgvLy8iIiLWSdWYAAAAAXRSTlMAQObYZgAAAQVJREFUSMftlM0RgjAQhV+0ATYK6i1Xb+iMd0qgBEqgBEuwBOxU2QDKsjvojQPvkJ/ZL5sXkgWrFirK4MibYUdE3OR2nEpuKz1/q8CdNxNQgthZCXYVLjyoDQftaKuniHHWRnPh2GCUetR2/9HsMAXyUT4/3UHwtQT2AggSCGKeSAsFnxBIOuAggdh3AKTL7pDuCyABcMb0aQP7aM4AnAbc/wHwA5D2wDHTTe56gIIOUA/4YYV2e1sg713PXdZJAuncdZMAGkAukU9OAn40O849+0ornPwT93rphWF0mgAbauUrEOthlX8Zu7P5A6kZyKCJy75hhw1Mgr9RAUvX7A3csGqZegEdniCx30c3agAAAABJRU5ErkJggg==" alt="cqhttp">
  </a>
    <a href="#">
    <img src="https://img.shields.io/badge/FengyeBot-v1.0-orange" alt="">
    </a>
    <a href="#">
    <img src="https://img.shields.io/badge/gocqhttp-v1.0.0--beta3-blue" alt="">
    </a>
</p>


---

## 安装

```
    go get github.com/3343780376/leafBot
```

## 内置插件

+ ### /echo插件
```
    /echo 123
    回复 ：123
```
+ ### 查询网易云歌曲
```
    查询歌曲 许嵩
```
+ ### 点歌
```
    点歌 5041604
```

+ ### 每日一图
    一图  1 
  
    即返回前一天的每日一图，最大为7，默认为0
```
  一图
```

## 基础使用
    
```go
    package main

import (
	"github.com/3343780376/leafBot"
	"os"
)

func init() {
	// 为bot添加weather响应器，命令为 ”/天气“ ,allies为命令别名，
	//参数格式为一个字符串数组，rule为一个结构体，响应前会先判断所以rules为true，weight为权重，block为是否阻断
	leafBot.AddCommandHandle(Weather, "/天气", nil, nil, 10, false)
}

func main() {
	dir, _ := os.Getwd()                             // 获取当前路径
	leafBot.LoadConfig(dir + "/example/config.json") //拼接配置文件路径，并且加载配置文件
	leafBot.InitBots()                               //初始化Bot
}

/*
	event: bot的event，里面包含了事件的所有字段
	bot: 触发事件的bot指针
	args ： 命令的参数，为一个数组
*/
func Weather(event leafBot.Event, bot *leafBot.Bot, args []string) {
	m := map[string]string{"北京": "晴", "山东": "下雨"}
	// 调用发送消息的api，会根据messageType自动回复
	bot.SendMsg(event.MessageType, event.UserId, event.GroupId,
		args[0]+"的天气为"+m[args[0]],
		false)
}

```

加载的配置文件内容为：

```json
{
  "bots": [
    {
      "name": "commit",
      "self_id": 123
    },
    {
     "name": "bot1",
      "self_id": 123
    }
  ],
  "host": "127.0.0.1",
  "port": 8080
}
```

+ bots :一个bot数组
+ bot : 包含了name字段和self_id字段，self_id为机器人qq号
+ host: gocq的ws上报地址
+ port : gocq的ws上报端口