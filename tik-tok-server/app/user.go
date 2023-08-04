package app

import (
	"tik-tok-server/app/common/request"
	"tik-tok-server/app/common/response"
	services "tik-tok-server/app/service"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BusinessFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
