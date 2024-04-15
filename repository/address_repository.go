package repository

import "katalisRobo/component-store/model"

type AddressRepository interface {
	Save(address *model.Address)
	FindByEmail(email string) []*model.Address
	FindByIdAndEmail(addressId string, email string) *model.Address
	Update(address *model.Address)
	Delete(addressId string, email string)
}
