package service

import (
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type CategoryService interface {
	Create(categoryCreateRequest *request.CategoryCreateUpdateRequest) *response.CategoryCreateUpdateResponse
	FindById(categoryId string) *response.CategoryResponse
	Update(categoryId string, categoryUpdateRequest *request.CategoryCreateUpdateRequest) *response.CategoryCreateUpdateResponse
	Delete(categoryId string)
}
