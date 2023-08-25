package video

import (
	"errors"
	"strconv"
	video2 "tik-tok-server/app/models/video"
	"tik-tok-server/app/service/basic/user"
)

type QueryVideoFlow struct {
	userId    int64
	videoList []*video2.Video
}

func QueryVideoById(userid int64) (*[]*video2.Video, error) {
	return NewQueryVideoFlow(userid).Do()
}

func NewQueryVideoFlow(userid int64) *QueryVideoFlow {
	return &QueryVideoFlow{
		userId: userid,
	}
}

// 真正操作的地方
func (q *QueryVideoFlow) Do() (*[]*video2.Video, error) {
	//校验数据
	if err := q.checkParams(); err != nil {
		return nil, err
	}
	//准备数据

	//打包数据
	if err := q.packData(); err != nil {
		return nil, err
	}
	return &q.videoList, nil
}

func (q *QueryVideoFlow) checkParams() error {
	if q.userId < 0 {
		return errors.New("userId错误")
	}
	return nil
}

func (q *QueryVideoFlow) prepareData() error {
	//
	return nil
}

func (q *QueryVideoFlow) packData() error {

	err := video2.NewVideoDao().QueryVideoList(&q.videoList, q.userId)

	//本来这里需要用userid查询这个作者信息，没有拉取代码,这里后面还需要将作者的信息填充进去
	err, author := user.UserService.GetUserInfo(strconv.FormatInt(q.userId, 10))
	if err != nil {
		return err
	}

	for index, _ := range q.videoList {
		q.videoList[index].Author = &author
	}

	if err != nil {
		return err
	}
	return nil
}
