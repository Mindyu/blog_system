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

type BlogType struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	TypeName  string         `gorm:"column:type_name" json:"type_name"`
	Note      sql.NullString `gorm:"column:note" json:"note"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (b *BlogType) TableName() string {
	return "blog_type"
}
