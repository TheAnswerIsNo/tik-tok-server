package user

import (
	"strconv"
	"tik-tok-server/app/models"
)

type User struct {
	models.ID
	UserName string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
