package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Comment struct {
	ID              int             `gorm:"column:id;primary_key" json:"id"`
	BlogID          int             `gorm:"column:blog_id" json:"blog_id"`
	BlogTitle       string          `gorm:"column:blog_title" json:"blog_title"`
	CommentContent  string          `gorm:"column:comment_content" json:"comment_content"`
	CommentUsername string          `gorm:"column:comment_username" json:"comment_username"`
	Status          int             `gorm:"column:status" json:"status"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at" json:"updated_at"`
	CommentReply    []*CommentReply `json:"comment_reply"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comment"
}


// 实现它的json序列化方法
func (this Comment) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Comment
	// 定义一个新的结构体
	tmp := struct {
		AliasCom
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		AliasCom:  (AliasCom)(this),
		CreatedAt: this.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: this.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmp)
}
