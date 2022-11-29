package model

import (
	"gorm.io/gorm"
	"time"
)

type Commodity struct {
	ID        uint64         `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Price     float32        `gorm:"column:price"`
	Count     int            `gorm:"column:price"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Commodity) TableName() string {
	return "commodities"
}
