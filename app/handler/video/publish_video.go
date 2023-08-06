package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"tik-tok-server/app/common/constants"
	"tik-tok-server/app/common/response"
	"tik-tok-server/app/service/video"
	"tik-tok-server/utils"
)

func PublishVideoHandler(c *gin.Context) {
	//获取参数
	//token参数(登录注册完成后开放)
	//token, err := c.Get("token")

	//if err != nil {
	//	response.Fail(c, -1, "token不存在")
	//}

	//title
	title := c.PostForm("title")
	if title == "" || len(title) > 30 {
		response.Fail(c, -1, "视频标题不能为空或长度超过30")
		return
	}

	//file参数
	form, err := c.MultipartForm()
	if err != nil {
		response.Fail(c, -1, "文件上传失败")
		return
	}

	files := form.File["data"]
	for _, file := range files {
		//判断文件名字后缀是否正确
		suffix := filepath.Ext(file.Filename)
		if suffix != ".mp4" {
			response.Fail(c, -1, "当前文件格式不支持")
			continue
		}

		//测试完成后开放
		//userId, err := c.Get("id")
		//if err != nil {
		//	response.Fail(c, -1, "用户id获取失败")
		//	return
		//}

		//测试数据
		userId := "1"
		randNum := utils.RandString(6)
		filename := fmt.Sprintf("%s_%s%s", userId, randNum, suffix)
		videoFilePath := constants.FILEPREFIX + filename

		if err := c.SaveUploadedFile(file, videoFilePath); err != nil {
			response.Fail(c, -1, "文件上传失败")
			continue
		}
		//保存截图
		imageFile := fmt.Sprintf("%s_%s%s", userId, randNum, ".jpg")
		imageFilePath := constants.IMAGEPREFIX + imageFile
		if err := utils.RunCmd(imageFilePath, videoFilePath); err != nil {
			response.Fail(c, -1, "图片保存失败")
			continue
		}
		//视频信息持久化
		if err := video.PostVideo(userId, title, videoFilePath, imageFilePath); err == nil {
			ResponseOK(c, "视频上传成功")
		}
	}

}
func ResponseOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"StatusCode": 1,
		"StatusMsg": msg})
}
