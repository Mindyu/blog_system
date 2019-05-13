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

type Attention struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	FocusUser   string    `gorm:"column:focus_user" json:"focus_user"`
	FocusedUser string    `gorm:"column:focused_user" json:"focused_user"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (a *Attention) TableName() string {
	return "attention"
}

// 实现它的json序列化方法
func (this Attention) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Attention
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
