package follow

import (
	"errors"
	"tik-tok-server/app/models"
)

const (
	Follow = 0
	Cancel = 1
)

// 定义错误
var (
	ErrIvalidAction = errors.New("未定义操作")
	ErrIvalidUser   = errors.New("用户不存在")
)

// 对用户操作进行处理
func FollowAction(userId, followedId int64, actionType int) error {
	// 可能存在点赞评论等操作
	if actionType != Follow && actionType != Cancel {
		return ErrIvalidAction
	}
	if userId == followedId {
		return ErrIvalidUser
	}
	if !models.NewUserDao().IsExistById(userId) {
		return ErrIvalidUser
	}
}
