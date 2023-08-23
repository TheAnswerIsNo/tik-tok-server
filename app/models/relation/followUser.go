package relation

import (
	"errors"
	"log"
	"sync"
	"tik-tok-server/app/models"
	"tik-tok-server/app/models/comment"
	"tik-tok-server/app/models/video"
	"tik-tok-server/global"

	"gorm.io/gorm"
)

// 创建user结构体
type UserInfo struct {
	UserId         int64              `json:"user_id,omitempty" gorm:"primaryKey"`
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

// 定义一些通用的错误
var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

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

// 根据用户ID查询用户信息
func (u *UserDao) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	// 从数据库查询用户信息
	global.App.DB.Where("user_id=?", userId).Select([]string{"user_id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	// 如果用户ID为零值，表示查询失败
	if userinfo.UserId == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// 根据用户ID判断用户是否存在
func (u *UserDao) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := global.App.DB.Where("user_id=?", id).Select("user_id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.UserId == 0 {
		return false
	}
	return true
}

// 根据用户ID获取关注列表
func (u *UserDao) GetFollowListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	// 从数据库查询用户的关注列表
	if err = global.App.DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].UserId == 0 {
		return ErrEmptyUserList
	}
	return nil
}

// 根据用户ID获取粉丝列表
func (u *UserDao) GetFollowerListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	// 从数据库查询用户的粉丝列表
	if err = global.App.DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id = ? AND r.user_info_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	return nil
}
