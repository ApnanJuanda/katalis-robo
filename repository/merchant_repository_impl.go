package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type MerchantRepositoryImpl struct {
	DB *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &MerchantRepositoryImpl{
		DB: db,
	}
}

func (repository MerchantRepositoryImpl) Save(merchant *model.Merchant) {
	err := repository.DB.Create(merchant).Error
	helper.PanicIfError(err)
}

func (repository MerchantRepositoryImpl) FindByEmail(email string) *model.Merchant {
	var merchant model.Merchant
	err := repository.DB.First(&merchant, "email = ?", email).Error
	if err != nil {
		return nil
	}
	return &merchant
}

func (repository MerchantRepositoryImpl) FindAll() []*model.Merchant {
	var merchants []*model.Merchant
	err := repository.DB.Order("updated_at desc, email").Find(&merchants).Error
	helper.PanicIfError(err)

	return merchants
}

func (repository MerchantRepositoryImpl) Update(merchant *model.Merchant) {
	err := repository.DB.Save(merchant).Error
	helper.PanicIfError(err)
}

func (repository MerchantRepositoryImpl) Delete(email string) {
	tx := repository.DB.Where("email = ?", email).Delete(model.Merchant{})
	helper.PanicIfError(tx.Error)
}
