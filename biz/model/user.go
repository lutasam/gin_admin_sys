package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"column:id"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	NickName  string         `gorm:"column:nickname"`
	Avatar    string         `gorm:"column:avatar"`
	Sign      string         `gorm:"column:sign"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
