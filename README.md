# go-gin-oms

用GIN框架编写的开箱即用的OMS后台



## 配置文件

```sh
# 默认使用gin的debug模式 修改对应配置
config.debug.yaml
config.test.yaml
config.release.yaml
```



## 目录结构

```sh
server
├── api                  # api接口，按版本和类型分类
│   └── v1               # v1版本
│       ├── system
│       │   ├── menu.go
│       │   ├── role.go
│       │   └── user.go
│       └── user
│           └── user.go
├── confii               # 配置文件夹
│   ├── constant.go
│   ├── jwt.go
│   ├── logconfig.go
│   ├── mysql.go
│   ├── redis.go
│   └── server.go
├── config.debug.yaml    # debug 配置文件
├── config.release.yaml  # release 配置文件
├── config.test.yaml     # test 配置文件
├── core                 # 核心类
│   ├── logger.go
│   ├── mysql.go
│   └── viper.go
├── global               # 全局变量
│   └── global.go
├── go.mod
├── go.sum
├── logs                 # 日志文件夹
├── main.go              # 入口文件
├── middleware           # 中间件
│   ├── jwt.go
│   └── logger.go
├── models               # 数据库模型
│   ├── common
│   │   └── by.go
│   ├── sys_menu.go
│   ├── sys_role.go
│   ├── sys_role_menu.go
│   └── sys_user.go
├── router              # 路由
│   ├── router.go
│   ├── v1.go
│   └── v2.go
└── utils               #工具类
└── web                 # 前端
```



## 运行

```sh
go run main.go # 等于 go run main.go -m=test 默认使用配置文件 config.debug.yaml
go run main.go -m=test # 使用配置文件 config.test.yaml
go run main.go -m=release # 使用配置文件 config.release.yaml
```

