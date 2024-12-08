package service

import (
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type AddressService interface {
	Create(email string, addressCreateRequest *request.AddressCreateUpdateRequest) *response.AddressResponse
	FindByEmail(email string) []*response.AddressResponse
	FindByIdAndEmail(addressId string, email string) *response.AddressResponse
	Update(addressId string, email string, addressUpdateRequest *request.AddressCreateUpdateRequest) *response.AddressResponse
	Delete(addressId string, email string)
}
