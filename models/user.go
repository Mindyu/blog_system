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

type User struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Password  string    `gorm:"column:password" json:"password"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Birthday  string    `gorm:"column:birthday" json:"birthday"`
	Education string    `gorm:"column:education" json:"education"`
	Hobby     string    `gorm:"column:hobby" json:"hobby"`
	Sign      string    `gorm:"column:sign" json:"sign"`
	RoleID    int       `gorm:"column:role_id" json:"role_id"`
	Status    int       `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
