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

type Friend struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Username1 string    `gorm:"column:username_1" json:"username_1"`
	Username2 string    `gorm:"column:username_2" json:"username_2"`
	Status    int       `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (f *Friend) TableName() string {
	return "friend"
}


// 实现它的json序列化方法
func (this Friend) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Friend
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