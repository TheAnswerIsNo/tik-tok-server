package main

import (
	"tik-tok-server/bootstrap"
	"tik-tok-server/router"
)

func main() {

	//初始化配置参数
	bootstrap.InitializeConfig()

	//初始化日志
	log := bootstrap.InitializeLog()
	log.Info("log init success!")

	//初始化redis
	bootstrap.InitializeRedis()
	log.Info("redis init success!")

	//初始化数据库
	DB := bootstrap.InitalizeDB()
	log.Info("mysql init success!")

	//关闭数据库连接
	defer func() {
		if DB != nil {
			db, _ := DB.DB()
			db.Close()
		}
	}()

	router.RunServer()

}
