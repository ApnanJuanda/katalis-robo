package service

import (
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type CustomerService interface {
	Create(customerCreateRequest *request.CustomerCreateRequest)
	FindByEmail(email string) *response.CustomerResponse
	FindAll() []*response.CustomerResponse
	Update(email string, customerUpdateRequest *request.CustomerUpdateRequest) *response.CustomerResponse
	Delete(email string)
}
