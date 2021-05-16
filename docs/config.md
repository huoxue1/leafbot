### 加载的配置文件内容为：

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
  "admin": 123,
  "host": "127.0.0.1",
  "port": 8080,
  "log_level": "info"
}
```

+ bots :一个bot数组
+ bot : 包含了name字段和self_id字段，self_id为机器人qq号
+ admin : 管理员账号
+ host: gocq的ws上报地址
+ port : gocq的ws上报端口
+ log_level : 日志等级，默认为info