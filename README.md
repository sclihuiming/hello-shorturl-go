# short url

将golang用于实际项目-阶段1

生成的文件结构如下：

```text
etc
└── shorturl-api.yaml         // 配置文件
internal
├── config
│   └── config.go             // 定义配置
├── handler
│   ├── expandhandler.go      // 实现expandHandler
│   ├── routes.go             // 定义路由处理
│   └── shortenhandler.go     // 实现shortenHandler
├── logic
│   ├── expandlogic.go        // 实现ExpandLogic
│   └── shortenlogic.go       // 实现ShortenLogic
├── svc
│   └── servicecontext.go     // 定义ServiceContext
└── types
    └── types.go              // 定义请求、返回结构体
shorturl.api
shorturl.go                   // main入口定义
```
