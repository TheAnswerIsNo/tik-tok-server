package userinfo

import (
	"errors"
	"log"
	"sync"
	"tik-tok-server/app/models"
	"tik-tok-server/app/models/comment"
	"tik-tok-server/app/models/video"
	"tik-tok-server/bootstrap"
)

//后面改成真正的用户信息！！！
// type UserInfo struct {
// 	Id              int64  `json:"id" gorm:"id,omitempty"`                             //用户id
// 	Name            string `json:"name" gorm:"name,omitempty"`                         //用户名称
// 	FollowCount     int64  `json:"follow_count" gorm:"follow_count,omitempty"`         //关注总数
// 	FollowerCount   int64  `json:"follower_count" gorm:"follower_count,omitempty"`     //粉丝总数
// 	IsFollow        bool   `json:"is_follow" gorm:"is_follow,omitempty"`               //true-已关注，false-未关注
// 	Avatar          string `json:"avatar" gorm:"avatar,omitempty"`                     //用户头像
// 	BackgroundImage string `json:"background_image" gorm:"background_image,omitempty"` //用户个人页顶部大图
// 	Signature       string `json:"signature" gorm:"signature,omitempty"`               //个人简介
// 	TotalFavorited  int64  `json:"total_favorited" gorm:"total_favorited,omitempty"`   //获赞数量
// 	WorkCount       int64  `json:"work_count" gorm:"work_count,omitempty"`             //作品数量
// 	FavoriteCount   int64  `json:"favorite_count" gorm:"favorite_count,omitempty"`     //点赞数量
// }

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

// 定义一些通用的错误
var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

// UserInfoDAO 表示用户信息数据访问对象
type UserInfoDAO struct {
}

// 单例模式，创建一个 UserInfoDAO 实例
var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

// 根据用户ID查询用户信息
func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	// 从数据库查询用户信息
	bootstrap.Db.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	// 如果用户ID为零值，表示查询失败
	if userinfo.UserId == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// 根据用户ID判断用户是否存在
func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := bootstrap.Db.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.UserId == 0 {
		return false
	}
	return true
}

// 根据用户ID获取关注列表
func (u *UserInfoDAO) GetFollowListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	// 从数据库查询用户的关注列表
	if err = bootstrap.Db.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].UserId == 0 {
		return ErrEmptyUserList
	}
	return nil
}

// 根据用户ID获取粉丝列表
func (u *UserInfoDAO) GetFollowerListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	// 从数据库查询用户的粉丝列表
	if err = bootstrap.Db.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id = ? AND r.user_info_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	return nil
}
