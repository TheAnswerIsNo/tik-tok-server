package comment

import common "tik-tok-server/app/models"

// Comment 评论实体
type Comment struct {
	common.ID
	common.Timestamp
	common.PseudoDeletion
}
