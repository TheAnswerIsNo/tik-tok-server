package request

// ListRequest 评论操作请求参数
type ListRequest struct {
	VideoId     int64 `json:"videoId"`
	ActionType  `json:"actionType"`
	CommentText string `json:"commentText"`
	CommentId   uint   `json:"commentId"`
}

type ActionType = uint8

const (
	Issue  ActionType = iota //发布评论
	Delete                   //删除评论
)
