package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	app "tik-tok-server/app/handler"
	"tik-tok-server/app/handler/interact/comment"
	"tik-tok-server/app/handler/relation"
	"tik-tok-server/app/handler/video"
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
		publishRouter := douyin.Group("/publish")
		{
			publishRouter.POST("/action", video.PublishVideoHandler)
		}

		//这个是一个负责登录注册的模块
		registerRouter := douyin.Group("/user")
		{
			registerRouter.POST("/register/", app.Register)
			registerRouter.POST("/Login/", app.Login)
		}
		//这个是用户信息
		authRouter := douyin.Group("/auth").Use(middleware.JWTAuth(middleware.AppGuardName))
		{
			authRouter.POST("/info/", app.Info)
		}

		// 关注
		followRouter := douyin.Group("/relation")
		{
			followRouter.POST("/action", relation.FollowActionHandler)
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
