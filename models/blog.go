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

type Blog struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	BlogTitle   string    `gorm:"column:blog_title" json:"blog_title"`
	BlogContent string    `gorm:"column:blog_content" json:"blog_content"`
	Keywords    string    `gorm:"column:keywords" json:"keywords"`
	Author      string    `gorm:"column:author" json:"author"`
	TypeID      int       `gorm:"column:type_id" json:"type_id"`
	Personal    int       `gorm:"column:personal" json:"personal"`
	ReadCount   int       `gorm:"column:read_count" json:"read_count"`
	ReplyCount  int       `gorm:"column:reply_count" json:"reply_count"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (b *Blog) TableName() string {
	return "blog"
}

// 实现它的json序列化方法
func (this Blog) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Blog
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
