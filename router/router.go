package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tik-tok-server/app/handler/interact/comment"
	"tik-tok-server/app/middleware"
	"tik-tok-server/global"
)

func setupRouter() *gin.Engine {

	routers := gin.Default()

	// 顶层路由/douyin
	douyin := routers.Group("/douyin")
	{
		//次级路由 根据各自模块命名
		commentRouter := douyin.Group("/comment").Use(middleware.JWTAuth(middleware.AppGuardName))
		{
			commentHandler := new(comment.CommentHandler)
			commentRouter.GET("/list/", commentHandler.GetCommentList)
			commentRouter.POST("/action/", commentHandler.CommentAction)
		}

	}

	//静态资源路由
	routers.Static("public", "./")
	return routers
}

func RunServer() {
	r := setupRouter()

	err := r.Run(fmt.Sprintf(":%d", global.App.Config.App.Port))
	if err != nil {
		return
	}
}
