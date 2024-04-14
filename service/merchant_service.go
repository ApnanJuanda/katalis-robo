package service

import (
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type MerchantService interface {
	Save(merchantCreateRequest *request.MerchantCreateRequest)
	FindByEmail(email string) *response.MerchantResponse
	FindAll() []*response.MerchantResponse
	Update(email string, merchantUpdateRequest *request.MerchantUpdateRequest) *response.MerchantResponse
	Delete(email string)
}
