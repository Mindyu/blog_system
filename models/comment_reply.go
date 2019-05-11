package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type CommentReply struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	CommentID     int       `gorm:"column:comment_id" json:"comment_id"`
	ReplyContent  string    `gorm:"column:reply_content" json:"reply_content"`
	ReplyUsername string    `gorm:"column:reply_username" json:"reply_username"`
	Status        int       `gorm:"column:status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (c *CommentReply) TableName() string {
	return "comment_reply"
}
