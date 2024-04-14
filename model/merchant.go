package model

import "time"

type Merchant struct {
	ID              string    `gorm:"primary_key;column:id"`
	Name            string    `gorm:"column:name"`
	Email           string    `gorm:"column:email"`
	Password        string    `gorm:"column:password"`
	PhoneNumber     string    `gorm:"column:phone_number"`
	AccountNumber   string    `gorm:"column:account_number"`
	IsLoginDisabled bool      `gorm:"column:is_login_disabled"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (c *Merchant) TableName() string {
	return "merchants"
}
