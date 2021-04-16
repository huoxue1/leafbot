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

### 1. 安装go-cqhttp

下载地址 :<https://github.com/Mrs4s/go-cqhttp/releases>

ps: 建议下载最新版本

### 2.配置go-cqhttp

```yaml
# go-cqhttp 默认配置文件

account: # 账号相关
  uin: 1233456 # QQ账号
  password: '' # 密码为空时使用扫码登录
  encrypt: false  # 是否开启密码加密
  relogin:        # 重连设置
    disabled: false
    delay: 3      # 重连延迟, 单位秒
    interval: 0   # 重连间隔
    max-times: 0  # 最大重连次数, 0为无限制

  # 是否使用服务器下发的新地址进行重连
  # 注意, 此设置可能导致在海外服务器上连接情况更差
  use-sso-address: true

heartbeat:
  disabled: false # 是否开启心跳事件上报
  # 心跳频率, 单位秒
  # -1 为关闭心跳
  interval: 5

message:
  # 上报数据类型
  # 可选: string,array
  post-format: string
  # 是否忽略无效的CQ码, 如果为假将原样发送
  ignore-invalid-cqcode: false
  # 是否强制分片发送消息
  # 分片发送将会带来更快的速度
  # 但是兼容性会有些问题
  force-fragment: false
  # 是否将url分片发送
  fix-url: false
  # 下载图片等请求网络代理
  proxy-rewrite: ''
  # 是否上报自身消息
  report-self-message: true
  # 移除服务端的Reply附带的At
  remove-reply-at: false
  # 为Reply附加更多信息
  extra-reply-data: false

output:
  # 日志等级 trace,debug,info,warn,error日志等级 trace,debug,info,warn,error
  log-level: warn
  # 是否启用 DEBUG
  debug: false # 开启调试模式

# 默认中间件锚点
default-middlewares: &default
  # 访问密钥, 强烈推荐在公网的服务器设置
  access-token: ''
  # 事件过滤器文件目录
  filter: ''
  # API限速设置
  # 该设置为全局生效
  # 原 cqhttp 虽然启用了 rate_limit 后缀, 但是基本没插件适配
  # 目前该限速设置为令牌桶算法, 请参考:
  # https://baike.baidu.com/item/%E4%BB%A4%E7%89%8C%E6%A1%B6%E7%AE%97%E6%B3%95/6597000?fr=aladdin
  rate-limit:
    enabled: false # 是否启用限速
    frequency: 1  # 令牌回复频率, 单位秒
    bucket: 1     # 令牌桶大小

servers:
  # HTTP 通信设置
  - http:
      # 是否关闭正向HTTP服务器
      disabled: true
      # 服务端监听地址
      host: 127.0.0.1
      # 服务端监听端口
      port: 5700
      # 反向HTTP超时时间, 单位秒
      # 最小值为5，小于5将会忽略本项设置
      timeout: 5
      middlewares:
        <<: *default # 引用默认中间件
      # 反向HTTP POST地址列表
      post:
        #- url: '' # 地址
        #  secret: ''           # 密钥
        #- url: 127.0.0.1:5701 # 地址
        #  secret: ''          # 密钥

  # 正向WS设置
  - ws:
      # 是否禁用正向WS服务器
      disabled: true
      # 正向WS服务器监听地址
      host: 127.0.0.1
      # 正向WS服务器监听端口
      port: 6700
      middlewares:
        <<: *default # 引用默认中间件

  - ws-reverse:
      # 是否禁用当前反向WS服务
      disabled: false
      # 反向WS Universal 地址
      # 注意 设置了此项地址后下面两项将会被忽略
      universal: ws://127.0.0.1:8080/cqhttp/ws
      # 反向WS API 地址
      api: ws://your_websocket_api.server
      # 反向WS Event 地址
      event: ws://your_websocket_event.server
      # 重连间隔 单位毫秒
      reconnect-interval: 3000
      middlewares:
        <<: *default # 引用默认中间件

  # 可添加更多
  #- ws-reverse:
  #- ws:
  #- http:

database: # 数据库相关设置
  leveldb:
    # 是否启用内置leveldb数据库
    # 启用将会增加10-20MB的内存占用和一定的磁盘空间
    # 关闭将无法使用 撤回 回复 get_msg 等上下文相关功能
    enable: true

```

复制该配置文件覆盖go-cqhttp文件夹下config.yaml文件

然后再命令行运行 
```shell
./go-cqhttp.exe
```

按照提示使用qq扫码登录


### 3. 下载<https://github.com/3343780376/leafBot/tree/master/example>
  双击运行


## 进阶自己构建，

### 1. 安装golang环境 

golang环境下载<https://studygolang.com/dl>

安装教程参考<https://www.runoob.com/go/go-environment.html>
    
### 2. 新建golang项目
创建main.go文件

复制下面代码

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