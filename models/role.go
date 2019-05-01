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

type Role struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	RoleName  sql.NullString `gorm:"column:role_name" json:"role_name"`
	Note      sql.NullString `gorm:"column:note" json:"note"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "role"
}
