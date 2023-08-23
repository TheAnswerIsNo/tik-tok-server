package relation

import (
	"errors"
	"tik-tok-server/app/models"
	user_info2 "tik-tok-server/app/service/relation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FollowerListResponse 包含关注者列表的响应结构
type FollowerListResponse struct {
	models.CommonResponse
	*user_info2.FollowerList
}

// QueryFollowerHandler 处理查询关注者列表的请求
func QueryFollowerHandler(c *gin.Context) {
	// 创建代理查询关注者列表的处理器并执行
	NewProxyQueryFollowerHandler(c).Do()
}

// ProxyQueryFollowerHandler 代理查询关注者列表的处理器结构
type ProxyQueryFollowerHandler struct {
	*gin.Context

	userId int64

	*user_info2.FollowerList
}

// NewProxyQueryFollowerHandler 创建一个代理查询关注者列表的处理器
func NewProxyQueryFollowerHandler(context *gin.Context) *ProxyQueryFollowerHandler {
	return &ProxyQueryFollowerHandler{Context: context}
}

// Do 执行查询关注者列表的操作
func (p *ProxyQueryFollowerHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, user_info2.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("准备数据出错")
		}
		return
	}
	p.SendOk("成功")
}

// parseNum 解析用户ID
func (p *ProxyQueryFollowerHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

// prepareData 准备关注者列表数据
func (p *ProxyQueryFollowerHandler) prepareData() error {
	list, err := user_info2.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

// SendError 发送错误响应
func (p *ProxyQueryFollowerHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

// SendOk 发送成功响应
func (p *ProxyQueryFollowerHandler) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
		FollowerList: p.FollowerList,
	})
}
