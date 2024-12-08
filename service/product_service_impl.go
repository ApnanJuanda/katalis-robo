package service

import (
	"github.com/go-playground/validator/v10"
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
	"katalisRobo/component-store/repository"
)

type ProductServiceImpl struct {
	ProductRepository  repository.ProductRepository
	CategoryRepository repository.CategoryRepository
	MerchantRepository repository.MerchantRepository
	GroupRepository    repository.GroupRepository
	Validate           *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository,
	merchantRepository repository.MerchantRepository, groupRepository repository.GroupRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
		MerchantRepository: merchantRepository,
		GroupRepository:    groupRepository,
		Validate:           validate,
	}
}

func (service ProductServiceImpl) Create(productCreateRequest *request.ProductCreateRequest, authUser *current.AuthUser) *response.ProductResponse {
	err := service.Validate.Struct(productCreateRequest)
	helper.PanicIfError(err)

	var productResponse *response.ProductResponse
	productResponse = nil

	category := service.CategoryRepository.FindByCategoryId(productCreateRequest.CategoryId)
	merchant := service.MerchantRepository.FindByEmail(authUser.UserEmail)
	if nil != category && nil != merchant {
		productModel, err := helper.ProductPopulator(productCreateRequest, merchant.ID)
		if err == nil {
			service.ProductRepository.Save(productModel)

			groupModel, err := helper.GroupPopulator(productModel.ID, productCreateRequest.CategoryId)
			if err == nil {
				service.GroupRepository.Save(groupModel)
				productResponse = service.FindById(productModel.ID)
			}
		}

	}
	return productResponse
}

func (service ProductServiceImpl) FindById(id string) *response.ProductResponse {
	productModel := service.ProductRepository.FindById(id)
	productResponse, err := helper.ProductResponsePopulator(productModel)
	if err != nil {
		return nil
	}
	return productResponse
}

func (service ProductServiceImpl) FindByMerchantId(merchantId string) []*response.ProductResponse {
	productsByMerchant := service.ProductRepository.FindByMerchantId(merchantId)
	var productsByMerchantResponse []*response.ProductResponse
	for _, productByMerchant := range productsByMerchant {
		productResponse, err := helper.ProductResponsePopulator(productByMerchant)
		if err == nil {
			productsByMerchantResponse = append(productsByMerchantResponse, productResponse)
		}
	}
	return productsByMerchantResponse
}

func (service ProductServiceImpl) Update(id string, productUpdateRequest *request.ProductUpdateRequest) *response.ProductResponse {
	err := service.Validate.Struct(productUpdateRequest)
	helper.PanicIfError(err)

	var productUpdatedModel *model.Product
	productModel := service.ProductRepository.FindById(id)
	if productModel != nil {
		productUpdatedModel, err = helper.ProductUpdatePopulator(productModel, productUpdateRequest)
		helper.PanicIfError(err)
		service.ProductRepository.Update(productUpdatedModel)
	}
	productUpdatedResponse, err := helper.ProductResponsePopulator(productUpdatedModel)
	helper.PanicIfError(err)
	return productUpdatedResponse
}

func (service ProductServiceImpl) Delete(id string) {
	service.ProductRepository.Delete(id)
}
