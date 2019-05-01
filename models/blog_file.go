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

type BlogFile struct {
	FileID    int            `gorm:"column:file_id;primary_key" json:"file_id"`
	BlogID    sql.NullInt64  `gorm:"column:blog_id" json:"blog_id"`
	ImgUrls   sql.NullString `gorm:"column:img_urls" json:"img_urls"`
	AttachURL sql.NullString `gorm:"column:attach_url" json:"attach_url"`
	Status    int            `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (b *BlogFile) TableName() string {
	return "blog_file"
}
