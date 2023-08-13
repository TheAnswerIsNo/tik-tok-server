package relation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"tik-tok-server/app/common/response"
	//"tik-tok-server/app/service/relation"
)

// 代理对象
type ProxyUser struct {
	*gin.Context
	userId     int64
	followId   int64
	antionType int
}

// 提供生成对象方法
func NewProxyUser(context *gin.Context) *ProxyUser {
	return &ProxyUser{Context: context}
}

// 提供对外函数
func FollowActionHandler(context *gin.Context) {
	NewProxyUser(context).Do()
}

// 实际准备
func (pUser *ProxyUser) Do() {
	if err := pUser.prepareNum(); err != nil {
		response.ValidateFail(pUser.Context, err.Error())
		return
	}

}

// 解析参数
func (pUser *ProxyUser) prepareNum() error {
	userId, isExists := pUser.Get("user_id")
	if !isExists {
		return errors.New("userId解析出错")
	}
	pUser.userId = userId.(int64)

	//	解析被关注用户的id
	followId := pUser.Query("to_user_id") //value string
	parseFollowId, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		return err
	}
	pUser.followId = parseFollowId

	//	解析action_type
	actionType := pUser.Query("action_type")
	parseActionType, err := strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		return err
	}
	pUser.antionType = int(parseActionType)
	return nil
}

//func (pUser *ProxyUser)startAction() error {
//	//err := relation.PostFollowAction(pUser.userId, pUser.followId, pUser.antionType)
//	//if err != nil {
//	//	return err
//	//}
//	//return nil
//}
