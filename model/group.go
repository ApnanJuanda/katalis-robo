package model

type Group struct {
	ID         string `gorm:"primary_key;column:id"`
	ProductId  string `gorm:"column:product_id"`
	CategoryId string `gorm:"column:category_id"`
}
