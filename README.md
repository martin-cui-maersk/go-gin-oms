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
├── web # 前端
├── server # 后端
├── LICENSE
```



## 运行

```sh
go run main.go # 等于 go run main.go -m=test 默认使用配置文件 config.debug.yaml
go run main.go -m=test # 使用配置文件 config.test.yaml
go run main.go -m=release # 使用配置文件 config.release.yaml
```

