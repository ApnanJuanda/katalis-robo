package service

import (
	"github.com/go-playground/validator/v10"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
	"katalisRobo/component-store/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		Validate:           validate,
	}
}

func (service CustomerServiceImpl) Create(customerCreateRequest *request.CustomerCreateRequest) {
	err := service.Validate.Struct(customerCreateRequest)
	helper.PanicIfError(err)
	customerModel, err := helper.CustomerPopulator(customerCreateRequest)
	helper.PanicIfError(err)

	service.CustomerRepository.Save(customerModel)
}

func (service CustomerServiceImpl) FindByEmail(email string) *response.CustomerResponse {
	customerModel := service.CustomerRepository.FindByEmail(email)
	if nil != customerModel {
		customerResponse, err := helper.CustomerResponsePopulator(customerModel)
		helper.PanicIfError(err)
		return customerResponse
	}
	return nil
}

func (service CustomerServiceImpl) FindAll() []*response.CustomerResponse {
	customerModels := service.CustomerRepository.FindAll()
	var customerResponses []*response.CustomerResponse
	for _, customerModel := range customerModels {
		customerResponse, err := helper.CustomerResponsePopulator(customerModel)
		if err == nil {
			customerResponses = append(customerResponses, customerResponse)
		}
	}
	return customerResponses
}

func (service CustomerServiceImpl) Update(email string, customerUpdateRequest *request.CustomerUpdateRequest) *response.CustomerResponse {
	err := service.Validate.Struct(customerUpdateRequest)
	helper.PanicIfError(err)

	var customerUpdatedModel *model.Customer
	customerModel := service.CustomerRepository.FindByEmail(email)
	if customerModel != nil {
		customerUpdatedModel, err = helper.CustomerUpdatePopulator(customerModel, customerUpdateRequest)
		helper.PanicIfError(err)
		service.CustomerRepository.Update(customerUpdatedModel)
	}

	customerUpdatedResponse, err := helper.CustomerResponsePopulator(customerUpdatedModel)
	helper.PanicIfError(err)
	return customerUpdatedResponse
}

func (service CustomerServiceImpl) Delete(email string) {
	service.CustomerRepository.Delete(email)
}
