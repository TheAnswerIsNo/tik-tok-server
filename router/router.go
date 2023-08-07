package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tik-tok-server/app/handler/interact/comment"
	"tik-tok-server/global"
)

func setupRouter() *gin.Engine {

	routers := gin.Default()

	// 顶层路由/douyin
	douyin := routers.Group("/douyin")
	{
		//次级路由 根据各自模块命名
		commentRouter := douyin.Group("/comment")
		{
			commentRouter.GET("/list", comment.QueryCommentHandler)
			commentRouter.POST("/action", comment.ActionCommentHandler)
		}
	}

	//静态资源路由
	routers.Static("/public", "./public")
	return routers
}

func RunServer() {
	r := setupRouter()

	err := r.Run(fmt.Sprintf(":%d", global.App.Config.App.Port))
	if err != nil {
		return
	}
}
