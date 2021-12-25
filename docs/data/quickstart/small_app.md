# 最小实例

## 新建项目
+ 新建文件夹example
+ 命令行输入
+ ```cd example && go mod init example```
+ 使用开发工具打开example项目，此时项目下应该有一个```go.mod```文件

## 安装leafBot依赖
在命令行输入

```go get github.com/huoxue1/leafBot```

此时查看go.mod文件，应该有类似如下内容
```go
module example

go 1.17

require (
	github.com/PuerkitoBio/goquery v1.7.0 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/guonaihong/gout v0.2.6 // indirect
	github.com/hjson/hjson-go v3.1.0+incompatible // indirect
	github.com/huoxue1/gg v1.3.1-0.20210909022355-795dba57682a // indirect
	github.com/huoxue1/leafBot v1.1.5 // indirect
	github.com/huoxue1/lorca v0.1.11 // indirect
	github.com/huoxue1/test3 v0.0.0-20210921063422-c43e6876777d // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mxschmitt/playwright-go v0.1100.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/tidwall/gjson v1.7.5 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/ugorji/go/codec v1.2.5 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

```

## 创建```main.go```文件
添加进去一下代码

```go
package main

import (
	"github.com/huoxue1/leafBot"
	"github.com/huoxue1/leafBot/cqhttp_reverse_ws_driver"
	"github.com/huoxue1/leafBot/message"
)

func init() {
	plugin := leafBot.NewPlugin("echo")
	plugin.SetHelp(map[string]string{"echo":"echo the params"})
	plugin.OnCommand("echo").AddHandle(func(ctx *leafBot.Context) {
		ctx.Send(message.Text(state.Args[:]))
	})
}

func main() {
	driver := cqhttp_reverse_ws_driver.NewDriver()
	leafBot.LoadDriver(driver)
	leafBot.InitBots()
	driver.Run()
}

```
## 解读

+ 第1行```package main```为包名，代表当前包可以为启动包
+ 第3-7行为导入项目依赖
+ 第9行为一个init方法，该方法为go语言中特殊方法，会在初始化项目的时候最先调用
+ 第10行为使用leafBot创建一个名为echo的插件
+ 第11行为该插件添加一个帮助
+ 第12行添加一个commang响应器，响应echo
+ 第13行通过event.send发送文本内容，内容为附带的参数
+ 第17行为项目入口方法
+ 第18行初始化了一个```cqhttp_reverse_ws_driver```驱动，驱动内容可以查看[驱动](../driver.md)
+ 第19行为leafBot加载了该驱动
+ 第20行初始化leafBot的配置内容
+ 第21行运行该驱动

## 运行

+ ### 编译

> go build main.go

+ ### 启动

> ./main.exe

此时可以看待日志输出
```go
time="2021-10-30T09:54:27+08:00" level=info msg="配置文件加载失败或者不存在"
time="2021-10-30T09:54:27+08:00" level=info msg="将生成默认配置文件"
time="2021-10-30T09:54:27+08:00" level=info msg="请输入机器人账号"
```
此时在命令行输入bot的账号,程序挺住运行停止后发现自动在config目录生成了```config.yml```配置文件，更多配置文件信息参考[配置](../config.md)

+ ### 再次运行

发现控制台输出如下内容
```go
[2021-10-30 09:57:10] [leafBot] [INFO] : 
 _                 __ ____        _
| |               / _|  _ \      | |
| |     ___  __ _| |_| |_) | ___ | |_
| |    / _ \/ _` |  _|  _ < / _ \| __|
| |___|  __/ (_| | | | |_) | (_) | |_
|______\___|\__,_|_| |____/ \___/ \__| 
[2021-10-30 09:57:10] [leafBot] [INFO] : web页面：http://127.0.0.1:3000 
[2021-10-30 09:57:10] [leafBot] [INFO] : 已加载插件 ==》 管理管理 
[2021-10-30 09:57:10] [leafBot] [INFO] : 已加载插件 ==》 配置重载 
[2021-10-30 09:57:10] [leafBot] [INFO] : 已加载插件 ==》 echo 
```

## 配置协议段上报

建议协议段使用go-cqhttp
驱动与协议段配置对应列表参考[驱动](../driver.md)