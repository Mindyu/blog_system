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

type CommentReply struct {
	ID           int       `gorm:"column:id;primary_key" json:"id"`
	CommentID    int       `gorm:"column:comment_id" json:"comment_id"`
	ReplyID      int       `gorm:"column:reply_id"json:"reply_id"`
	ReplyType    int       `gorm:"column:reply_type"json:"reply_type"`
	ReplyContent string    `gorm:"column:reply_content" json:"reply_content"`
	FromUsername string    `gorm:"column:from_username" json:"from_username"`
	ToUsername   string    `gorm:"column:to_username" json:"to_username"`
	Status       int       `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (c *CommentReply) TableName() string {
	return "comment_reply"
}

// 实现它的json序列化方法
func (this CommentReply) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom CommentReply
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
