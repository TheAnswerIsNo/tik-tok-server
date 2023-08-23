package relation

import (
	"errors"
	"tik-tok-server/app/models/relation"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

// FollowList 表示关注列表结构
type FollowList struct {
	UserList []*relation.UserInfo `json:"user_list"`
}

// QueryFollowList 查询用户的关注列表
func QueryFollowList(userId int64) (*FollowList, error) {
	return NewQueryFollowListFlow(userId).Do()
}

// QueryFollowListFlow 查询关注列表的处理流程
type QueryFollowListFlow struct {
	userId int64

	userList []*relation.UserInfo

	*FollowList
}

// NewQueryFollowListFlow 创建一个 QueryFollowListFlow 实例
func NewQueryFollowListFlow(userId int64) *QueryFollowListFlow {
	return &QueryFollowListFlow{userId: userId}
}

// Do 执行查询关注列表的操作
func (q *QueryFollowListFlow) Do() (*FollowList, error) {
	var err error
	if err = q.checkNum(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}

	return q.FollowList, nil
}

// checkNum 检查用户是否存在
func (q *QueryFollowListFlow) checkNum() error {
	if !relation.NewUserDao().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

// prepareData 准备关注列表数据
func (q *QueryFollowListFlow) prepareData() error {
	var userList []*relation.UserInfo
	err := relation.NewUserDao().GetFollowListByUserId(q.userId, &userList)
	if err != nil {
		return err
	}
	for i := range userList {
		userList[i].IsFollow = true // 当前用户的关注列表，因此 isFollow 设置为 true
	}
	q.userList = userList
	return nil
}

// packData 打包关注列表数据
func (q *QueryFollowListFlow) packData() error {
	q.FollowList = &FollowList{UserList: q.userList}
	return nil
}
