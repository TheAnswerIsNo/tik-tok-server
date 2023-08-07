package comment

import (
	"go.uber.org/zap"
	common "tik-tok-server/app/models"
	"tik-tok-server/global"
	"time"
)

// Comment 评论实体
type Comment struct {
	common.ID
	UserId  int64  `json:"user_id" column:"user_id"`
	VideoId int64  `json:"video_id" column:"video_id"`
	Content string `json:"content" column:"content"`
	common.Timestamp
	common.PseudoDeletion
}

// dao 评论数据操作
type dao struct {
}

var (
	commentDao dao
)

func NewDao() *dao {
	return &commentDao
}

// QueryCommentListByVideoId 根据视频id获取评论列表
func (*dao) QueryCommentListByVideoId(videoId int64) ([]Comment, error) {
	var commentList []Comment

	if err := global.App.DB.Where("video_id = ? and delete_time is null", videoId).Order("create_time desc").Find(&commentList).Error; err != nil {
		return nil, err
	}
	global.App.Log.Info("视频评论列表查询成功")
	return commentList, nil
}

// CreateComment 添加评论数据
func (*dao) CreateComment(VideoId int64, content string) error {

	comment := Comment{VideoId: VideoId, Content: content, UserId: 1, Timestamp: common.Timestamp{CreateTime: time.Now(), UpdateTime: time.Now()}}

	if result := global.App.DB.Create(&comment); result.Error != nil {
		global.App.Log.Warn(result.Error.Error(), zap.Any("err", result.Error))
		return result.Error
	}
	global.App.Log.Info("添加评论数据成功")

	return nil
}

// DeleteComment 删除评论数据
func (*dao) DeleteComment(comment *Comment) error {

	if err := global.App.DB.Model(&comment).Updates(map[string]interface{}{"delete_time": time.Now(), "update_time": time.Now()}).Error; err != nil {
		return err
	}
	global.App.Log.Info("删除评论数据成功")

	return nil
}
