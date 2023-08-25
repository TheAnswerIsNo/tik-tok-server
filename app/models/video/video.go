package video

import (
	"sync"
	"tik-tok-server/app/models/user"
	"tik-tok-server/global"
	"time"
)

type Video struct {
	Id     int64 `json:"id,omitempty"`
	UserId int64 `json:"-"`
	//下面需要修改为真正的User
	Author        *user.User `json:"author,omitempty" gorm:"-"`
	PlayUrl       string     `json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count"`
	CommentCount  int64      `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
	CreateTime    time.Time  `json:"-" gorm:"autoUpdateTime"`
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
