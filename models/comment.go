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

type Comment struct {
	ID              int            `gorm:"column:id;primary_key" json:"id"`
	BlogID          int            `gorm:"column:blog_id" json:"blog_id"`
	BlogTitle       string         `gorm:"column:blog_title" json:"blog_title"`
	CommentContent  string         `gorm:"column:comment_content" json:"comment_content"`
	CommentUsername string         `gorm:"column:comment_username" json:"comment_username"`
	Status          int            `gorm:"column:status" json:"status"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	CommentReply    []*CommentReply `json:"comment_reply"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comment"
}
