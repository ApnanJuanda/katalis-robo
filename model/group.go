package model

type Group struct {
	ID         string  `gorm:"primary_key;column:id"`
	ProductId  string  `gorm:"column:product_id"`
	CategoryId string  `gorm:"column:category_id"`
	Product    Product `gorm:"foreignKey:product_id;references:id"`
	Category   Product `gorm:"foreignKey:category_id;references:id"`
}
