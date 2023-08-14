package comment

//评论模块控制层
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentHandler struct{}

func (CommentHandler) GetCommentList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (CommentHandler) CommentAction(context *gin.Context) {

}
