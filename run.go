package leafBot

import (
	"fmt" //nolint:gci
	log "github.com/sirupsen/logrus"
	"os" //nolint:gci
)

// run
/*
   @Description:
   @param bot *Bot
*/
func run(bot *Bot) {
	log.Infoln("已加载bot：" + bot.Name)

}

func writeGoConfig(selfID int) {
	err := os.WriteFile("config.yml", []byte(fmt.Sprintf(gocqConfig, selfID)), 0666)
	if err != nil {
		return
	}
	log.Infoln("已生成gocq配置文件，将其复制到gocq目录即可使用")
}

var gocqConfig = "# go-cqhttp 默认配置文件\n\naccount: # 账号相关\n  uin: %d # QQ账号\n  password: '' # 密码为空时使用扫码登录\n  encrypt: false  # 是否开启密码加密\n  status: 0      # 在线状态 请参考 https://github.com/Mrs4s/go-cqhttp/blob/dev/docs/config.md#在线状态\n  relogin: # 重连设置\n    disabled: false\n    delay: 3      # 重连延迟, 单位秒\n    interval: 0   # 重连间隔\n    max-times: 0  # 最大重连次数, 0为无限制\n\n  # 是否使用服务器下发的新地址进行重连\n  # 注意, 此设置可能导致在海外服务器上连接情况更差\n  use-sso-address: true\n\nheartbeat:\n  disabled: false # 是否开启心跳事件上报\n  # 心跳频率, 单位秒\n  # -1 为关闭心跳\n  interval: 5\n\nmessage:\n  # 上报数据类型\n  # 可选: string,array\n  post-format: array\n  # 是否忽略无效的CQ码, 如果为假将原样发送\n  ignore-invalid-cqcode: false\n  # 是否强制分片发送消息\n  # 分片发送将会带来更快的速度\n  # 但是兼容性会有些问题\n  force-fragment: false\n  # 是否将url分片发送\n  fix-url: false\n  # 下载图片等请求网络代理\n  proxy-rewrite: ''\n  # 是否上报自身消息\n  report-self-message: false\n  # 移除服务端的Reply附带的At\n  remove-reply-at: false\n  # 为Reply附加更多信息\n  extra-reply-data: false\n\noutput:\n  # 日志等级 trace,debug,info,warn,error\n  log-level: warn\n  # 是否启用 DEBUG\n  debug: false # 开启调试模式\n\n# 默认中间件锚点\ndefault-middlewares: &default\n  # 访问密钥, 强烈推荐在公网的服务器设置\n  access-token: ''\n  # 事件过滤器文件目录\n  filter: ''\n  # API限速设置\n  # 该设置为全局生效\n  # 原 cqhttp 虽然启用了 rate_limit 后缀, 但是基本没插件适配\n  # 目前该限速设置为令牌桶算法, 请参考:\n  # https://baike.baidu.com/item/%E4%BB%A4%E7%89%8C%E6%A1%B6%E7%AE%97%E6%B3%95/6597000?fr=aladdin\n  rate-limit:\n    enabled: false # 是否启用限速\n    frequency: 1  # 令牌回复频率, 单位秒\n    bucket: 1     # 令牌桶大小\n\nservers:\n  # HTTP 通信设置\n  - http:\n      # 是否关闭正向HTTP服务器\n      disabled: true\n      # 服务端监听地址\n      host: 127.0.0.1\n      # 服务端监听端口\n      port: 5700\n      # 反向HTTP超时时间, 单位秒\n      # 最小值为5，小于5将会忽略本项设置\n      timeout: 5\n      middlewares:\n        <<: *default # 引用默认中间件\n      # 反向HTTP POST地址列表\n      post:\n      #- url: '' # 地址\n      #  secret: ''           # 密钥\n      #- url: 127.0.0.1:5701 # 地址\n      #  secret: ''          # 密钥\n\n  # 正向WS设置\n  - ws:\n      # 是否禁用正向WS服务器\n      disabled: true\n      # 正向WS服务器监听地址\n      host: 127.0.0.1\n      # 正向WS服务器监听端口\n      port: 6700\n      middlewares:\n        <<: *default # 引用默认中间件\n\n  - ws-reverse:\n      # 是否禁用当前反向WS服务\n      disabled: false\n      # 反向WS Universal 地址\n      # 注意 设置了此项地址后下面两项将会被忽略\n      universal: ws://127.0.0.1:8080/cqhttp/ws\n      # 反向WS API 地址\n      api: ws://your_websocket_api.server\n      # 反向WS Event 地址\n      event: ws://your_websocket_event.server\n      # 重连间隔 单位毫秒\n      reconnect-interval: 3000\n      middlewares:\n        <<: *default # 引用默认中间件\n  # pprof 性能分析服务器, 一般情况下不需要启用.\n  # 如果遇到性能问题请上传报告给开发者处理\n  # 注意: pprof服务不支持中间件、不支持鉴权. 请不要开放到公网\n  - pprof:\n      # 是否禁用pprof性能分析服务器\n      disabled: true\n      # pprof服务器监听地址\n      host: 127.0.0.1\n      # pprof服务器监听端口\n      port: 7700\n\n  # 可添加更多\n  #- ws-reverse:\n  #- ws:\n  #- http:\n  #- pprof:\n\ndatabase: # 数据库相关设置\n  leveldb:\n    # 是否启用内置leveldb数据库\n    # 启用将会增加10-20MB的内存占用和一定的磁盘空间\n    # 关闭将无法使用 撤回 回复 get_msg 等上下文相关功能\n    enable: true\n"
