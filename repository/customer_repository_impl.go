package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		DB: db,
	}
}

func (repository CustomerRepositoryImpl) Save(customer *model.Customer) {
	err := repository.DB.Create(customer).Error
	helper.PanicIfError(err)
}

func (repository CustomerRepositoryImpl) FindByEmail(email string) *model.Customer {
	var customer model.Customer
	err := repository.DB.First(&customer, "email = ?", email).Error
	if err != nil {
		return nil
	}
	return &customer
}

func (repository CustomerRepositoryImpl) FindAll() []*model.Customer {
	var customers []*model.Customer
	err := repository.DB.Order("updated_at desc, email").Find(&customers).Error
	helper.PanicIfError(err)

	return customers
}

func (repository CustomerRepositoryImpl) Update(customer *model.Customer) {
	err := repository.DB.Save(customer).Error
	helper.PanicIfError(err)
}

func (repository CustomerRepositoryImpl) Delete(email string) {
	tx := repository.DB.Where("email = ?", email).Delete(&model.Customer{})
	helper.PanicIfError(tx.Error)
}
