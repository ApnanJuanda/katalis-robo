package model

import "time"

type Category struct {
	ID        string    `gorm:"primary_key;column:id"`
	ProductId string    `gorm:"column:product_id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (c *Category) TableName() string {
	return "categories"
}
