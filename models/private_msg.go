package models

import (
	"database/sql"
	"encoding/json"
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

// 实现它的json序列化方法
func (this PrivateMsg) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom PrivateMsg
	// 定义一个新的结构体
	tmp := struct {
		AliasCom
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		AliasCom:  (AliasCom)(this),
		CreatedAt: this.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: this.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmp)
}
