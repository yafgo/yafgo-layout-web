// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "t_user"

// User mapped from table <t_user>
type User struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string         `gorm:"column:username;not null;comment:用户名" json:"username"`
	Password  string         `gorm:"column:password;not null;comment:密码" json:"password"`
	Phone     string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`
	Nickname  string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`
	Avatar    string         `gorm:"column:avatar;not null;comment:头像" json:"avatar"`
	Name      string         `gorm:"column:name;not null;comment:用户姓名" json:"name"`
	Gender    string         `gorm:"column:gender;not null;comment:男｜女" json:"gender"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}