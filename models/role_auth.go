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

type RoleAuth struct {
	ID     int `gorm:"column:id;primary_key" json:"id"`
	RoleID int `gorm:"column:role_id" json:"role_id"`
	AuthID int `gorm:"column:auth_id" json:"auth_id"`
}

// TableName sets the insert table name for this struct type
func (r *RoleAuth) TableName() string {
	return "role_auth"
}
