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

type Log struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	Username  string         `gorm:"column:username" json:"username"`
	CallAPI   string         `gorm:"column:call_api" json:"call_api"`
	Params    sql.NullString `gorm:"column:params" json:"params"`
	Operation string         `gorm:"column:operation" json:"operation"`
	Level     sql.NullString `gorm:"column:level" json:"level"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (l *Log) TableName() string {
	return "log"
}
