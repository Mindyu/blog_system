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

type PrivateMsg struct {
	ID         int            `gorm:"column:id;primary_key" json:"id"`
	Sender     sql.NullString `gorm:"column:sender" json:"sender"`
	Receiver   sql.NullString `gorm:"column:receiver" json:"receiver"`
	MsgContent sql.NullString `gorm:"column:msg_content" json:"msg_content"`
	Status     int            `gorm:"column:status" json:"status"`
	IsRead     sql.NullInt64  `gorm:"column:is_read" json:"is_read"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (p *PrivateMsg) TableName() string {
	return "private_msg"
}
