package video

import (
	"errors"
	"strconv"
	"tik-tok-server/app/models/video"
)

//type PVideo struct {
//	video *video.Video
//}

type PVideoFlow struct {
	userId        string
	title         string
	videoFilePath string
	imageFilePath string

	//pvideo *PVideo

	video *video.Video
}

func PostVideo(userId string, title string, videoFilePath string, imageFilePath string) error {
	return NewPostVideoFlow(userId, title, videoFilePath, imageFilePath).Do()
}

func NewPostVideoFlow(userId string, title string, videoFilePath string, imageFilePath string) *PVideoFlow {
	return &PVideoFlow{
		userId:        userId,
		title:         title,
		videoFilePath: videoFilePath,
		imageFilePath: imageFilePath,
	}
}

func (p *PVideoFlow) Do() error {
	//参数校验
	if err := p.checkParams(); err != nil {
		return err
	}
	//准备数据
	//打包数据
	if err := p.packData(); err != nil {
		return err
	}
	if err := video.NewVideoDao().InsertVideo(p.video); err != nil {
		return err
	}
	return nil

}
func (p *PVideoFlow) checkParams() error {
	if p.userId == "" || p.title == "" || p.videoFilePath == "" || p.imageFilePath == "" {
		return errors.New("参数不全")
	}
	return nil
}
func (p *PVideoFlow) packData() error {

	temp, err := strconv.ParseInt(p.userId, 10, 64)
	if err != nil {
		return errors.New("数据转换失败")
	}
	p.video = &video.Video{
		UserId:   temp,
		Title:    p.title,
		PlayUrl:  p.videoFilePath,
		CoverUrl: p.imageFilePath,
	}
	return nil
}
