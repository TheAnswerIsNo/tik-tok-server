package main

import (
	"tik-tok-server/bootstrap"
	"tik-tok-server/global"
	"tik-tok-server/router"
)

func main() {

	//初始化配置参数
	bootstrap.InitializeConfig()

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	//初始化redis
	global.App.Redis = bootstrap.InitializeRedis()
	global.App.Log.Info("redis init success!")

	//初始化数据库
	global.App.DB = bootstrap.InitalizeDB()
	global.App.Log.Info("mysql init success!")

	//关闭数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	router.RunServer()
}
