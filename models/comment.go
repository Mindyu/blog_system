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
	BlogID          sql.NullInt64  `gorm:"column:blog_id" json:"blog_id"`
	BlogTitle       sql.NullString `gorm:"column:blog_title" json:"blog_title"`
	CommentContent  sql.NullString `gorm:"column:comment_content" json:"comment_content"`
	CommentUsername sql.NullString `gorm:"column:comment_username" json:"comment_username"`
	Status          sql.NullInt64  `gorm:"column:status" json:"status"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comment"
}
