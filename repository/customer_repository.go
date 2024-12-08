package repository

import "katalisRobo/component-store/model"

type CustomerRepository interface {
	Save(customer *model.Customer)
	FindByEmail(email string) *model.Customer
	FindAll() []*model.Customer
	Update(customer *model.Customer)
	Delete(email string)
}
