# 简洁GO项目架构

## 项目结构

```bash
.
├── README.md
├── app // 程序逻辑
│   ├── controller      // 控制层
│   │   └── user.go
│   ├── middleware      // 中间件
│   │   └── cors.go
│   ├── route.go        // 路由
│   └── service         // 服务层
│       └── user.go
├── cmd                 // 入口
│   └── go-clean-arch
│       └── server.go
├── config              // 配置   
│   └── config.yaml
├── constant            // 常用
│   └── err.go
├── doc                 // 文档说明
│   └── api.json
├── domain              // 业务领域层
│   └── user.go
├── go.mod
├── infra               // 基础设施层
│   └── mysql
│       └── user.go
├── test
└── tool
```

## 参考
- [简洁go架构](https://github.com/bxcodec/go-clean-arch)
- [uncle bob的架构思想](https://cleancoders.com/)
```