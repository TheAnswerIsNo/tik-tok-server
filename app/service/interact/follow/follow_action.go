package follow

import "errors"

const (
	FOLLOW = 0
	CANCEL = 1
)

// 定义错误
var (
	ErrIvalidAction = errors.New("未定义操作")
	ErrIvalidUser   = errors.New("用户不存在")
)
