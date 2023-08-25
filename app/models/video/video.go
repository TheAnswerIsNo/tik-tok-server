package video

import (
	"sync"
	"tik-tok-server/app/models/user"
	"tik-tok-server/global"
	"time"
)

//	type Video struct {
//		Id     int64 `json:"id,omitempty"`
//		UserId int64 `json:"-"`
//		//下面需要修改为真正的User
//		Author        *user.User `json:"author,omitempty" gorm:"-"`
//		PlayUrl       string     `json:"play_url,omitempty"`
//		CoverUrl      string     `json:"cover_url,omitempty"`
//		FavoriteCount int64      `json:"favorite_count"`
//		CommentCount  int64      `json:"comment_count"`
//		IsFavorite    bool       `json:"is_favorite"`
//		Title         string     `json:"title"`
//		CreateTime    time.Time  `json:"-" gorm:"autoUpdateTime"`
//	}
type Video struct {
	Id            int64      `json:"id,omitempty" gorm:"type:int;primaryKey;unique;size:20;not NULL;AUTO_INCREMENT" column:"id" comment:"视频唯一标识"`
	UserId        int64      `json:"-" gorm:"size:20;" column:"user_id" comment:"当前作者id"`
	Author        *user.User `json:"author,omitempty" gorm:"-"`
	PlayUrl       string     `json:"play_url,omitempty" gorm:"type:varchar(255);" column:"play_url" comment:"视频播放地址"`
	CoverUrl      string     `json:"cover_url,omitempty" gorm:"type:varchar(255);" column:"cover_url" comment:"视频封面地址"`
	FavoriteCount int64      `json:"favorite_count" gorm:"type:int;size=100;default:0;" column:"favorite_count" comment:"视频的点赞总数"`
	CommentCount  int64      `json:"comment_count" gorm:"type:int;size=100;default:0;" column:"comment_count" comment:"视频的评论总数"`
	IsFavorite    bool       `json:"is_favorite" gorm:"type:bool;size=1;default:0;" column:"is_favorite" comment:"是否点赞"`
	Title         string     `json:"title" gorm:"type:varchar(255);" column:"title" comment:"视频标题"`
	CreateTime    time.Time  `json:"-" gorm:"type:datetime;autoUpdateTime" column:"create_time" comment:"视频发布时间"`
}

func (Video) TabelName() string {
	return "video"
}

type VideoDao struct {
}

var (
	videodao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDao() *VideoDao {
	videoOnce.Do(func() {
		videodao = new(VideoDao)
	})
	return videodao
}

func (*VideoDao) InsertVideo(video *Video) error {
	if err := global.App.DB.Create(&video); err != nil {
		return err.Error
	}
	return nil
}

func (*VideoDao) QueryVideoList(videoList *[]*Video, userid int64) error {
	err := global.App.DB.Where("user_id=?", userid).Find(videoList).Error
	if err != nil {
		return err
	}
	return nil
}
func (*VideoDao) QueryFeedVideoList(videoList *[]*Video, latestime time.Time, maxcount int) error {
	err := global.App.DB.
		Where("create_time<?", latestime).
		Limit(maxcount).Order("create_time desc").
		Find(videoList).Error
	if err != nil {
		return err
	}
	return nil
}
