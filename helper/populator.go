package helper

import (
	"crypto/rand"
	"fmt"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/model"
	"math"
	"math/big"
)

/*
this method to populate data from customerCreateRequest to customerModel
*/
func CustomerPopulator(customerRegister *request.CustomerCreateRequest) (*model.Customer, error) {
	_, err := GetAESDecrypted(customerRegister.Password)
	PanicIfError(err)

	return &model.Customer{
		ID:              GenerateCustomerId(4, "customer"),
		Name:            customerRegister.Name,
		Email:           customerRegister.Email,
		Password:        customerRegister.Password,
		PhoneNumber:     customerRegister.PhoneNumber,
		IsLoginDisabled: false,
	}, nil
}

/*
this method to populate data from customerModel to customerResponse
*/
func CustomerResponsePopulator(customerModel *model.Customer) (*response.CustomerResponse, error) {

	return &response.CustomerResponse{
		Name:        customerModel.Name,
		Email:       customerModel.Email,
		PhoneNumber: customerModel.PhoneNumber,
	}, nil
}

func CustomerUpdatePopulator(customerModel *model.Customer, customerUpdateRequest *request.CustomerUpdateRequest) (*model.Customer, error) {
	customerModel.Name = customerUpdateRequest.Name
	customerModel.PhoneNumber = customerUpdateRequest.PhoneNumber
	return customerModel, nil
}

/*
this method to generate customer id
*/
func GenerateCustomerId(maxDigits uint32, typeAccount string) string {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	if typeAccount == "customer" {
		return fmt.Sprintf("AC%0*d", maxDigits, bi)
	} else {
		return fmt.Sprintf("SPL%0*d", maxDigits, bi)
	}
}

func MerchantPopulator(merchantCreateRequest *request.MerchantCreateRequest) (*model.Merchant, error) {
	_, err := GetAESEncrypted(merchantCreateRequest.Password)
	PanicIfError(err)

	return &model.Merchant{
		ID:              GenerateCustomerId(4, "supplier"),
		Name:            merchantCreateRequest.Name,
		Email:           merchantCreateRequest.Email,
		Password:        merchantCreateRequest.Password,
		PhoneNumber:     merchantCreateRequest.PhoneNumber,
		AccountNumber:   merchantCreateRequest.AccountNumber,
		IsLoginDisabled: false,
	}, nil
}

func MerchantResponsePopulator(merchantModel *model.Merchant) (*response.MerchantResponse, error) {

	return &response.MerchantResponse{
		Name:        merchantModel.Name,
		Email:       merchantModel.Email,
		PhoneNumber: merchantModel.PhoneNumber,
	}, nil
}

func MerchantUpdatePopulator(merchantModel *model.Merchant, merchantUpdateRequest *request.MerchantUpdateRequest) (*model.Merchant, error) {
	merchantModel.Name = merchantUpdateRequest.Name
	merchantModel.PhoneNumber = merchantUpdateRequest.PhoneNumber
	merchantModel.AccountNumber = merchantModel.AccountNumber
	return merchantModel, nil
}
