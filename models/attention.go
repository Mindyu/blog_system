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

type Attention struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	FocusUserID   int       `gorm:"column:focus_user_id" json:"focus_user_id"`
	FocusedUserID int       `gorm:"column:focused_user_id" json:"focused_user_id"`
	Status        int       `gorm:"column:status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (a *Attention) TableName() string {
	return "attention"
}
