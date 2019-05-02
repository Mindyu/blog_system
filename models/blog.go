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

type Blog struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	BlogTitle   string    `gorm:"column:blog_title" json:"blog_title"`
	BlogContent string    `gorm:"column:blog_content" json:"blog_content"`
	Keywords    string    `gorm:"column:keywords" json:"keywords"`
	Author      string    `gorm:"column:author" json:"author"`
	TypeID      int       `gorm:"column:type_id" json:"type_id"`
	Personal    int       `gorm:"column:personal" json:"personal"`
	ThumbUp     int       `gorm:"column:thumb_up" json:"thumb_up"`
	ThumbDown   int       `gorm:"column:thumb_down" json:"thumb_down"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (b *Blog) TableName() string {
	return "blog"
}
