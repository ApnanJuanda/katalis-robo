package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository ProductRepositoryImpl) Save(product *model.Product) {
	err := repository.DB.Create(product).Error
	helper.PanicIfError(err)
}

func (repository ProductRepositoryImpl) FindById(id string) *model.Product {
	var product model.Product
	err := repository.DB.Model(&product).Joins("Merchant").First(&product, "products.id = ?", id).Error
	helper.PanicIfError(err)
	return &product
}

func (repository ProductRepositoryImpl) FindByMerchantId(merchantId string) []*model.Product {
	var productsByMerchant []*model.Product
	err := repository.DB.Model(&productsByMerchant).Joins("Merchant").Order("updated_at, name").
		Find(&productsByMerchant, "products.merchant_id = ?", merchantId).Error
	helper.PanicIfError(err)

	return productsByMerchant
}

func (repository ProductRepositoryImpl) Update(product *model.Product) {
	err := repository.DB.Save(product).Error
	helper.PanicIfError(err)
}

func (repository ProductRepositoryImpl) Delete(id string) {
	tx := repository.DB.Where("id = ?", id).Delete(&model.Product{})
	helper.PanicIfError(tx.Error)
}
