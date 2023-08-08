package handler

import (
	"github.com/gin-gonic/gin"
	"tik-tok-server/app/common/request"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/service"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BusinessFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := service.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
