package models

import (
	"log"
	"sync"
	"tik-tok-server/bootstrap"
)

// 创建user结构体
type UserInfo struct {
	Id             int32       `json:"id,omitempty" :"id"`
	Name           string      `json:"name,omitempty" :"name"`
	sex            bool        `json:"sex,omitempty" :"sex"`
	getLike        int64       `json:"get_like,omitempty" :"get_like"`
	Friends        int64       `json:"friends,omitempty" :"friends"`
	Follows        []*UserInfo `json:"follow,omitempty" :"follows"`
	Followed       []*UserInfo `json:"followed,omitempty" :"followed"`
	Videos         []*Video    `:"videos"`
	FavoriteVideos []*Video    `:"favorite_videos"`
	Comments       []*Comment  `:"comments"`
}

type UserDao struct {
}

var (
	userOnce sync.Once
	userDao  *UserDao
)

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = new(UserDao)
	})
	return userDao
}

// 通过Id查找用户是否存在
func (user *UserDao) IsExistById(id int64) bool {
	var userInfo UserInfo
	db := bootstrap.InitalizeDB()
	if err := db.Where("id = ?", id).Select("id").First(&userInfo).Error; err != nil {
		log.Println(err)
	}
	if userInfo.Id == 0 {
		return false
	}
	return true
}
