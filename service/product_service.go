package service

import (
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type ProductService interface {
	Create(productCreateRequest *request.ProductCreateRequest, authUser *current.AuthUser) *response.ProductResponse
	FindById(id string) *response.ProductResponse
	FindByMerchantId(merchantId string) []*response.ProductResponse
	Update(id string, productUpdateRequest *request.ProductUpdateRequest) *response.ProductResponse
	Delete(id string)
}
