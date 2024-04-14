package model

import "time"

type Product struct {
	ID         string    `gorm:"primary_key;column:id"`
	MerchantId string    `gorm:"column:merchant_id"`
	Name       string    `gorm:"column:name"`
	Price      int       `gorm:"column:price"`
	Stock      int       `gorm:"column:stock"`
	ImageUrl   string    `gorm:"column:image_url"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Merchant   Merchant  `gorm:"foreignKey:merchant_id;references:id"`
}

func (c *Product) TableName() string {
	return "products"
}
