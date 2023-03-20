package main

import (
	"github.com/huoxue1/leafbot/driver/cqhttp_default_driver"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"

	_ "github.com/Mrs4s/go-cqhttp/db/leveldb" // leveldb

	// _ "github.com/Mrs4s/go-cqhttp/modules/pprof" // pprof 性能分析

	_ "github.com/Mrs4s/go-cqhttp/modules/silk" // si
)

func init() {
	plugin := leafbot.NewPlugin("测试")
	plugin.OnCommand("测试", leafbot.Option{
		Weight: 0,
		Block:  false,
		Allies: nil,
		Rules: []leafbot.Rule{func(ctx *leafbot.Context) bool {
			return true
		}},
	}).Handle(func(ctx *leafbot.Context) {
		ctx.Send(message.Text("123"))
	})
	plugin.OnStart("开头").Handle(func(ctx *leafbot.Context) {
		ctx.Send(message.Text("onStart匹配成功"))
	})
	plugin.OnEnd("结束").Handle(func(ctx *leafbot.Context) {
		ctx.Send("onEnd匹配成功")
	})
	plugin.OnRegex(`我的(.*?)时小明`).Handle(func(ctx *leafbot.Context) {
		log.Infoln(ctx.State.RegexResult)
		ctx.Send(message.Text("正则匹配成功"))
	})
}

func main() {
	// 创建一个驱动
	driver := cqhttp_default_driver.NewDriver()
	// 注册驱动
	leafbot.LoadDriver(driver)
	// 初始化Bot
	leafbot.InitBots(leafbot.DefaultConfig)
	// 运行驱动
	driver.Run()
}

func runChild() error {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd.Wait()
}
