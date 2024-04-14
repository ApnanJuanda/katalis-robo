package service

import (
	"github.com/go-playground/validator/v10"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
	"katalisRobo/component-store/repository"
)

type MerchantServiceImpl struct {
	MerchantRepository repository.MerchantRepository
	Validate           *validator.Validate
}

func NewMerchantService(merchantRepository repository.MerchantRepository, validate *validator.Validate) MerchantService {
	return &MerchantServiceImpl{
		MerchantRepository: merchantRepository,
		Validate:           validate,
	}
}

func (service MerchantServiceImpl) Save(merchantCreateRequest *request.MerchantCreateRequest) {
	err := service.Validate.Struct(merchantCreateRequest)
	helper.PanicIfError(err)
	merchantModel, err := helper.MerchantPopulator(merchantCreateRequest)
	helper.PanicIfError(err)
	service.MerchantRepository.Save(merchantModel)
}

func (service MerchantServiceImpl) FindByEmail(email string) *response.MerchantResponse {
	merchantModel := service.MerchantRepository.FindByEmail(email)
	if nil != merchantModel {
		merchantResponse, err := helper.MerchantResponsePopulator(merchantModel)
		helper.PanicIfError(err)
		return merchantResponse
	}
	return nil
}

func (service MerchantServiceImpl) FindAll() []*response.MerchantResponse {
	merchantModels := service.MerchantRepository.FindAll()
	var merchantResponses []*response.MerchantResponse
	for _, merchantModel := range merchantModels {
		merchantresponse, err := helper.MerchantResponsePopulator(merchantModel)
		helper.PanicIfError(err)
		merchantResponses = append(merchantResponses, merchantresponse)
	}
	return merchantResponses
}

func (service MerchantServiceImpl) Update(email string, merchantUpdateRequest *request.MerchantUpdateRequest) *response.MerchantResponse {
	err := service.Validate.Struct(merchantUpdateRequest)
	helper.PanicIfError(err)

	// find by email
	var merchantUpdatedModel *model.Merchant
	merchantModel := service.MerchantRepository.FindByEmail(email)
	if nil != merchantModel {
		merchantUpdatedModel, err = helper.MerchantUpdatePopulator(merchantModel, merchantUpdateRequest)
		helper.PanicIfError(err)
		service.MerchantRepository.Update(merchantUpdatedModel)
	}

	merchantUpdatedResponse, err := helper.MerchantResponsePopulator(merchantUpdatedModel)
	helper.PanicIfError(err)
	return merchantUpdatedResponse
}

func (service MerchantServiceImpl) Delete(email string) {
	service.MerchantRepository.Delete(email)
}
