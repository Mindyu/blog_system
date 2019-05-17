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
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	Sender     string    `gorm:"column:sender" json:"sender"`
	Receiver   string    `gorm:"column:receiver" json:"receiver"`
	MsgContent string    `gorm:"column:msg_content" json:"msg_content"`
	Status     int       `gorm:"column:status" json:"status"`
	IsRead     int       `gorm:"column:is_read" json:"is_read"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (p *PrivateMsg) TableName() string {
	return "private_msg"
}
