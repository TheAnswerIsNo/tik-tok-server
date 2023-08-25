package video

import (
	"strconv"
	"tik-tok-server/app/models/video"
	"tik-tok-server/app/service/basic/user"
	"time"
)

const MaxCount = 30

type FeedVideoResponse struct {
	VideoList []*video.Video `json:"video_list"`
	//因为response中nexttime是integer
	NextTime int64 `json:"next_time"`
}

type FeedVideoResponseFlow struct {
	userId     int64
	latestTime time.Time

	videolist []*video.Video
	nexttime  int64

	feedvideoresponse *FeedVideoResponse
}

func QueryFeedVideoInCondition(userid int64, latesttime time.Time) ([]*video.Video, int64, error) {
	return NewFeedVideoResponseFlow(userid, latesttime).Do()
}

func NewFeedVideoResponseFlow(userid int64, latesttime time.Time) *FeedVideoResponseFlow {
	return &FeedVideoResponseFlow{
		userId:     userid,
		latestTime: latesttime,
	}
}

func (f *FeedVideoResponseFlow) Do() ([]*video.Video, int64, error) {
	//校验数据
	f.checkParmas()

	//准备数据
	f.prepareData()

	//打包数据
	f.packData()

	return f.videolist, f.nexttime, nil
}
func (f *FeedVideoResponseFlow) checkParmas() error {
	if f.latestTime.IsZero() {
		f.latestTime = time.Now()
	}
	return nil
}

func (f *FeedVideoResponseFlow) prepareData() error {
	err := video.NewVideoDao().QueryFeedVideoList(&f.videolist, f.latestTime, MaxCount)
	if err != nil {
		return err
	}
	//填充作者信息
	for k, _ := range f.videolist {
		err, author := user.UserService.GetUserInfo(strconv.FormatInt(f.videolist[k].UserId, 10))
		if err != nil {
			return err
		}
		f.videolist[k].Author = &author
	}

	if f.userId > 0 {
		//如果当前已经登录，需要填充是否点赞，是否关注的数据
		//填充点赞数据

		//填充关注

	}

	size := len(f.videolist)
	if size >= 1 {
		f.latestTime = f.videolist[size-1].CreateTime
	}

	//返回一个ms级别的时间戳
	f.nexttime = f.latestTime.UnixNano() / 1e6

	return nil
}

func (f *FeedVideoResponseFlow) packData() error {
	f.feedvideoresponse = &FeedVideoResponse{
		VideoList: f.videolist,
		NextTime:  f.nexttime,
	}
	return nil
}
