package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/models/video"
	video2 "tik-tok-server/app/service/video"
	"time"
)

func FeedVideo(c *gin.Context) {
	//传入时间的格式化处理
	//test
	latestTime1 := c.Query("latest_time")
	var latestTime2 time.Time

	if latestTime1 == "" {
		latestTime2 = time.Now()
	} else {
		intTime, err := strconv.ParseInt(latestTime1, 10, 64)
		if err != nil {
			latestTime2 = time.Unix(0, intTime*1e6)
		}
	}

	_, ok := c.Get("token")
	if !ok {
		//response.Fail(c, -1, "token不存在")
	}
	//var userId int64
	//userId = 0
	userId, ok := c.Get("id")
	var user_id int64
	if !ok {
		userId = 0
		user_id = int64(userId.(int))
	} else {
		user_id2, err := strconv.ParseInt(userId.(string), 10, 64)
		user_id = user_id2
		if err != nil {
			response.Fail(c, -1, "数据转换失败")
			return
		}
	}

	videoList, nextTime, err := video2.QueryFeedVideoInCondition(user_id, latestTime2)

	if err == nil {
		//data := video2.FeedVideoResponse{
		//	VideoList: videoList,
		//	NextTime:  nextTime,
		//}
		//response.Success(c, data)
		ResponseFeedOk(c, videoList, nextTime, "视频流返回成功")
		return
	}

}
func ResponseFeedOk(c *gin.Context, videos []*video.Video, nexttime int64, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
		"next_time":   nexttime,
		"video_list":  videos,
	})
}
