package model

type Commodity struct {
	ID    uint64  `gorm:"column:id"`
	Name  string  `gorm:"column:name"`
	Price float32 `gorm:"column:price"`
	Count int     `gorm:"column:price"`
}

func (Commodity) TableName() string {
	return "commodities"
}
