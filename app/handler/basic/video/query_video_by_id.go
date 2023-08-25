package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/models/video"
	video2 "tik-tok-server/app/service/basic/video"
)

type ResultVideo struct {
	StatusCode int64           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	VideoList  *[]*video.Video `json:"video_list"`
}

func QueryVideoById(c *gin.Context) {
	//获取参数
	_, ok := c.Get("token")
	if !ok {
		response.Fail(c, -1, "token不存在，未登录")
		return
	}
	id, ok := c.Get("id")
	if !ok {
		response.Fail(c, -1, "userid不存在，未登录")
		return
	}
	userId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		response.Fail(c, -1, "参数转换类型失败")
		return
	}

	videoList, err := video2.QueryVideoById(userId)

	if err == nil {
		resultvideo := &ResultVideo{
			StatusCode: 0,
			StatusMsg:  "查询成功",
			VideoList:  videoList,
		}
		ResponseOk(c, resultvideo)
	}

}
func ResponseOk(c *gin.Context, resultVideo *ResultVideo) {
	c.JSON(http.StatusOK, resultVideo)
}
