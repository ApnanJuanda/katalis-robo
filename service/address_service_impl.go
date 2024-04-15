package service

import (
	"github.com/go-playground/validator/v10"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
	"katalisRobo/component-store/repository"
)

type AddressServiceImpl struct {
	AddressRepository  repository.AddressRepository
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewAddressService(addressRepository repository.AddressRepository, customerRepository repository.CustomerRepository,
	validate *validator.Validate) AddressService {
	return &AddressServiceImpl{
		AddressRepository:  addressRepository,
		CustomerRepository: customerRepository,
		Validate:           validate,
	}
}

func (service AddressServiceImpl) Create(email string, addressCreateRequest *request.AddressCreateUpdateRequest) *response.AddressResponse {
	err := service.Validate.Struct(addressCreateRequest)
	helper.PanicIfError(err)
	customerModel := service.CustomerRepository.FindByEmail(email)
	if customerModel != nil {
		addressModel, err := helper.AddressPopulator(addressCreateRequest, customerModel.ID)
		if err == nil {
			service.AddressRepository.Save(addressModel)
			addressResponse := service.FindByIdAndEmail(addressModel.ID, email)
			return addressResponse
		}
	}
	return nil
}

func (service AddressServiceImpl) FindByEmail(email string) []*response.AddressResponse {
	addressesByEmail := service.AddressRepository.FindByEmail(email)
	var addressesByEmailResponse []*response.AddressResponse
	for _, addressByEmail := range addressesByEmail {
		addressResponse, err := helper.AddressResponsePopulator(addressByEmail)
		if err == nil {
			addressesByEmailResponse = append(addressesByEmailResponse, addressResponse)
		}
	}
	return addressesByEmailResponse
}

func (service AddressServiceImpl) FindByIdAndEmail(addressId string, email string) *response.AddressResponse {
	var addressResponse *response.AddressResponse
	addressResponse = nil
	addressModel := service.AddressRepository.FindByIdAndEmail(addressId, email)
	if addressModel != nil {
		addressResponse, _ = helper.AddressResponsePopulator(addressModel)
	}
	return addressResponse
}

func (service AddressServiceImpl) Update(addressId string, email string, addressUpdateRequest *request.AddressCreateUpdateRequest) *response.AddressResponse {
	err := service.Validate.Struct(addressUpdateRequest)
	helper.PanicIfError(err)

	var addressUpdatedModel *model.Address
	addressModel := service.AddressRepository.FindByIdAndEmail(addressId, email)
	if addressModel != nil {
		addressUpdatedModel, err = helper.AddressUpdatePopulator(addressModel, addressUpdateRequest)
		if err == nil {
			service.AddressRepository.Update(addressUpdatedModel)
		}
	}
	addressUpdatedResponse, err := helper.AddressResponsePopulator(addressUpdatedModel)
	helper.PanicIfError(err)
	return addressUpdatedResponse
}

func (service AddressServiceImpl) Delete(addressId string, email string) {
	service.AddressRepository.Delete(addressId, email)
}
