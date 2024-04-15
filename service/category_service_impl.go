package service

import (
	"github.com/go-playground/validator/v10"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
	"katalisRobo/component-store/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (service CategoryServiceImpl) Create(categoryCreateRequest *request.CategoryCreateUpdateRequest) *response.CategoryCreateUpdateResponse {
	err := service.Validate.Struct(categoryCreateRequest)
	helper.PanicIfError(err)
	categoryModel, err := helper.CategoryPopulator(categoryCreateRequest)
	helper.PanicIfError(err)

	service.CategoryRepository.Save(categoryModel)
	categoryResponse, err := helper.CategoryCreateUpdateResponse(categoryModel)
	return categoryResponse
}

func (service CategoryServiceImpl) FindById(categoryId string) *response.CategoryResponse {
	var groups []*model.Group
	groups = service.CategoryRepository.FindById(categoryId)
	if len(groups) > 0 {
		var productResponses []*response.ProductResponse
		for _, group := range groups {
			productResponse, err := helper.ProductResponsePopulator(&group.Product)
			if err == nil {
				productResponses = append(productResponses, productResponse)
			}
		}
		return &response.CategoryResponse{
			Name:     groups[0].Category.Name,
			Products: productResponses,
		}
	}
	return nil
}

func (service CategoryServiceImpl) Update(categoryId string, categoryUpdateRequest *request.CategoryCreateUpdateRequest) *response.CategoryCreateUpdateResponse {
	err := service.Validate.Struct(categoryUpdateRequest)
	helper.PanicIfError(err)

	var categoryUpdatedModel *model.Category
	categoryModel := service.CategoryRepository.FindByCategoryId(categoryId)
	if categoryModel != nil {
		categoryUpdatedModel, err = helper.CategoryUpdatePopulator(categoryModel, categoryUpdateRequest)
		helper.PanicIfError(err)
		service.CategoryRepository.Update(categoryUpdatedModel)
	}

	categoryUpdatedResponse, err := helper.CategoryCreateUpdateResponse(categoryUpdatedModel)
	helper.PanicIfError(err)
	return categoryUpdatedResponse
}

func (service CategoryServiceImpl) Delete(categoryId string) {
	service.CategoryRepository.Delete(categoryId)
}
