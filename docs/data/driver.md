# driver列表

## cqhttp-http-driver
通过http通信方式与onebot端进行通信。

gocq的host和port配置分别对应leafBot的post_host和post_port

gocq的url配置分别对应leafBot的listen_host和listen_port

> 使用http链接方式可能会使oncnnect和ondisconnect插件失效
### gocq关键配置

```yaml
servers:
  # HTTP 通信设置
  - http:
      # 是否关闭正向HTTP服务器
      disabled: false
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
        - url: http://127.0.0.1:8081 # 地址
        #  secret: ''          # 密钥

```

### leafBot关键配置

```yaml
listen_host: 127.0.0.1

listen_port: 8081

# cqhttp_http_driver相关配置
web_hook:
    -
        post_host: 127.0.0.1
        post_port: 5700
        self_id: 1603214019
```


## cqhttp-positive-ws-driver
gocq作为服务端，leafBot最为客户端，通过websocket的链接方式进行通信
>使用正向ws会导致不能一个leafBot链接多个gocq

### gocq关键配置

```yaml

# 正向WS设置
- ws:
    # 是否禁用正向WS服务器
    disabled: true
    # 正向WS服务器监听地址
    host: 127.0.0.1
    # 正向WS服务器监听端口
    port: 8080
    middlewares:
      <<: *default # 引用默认中间件

```

### leafBot关键配置

```yaml
# bot运行地址，若和gocq在同一台机器则只需要填写127.0.0.1即可，否则填写0.0.0.0，gocq配置你的公网地址
host: 127.0.0.1
# bot运行端口
port: 8080
```

## cqhttp-reverse-ws-driver
gocq作为客户端，leafBot作为服务的，通过websocket的链接方式进行通信
>推荐使用方式，可支持多q

### gocq关键配置

```yaml

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

```

### leafBot关键配置

```yaml
# bot运行地址，若和gocq在同一台机器则只需要填写127.0.0.1即可，否则填写0.0.0.0，gocq配置你的公网地址
host: 127.0.0.1
# bot运行端口
port: 8080
```

## 自定义实现
leafBot定义了一套关于driver的接口

```go
// Driver
// @Description:
//
type Driver interface {
	// Run
	// @Description: 运行该驱动的接口，该接口应该为阻塞式运行
	//
	Run()
	// GetEvent
	// @Description: 返回一个chan，该chan为事件传递的chan
	// @return chan
	//
	GetEvent() chan []byte

	OnConnect(func(selfId int64, host string, clientRole string))
	OnDisConnect(func(selfId int64))

	// GetBot
	// @Description: 获取一个实现了APi接口的bot
	// @param int64 bot的id
	// @return interface{}
	//
	GetBot(int64) interface{}
	// GetBots
	// @Description: 获取所有bot
	// @return map[int64]interface{}
	//
	GetBots() map[int64]interface{}
	
	// 设置一些配置信息
	SetConfig(config map[string]interface{})
	// 添加一个webhook监听
	AddWebHook(selfID int64, postHost string, postPort int)
	// 配置access——token
	SetToken(token string)
}
```
实现该接口即可作为leafBot的driver进行加载