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

type Friend struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	Username1 sql.NullString `gorm:"column:username_1" json:"username_1"`
	Username2 sql.NullString `gorm:"column:username_2" json:"username_2"`
	Status    int            `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (f *Friend) TableName() string {
	return "friend"
}
