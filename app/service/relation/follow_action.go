package relation

import (
	"errors"
	"tik-tok-server/app/models/relation"
)

const (
	FOLLOW = 1
	CANCEL = 2
)

// 自定义错误
var (
	ErrInvalidAct  = errors.New("未定义操作")
	ErrInavlidUser = errors.New("关注用户不存在")
)

type FollowActionFlow struct {
	userId     int64
	followedId int64
	actionType int
}

// 创建实例
func NewFollowActionFlow(userId int64, followedId int64, actionType int) *FollowActionFlow {
	return &FollowActionFlow{
		userId:     userId,
		followedId: followedId,
		actionType: actionType,
	}
}

func (fc *FollowActionFlow) Do() error {
	var err error
	if err = fc.check(); err != nil {
		return err
	}
	if err = fc.packData(); err != nil {
		return err
	}
	return nil
}

// 检查
func (fc *FollowActionFlow) check() error {
	//	关注用户已注销
	if !relation.NewUserDao().IsExistById(fc.followedId) {
		// 待解决，如何自定义错误
		return ErrInavlidUser
	}
	if fc.actionType != FOLLOW && fc.actionType != CANCEL {
		return ErrInvalidAct
	}
	// 自己关注自己是否需要
	return nil
}

func (fc *FollowActionFlow) packData() error {
	userDao := relation.NewUserDao()
	var err error
	switch fc.actionType {
	case FOLLOW:
		err = userDao.AddUserFollow(fc.userId, fc.followedId)
	case CANCEL:
		err = userDao.CancelUserFollow(fc.userId, fc.followedId)
	default:
		return ErrInvalidAct
	}
	return err
}
