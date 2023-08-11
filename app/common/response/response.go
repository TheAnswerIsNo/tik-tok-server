package responsepackage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tik-tok-server/global"
)

// Response 响应结构体
type Response struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"data"`
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		"success",
		data,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, statusCode int, statusMsg string) {
	c.JSON(http.StatusOK, Response{
		statusCode,
		statusMsg,
		nil,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenFail token鉴权失败
func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}
