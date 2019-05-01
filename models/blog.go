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
	ID          int            `gorm:"column:id;primary_key" json:"id"`
	BlogTitle   sql.NullString `gorm:"column:blog_title" json:"blog_title"`
	BlogContent sql.NullString `gorm:"column:blog_content" json:"blog_content"`
	Keywords    sql.NullString `gorm:"column:keywords" json:"keywords"`
	Author      sql.NullString `gorm:"column:author" json:"author"`
	TypeID      sql.NullInt64  `gorm:"column:type_id" json:"type_id"`
	Personal    sql.NullInt64  `gorm:"column:personal" json:"personal"`
	ThumbUp     sql.NullInt64  `gorm:"column:thumb_up" json:"thumb_up"`
	ThumbDown   sql.NullInt64  `gorm:"column:thumb_down" json:"thumb_down"`
	Status      int            `gorm:"column:status" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (b *Blog) TableName() string {
	return "blog"
}
