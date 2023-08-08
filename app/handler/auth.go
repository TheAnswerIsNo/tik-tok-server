package handler

import (
	"github.com/gin-gonic/gin"
	"tik-tok-server/app/common/request"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/service"
)

// Login 进行入参校验，并调用UserService和jwtService服务颁发token
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := service.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := service.JwtService.CreateToken(service.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// Info 通过 JWTAuth 中间件校验 Token 识别的用户 ID 来获取用户信息
func Info(c *gin.Context) {
	err, user := service.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
