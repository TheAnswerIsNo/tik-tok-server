package comment

//评论模块控制层
import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/models/comment/request"
	"tik-tok-server/app/service/interact/comment"
)

type Handler struct {
}

var (
	commentHandler Handler
)

func QueryCommentHandler(c *gin.Context) {
	err := commentHandler.getCommentList(c)
	if err != nil {
		response.BusinessFail(c, "查询失败")
		return
	}
}

func ActionCommentHandler(c *gin.Context) {
	err := commentHandler.commentAction(c)
	if err != nil {
		response.BusinessFail(c, "评论操作失败")
		return
	}
}

// getCommentList 获取评论列表
func (*Handler) getCommentList(context *gin.Context) error {

	rawVideoId := context.Query("videoId")

	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}

	list, err := comment.PackageCommentList(0, videoId)
	if err != nil {
		return err
	}
	response.Success(context, list)

	return nil
}

// commentAction 评论互动
func (*Handler) commentAction(context *gin.Context) error {
	var requestParam request.ListRequest

	if err := context.BindJSON(&requestParam); err != nil {
		return err
	}

	if err := comment.Operation(&requestParam); err != nil {
		return err
	}
	response.SuccessNoData(context, "评论操作成功")

	return nil
}
