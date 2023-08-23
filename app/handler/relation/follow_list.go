package relation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"tik-tok-server/app/models"
	"tik-tok-server/app/service/relation"
)

// FollowListResponse 表示关注列表响应结构
type FollowListResponse struct {
	models.CommonResponse
	*relation.FollowList
}

// QueryFollowListHandler 处理查询关注列表的请求
func QueryFollowListHandler(c *gin.Context) {
	NewProxyQueryFollowList(c).Do()
}

// ProxyQueryFollowList 代理查询关注列表的处理
type ProxyQueryFollowList struct {
	*gin.Context

	userId int64

	*relation.FollowList
}

// NewProxyQueryFollowList 创建一个 ProxyQueryFollowList 实例
func NewProxyQueryFollowList(context *gin.Context) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Context: context}
}

// Do 执行查询关注列表的操作
func (p *ProxyQueryFollowList) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendError(err.Error())
		return
	}
	p.SendOk("请求成功")
}

// parseNum 解析用户ID
func (p *ProxyQueryFollowList) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

// prepareData 准备关注列表数据
func (p *ProxyQueryFollowList) prepareData() error {
	list, err := relation.QueryFollowList(p.userId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}

// SendError 发送错误响应
func (p *ProxyQueryFollowList) SendError(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg},
	})
}

// SendOk 发送成功响应
func (p *ProxyQueryFollowList) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0, StatusMsg: msg},
		FollowList:     p.FollowList,
	})
}
