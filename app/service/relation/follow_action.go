package relation

import "tik-tok-server/app/models/relation"

const (
	FOLLOW = 1
	CANCEL = 2
)

// 自定义错误
var ()

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
	if err = fc.publish(); err != nil {
		return err
	}
	return nil
}

// 检查
func (fc *FollowActionFlow) check() error {
	//	关注用户已注销
	if !relation.NewUserDao().IsExistById(fc.followedId) {
		return error
	}
	return nil
}
