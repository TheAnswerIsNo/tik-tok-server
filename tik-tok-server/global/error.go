package global

type CustomError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{-1, "业务逻辑错误"},
	ValidateError: CustomError{42200, "请求参数错误"},
	TokenError:    CustomError{-2, "登录授权失效"},
}
