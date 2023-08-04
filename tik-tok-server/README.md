# tik-tok-server
青训营大项目:极简抖音

## 1.项目结构
    tik-tok-server
    ├── app
    │   ├── common                公共模块
    │   ├── handler               控制层(处理请求数据)
    │   ├── middleware            中间件(实现jwt授权)
    │   ├── models                应用持久层(对数据库的CURD)
    │   └── service               业务逻辑层
    ├── bootstrap                 初始化配置
    ├── config                    配置结构体
    ├── global                    全局配置
    ├── logs                      日志
    ├── public                    静态资源
    ├── router                    路由分发
    └── utils                     工具包
## 2.运行项目

$ go mod tidy

修改config.yaml配置文件

$ go build