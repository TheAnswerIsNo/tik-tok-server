package video

import (
	"gorm.io/gorm/utils/tests"
	"sync"
	"tik-tok-server/bootstrap"
	"time"
)

type Video struct {
	Id     int64 `json:"id,omitempty"`
	UserId int64 `json:"-"`
	//下面需要修改为真正的User
	Author        *tests.User `json:"author,omitempty" gorm:"-"`
	PlayUrl       string      `json:"play_url,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount int64       `json:"favorite_count,omitempty"`
	CommentCount  int64       `json:"comment_count,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty"`
	Title         string      `json:"title,omitempty"`
	CreateTime    time.Time   `json:"-" gorm:"autoUpdateTime"`
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
	if err := bootstrap.Db.Create(&video); err != nil {
		return err.Error
	}
	return nil
}
