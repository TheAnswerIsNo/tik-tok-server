package relation

import (
	"tik-tok-server/app/cache"
	"tik-tok-server/app/models/relation"
)

// FollowerList 包含用户关注者列表的结构
type FollowerList struct {
	UserList []*relation.UserInfo `json:"user_list"`
}

// QueryFollowerList 查询用户关注者列表
func QueryFollowerList(userId int64) (*FollowerList, error) {
	return NewQueryFollowerListFlow(userId).Do()
}

// QueryFollowerListFlow 查询用户关注者列表的流程
type QueryFollowerListFlow struct {
	userId int64

	userList []*relation.UserInfo

	*FollowerList
}

// NewQueryFollowerListFlow 创建一个查询用户关注者列表的流程
func NewQueryFollowerListFlow(userId int64) *QueryFollowerListFlow {
	return &QueryFollowerListFlow{userId: userId}
}

// Do 执行查询用户关注者列表的操作
func (q *QueryFollowerListFlow) Do() (*FollowerList, error) {
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
	return q.FollowerList, nil
}

// checkNum 检查用户是否存在
func (q *QueryFollowerListFlow) checkNum() error {
	// 使用数据访问对象检查用户是否存在
	if !relation.NewUserDao().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

// prepareData 准备关注者列表数据
func (q *QueryFollowerListFlow) prepareData() error {
	// 从数据访问对象获取用户的关注者列表
	err := relation.NewUserDao().GetFollowerListByUserId(q.userId, &q.userList)
	if err != nil {
		return err
	}


	//填充每个用户的关注状态
	for _, v := range q.userList {
		v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(q.userId, v.UserId) //更新关注状态
	}
	return nil
}

// packData 打包关注者列表数据
func (q *QueryFollowerListFlow) packData() error {
	q.FollowerList = &FollowerList{UserList: q.userList}
	return nil
}
