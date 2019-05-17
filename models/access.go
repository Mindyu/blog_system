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

type Access struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	Username   string    `gorm:"column:username" json:"username"`
	Ip         string    `gorm:"column:ip" json:"ip"`
	AccessTime time.Time `gorm:"column:access_time" json:"access_time"`
}

// TableName sets the insert table name for this struct type
func (l *Access) TableName() string {
	return "access"
}

// 实现它的json序列化方法
func (this Access) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasCom Access
	// 定义一个新的结构体
	tmp := struct {
		AliasCom
		AccessTime string `json:"access_time"`
	}{
		AliasCom:   (AliasCom)(this),
		AccessTime: this.AccessTime.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmp)
}
