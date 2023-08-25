package user

import (
	"errors"
	"strconv"
	"tik-tok-server/app/common/request"
	"tik-tok-server/app/models/user"
	"tik-tok-server/global"
	"tik-tok-server/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, newUser user.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&user.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	newUser = user.User{UserName: params.UserName, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&newUser).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *user.User) {
	err = global.App.DB.Where("username = ?", params.UserName).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user user.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
