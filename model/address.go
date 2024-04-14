package model

import "time"

type Address struct {
	ID         string    `gorm:"primary_key;column:id"`
	CustomerId string    `gorm:"column:customer_id"`
	Detail     string    `gorm:"column:detail"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (c *Address) TableName() string {
	return "addresses"
}
