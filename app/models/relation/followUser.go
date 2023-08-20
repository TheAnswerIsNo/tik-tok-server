package relation

import (
	"gorm.io/gorm"
	"log"
	"sync"
	"tik-tok-server/app/models"
	"tik-tok-server/app/models/comment"
	"tik-tok-server/app/models/video"
	"tik-tok-server/global"
)

// 创建user结构体
type UserInfo struct {
	UserId         int32              `json:"user_id,omitempty" gorm:"primaryKey"`
	Name           string             `json:"name,omitempty"`
	FollowCount    int64              `json:"follow_count" gorm:"column:follow_count"`
	FollowerCount  int64              `json:"follower_count" gorm:"column:follower_count"`
	IsFollow       bool               `json:"is_follow" gorm:"column:is_follow"`
	User           *models.User       `json:"user"`
	Follows        []*UserInfo        `json:"follows,omitempty" gorm:"many2many:user_relation;joinForeignKey:user_id;joinReferences:user_id"`
	Videos         []*video.Video     `json:"videos,omitempty"`
	FavoriteVideos []*video.Video     `json:"favorite_videos,omitempty" gorm:"many2many:user_favor_videos;"`
	Comments       []*comment.Comment `json:"comments,omitempty"`
}

// 生成对应表
func (UserInfo) TableName() string {
	return "user_info"
}

// Dao结构体 用于封装数据库操作
type UserDao struct {
}

var (
	userOnce sync.Once
	userDao  *UserDao
)

// 实例化函数
func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = new(UserDao)
	})
	return userDao
}

// 通过Id查找用户是否存在
func (*UserDao) IsExistById(id int64) bool {
	var userInfo UserInfo
	db := global.App.DB
	if err := db.Where("user_id = ?", id).Select("id").First(&userInfo).Error; err != nil {
		log.Println(err)
	}
	if userInfo.UserId == 0 {
		return false
	}
	return true
}

// 用户关注
// userId: 当前用户id，followedId:被关注的对象的id
func (user *UserDao) AddUserFollow(userId, followedId int64) error {
	return global.App.DB.Transaction(func(tx *gorm.DB) error {
		// 更新当前用户的关注数
		if err := tx.Exec("UPDATE user_info SET follow_count = follow_count + 1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		// 更新被关注用户的粉丝数
		if err := tx.Exec("UPDATE user_info SET follower_count = follower_count + 1 WHERE id = ?", followedId).Error; err != nil {
			return err
		}
		//	更新用户关系
		if err := tx.Exec("INSERT INTO `user_relation` (`user_info_id`, `follow_id`) VALUES (?, ?)", userId, followedId).Error; err != nil {
			return err
		}
		return nil
	})
}

// 取消关注
func (user *UserDao) CancelUserFollow(userId, followedId int64) error {
	return global.App.DB.Transaction(func(tx *gorm.DB) error {
		// 更新当前用户的关注数
		if err := tx.Exec("UPDATE user_info SET follow_count = follow_count - 1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		// 更新被关注用户的粉丝数
		if err := tx.Exec("UPDATE user_info SET follower_count = follower_count - 1 WHERE id = ?", followedId).Error; err != nil {
			return err
		}
		//	更新用户关系
		if err := tx.Exec("DELETE FROM `user_relation` (`user_info_id`, `follow_id`) VALUES (?, ?)", userId, followedId).Error; err != nil {
			return err
		}
		return nil
	})
}
