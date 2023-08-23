package comment

//评论模块逻辑处理层
import (
	"tik-tok-server/app/models/comment"
	"tik-tok-server/app/models/comment/request"
)

// PackageCommentList 封装评论列表数据
func PackageCommentList(userId, videoId int64) ([]comment.Comment, error) {

	commentList, err := comment.NewDao().QueryCommentListByVideoId(videoId)

	return commentList, err
}

// Operation 评论操作
func Operation(param *request.ListRequest) error {
	switch param.ActionType {
	case request.Issue:
		err := comment.NewDao().CreateComment(param.VideoId, param.CommentText)
		return err
	case request.Delete:
		var c comment.Comment
		c.ID.ID = param.CommentId
		err := comment.NewDao().DeleteComment(&c)
		return err
	}
	return nil
}
