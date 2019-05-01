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

type Auth struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	AuthName  string         `gorm:"column:auth_name" json:"auth_name"`
	Note      sql.NullString `gorm:"column:note" json:"note"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (a *Auth) TableName() string {
	return "auth"
}
