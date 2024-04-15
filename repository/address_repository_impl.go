package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type AddressRepositoryImpl struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		DB: db,
	}
}

func (repository AddressRepositoryImpl) Save(address *model.Address) {
	err := repository.DB.Create(address).Error
	helper.PanicIfError(err)
}

func (repository AddressRepositoryImpl) FindByEmail(email string) []*model.Address {
	var addressByEmail []*model.Address
	err := repository.DB.Model(&addressByEmail).Joins("Customer").Order("updated_at").
		Find(&addressByEmail, "customers.email = ?", email).Error
	if err != nil {
		return nil
	}
	return addressByEmail
}

func (repository AddressRepositoryImpl) FindByIdAndEmail(addressId string, email string) *model.Address {
	var address model.Address
	err := repository.DB.Model(&address).Joins("Customer").
		First(&address, "addresses.id = ? && customers.Email = ?", addressId, email).Error
	if err != nil {
		return nil
	}
	return &address
}

func (repository AddressRepositoryImpl) Update(address *model.Address) {
	err := repository.DB.Save(address).Error
	helper.PanicIfError(err)
}

func (repository AddressRepositoryImpl) Delete(addressId string, email string) {
	var address model.Address
	err := repository.DB.Model(&address).Joins("Customer").
		First(&address, "addresses.id = ? && customers.Email = ?", addressId, email).Error
	if err == nil {
		err = repository.DB.Delete(&address).Error
		helper.PanicIfError(err)
	}
}
