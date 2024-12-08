package helper

import (
	"crypto/rand"
	"errors"
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
		ID:              GenerateId(4, "customer"),
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
func GenerateId(maxDigits uint32, typeAccount string) string {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	if typeAccount == "customer" {
		return fmt.Sprintf("AC%0*d", maxDigits, bi)
	} else if typeAccount == "merchant" {
		return fmt.Sprintf("SPL%0*d", maxDigits, bi)
	} else if typeAccount == "product" {
		return fmt.Sprintf("P%0*d", maxDigits, bi)
	} else if typeAccount == "group" {
		return fmt.Sprintf("G%0*d", maxDigits, bi)
	} else if typeAccount == "address" {
		return fmt.Sprintf("ADD%0*d", maxDigits, bi)
	} else if typeAccount == "category" {
		return fmt.Sprintf("C%0*d", maxDigits, bi)
	}
	return ""
}

func MerchantPopulator(merchantCreateRequest *request.MerchantCreateRequest) (*model.Merchant, error) {
	_, err := GetAESDecrypted(merchantCreateRequest.Password)
	PanicIfError(err)

	return &model.Merchant{
		ID:              GenerateId(4, "merchant"),
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

func ProductPopulator(productCreateRequest *request.ProductCreateRequest, merchantId string) (*model.Product, error) {
	return &model.Product{
		ID:         GenerateId(4, "product"),
		Name:       productCreateRequest.Name,
		Price:      productCreateRequest.Price,
		Stock:      productCreateRequest.Stock,
		ImageUrl:   productCreateRequest.ImageUrl,
		MerchantId: merchantId,
	}, nil
}

func ProductResponsePopulator(productModel *model.Product) (*response.ProductResponse, error) {
	merchantResponse, err := MerchantResponsePopulator(&productModel.Merchant)
	if err == nil {
		return &response.ProductResponse{
			Name:     productModel.Name,
			Merchant: merchantResponse,
			Price:    productModel.Price,
			Stock:    productModel.Stock,
			ImageUrl: productModel.ImageUrl,
		}, nil
	} else {
		return nil, errors.New("Terjadi kesalahan")
	}
}

func ProductUpdatePopulator(productModel *model.Product, productUpdateRequest *request.ProductUpdateRequest) (*model.Product, error) {
	productModel.Name = productUpdateRequest.Name
	productModel.Price = productUpdateRequest.Price
	productModel.Stock = productUpdateRequest.Stock
	productModel.ImageUrl = productUpdateRequest.ImageUrl

	return productModel, nil
}

func GroupPopulator(productId string, categoryId string) (*model.Group, error) {
	return &model.Group{
		ID:         GenerateId(4, "group"),
		ProductId:  productId,
		CategoryId: categoryId,
	}, nil
}

func AddressPopulator(addressCreateRequest *request.AddressCreateUpdateRequest, customerId string) (*model.Address, error) {
	return &model.Address{
		ID:         GenerateId(4, "product"),
		CustomerId: customerId,
		Detail:     addressCreateRequest.Detail,
	}, nil
}

func AddressResponsePopulator(addressModel *model.Address) (*response.AddressResponse, error) {
	customerResponse, err := CustomerResponsePopulator(&addressModel.Customer)
	if err == nil {
		return &response.AddressResponse{
			Customer: customerResponse,
			Address:  addressModel.Detail,
		}, nil
	} else {
		return nil, errors.New("Terjadi kesalahan")
	}
}

func AddressUpdatePopulator(addressModel *model.Address, addressUpdateRequest *request.AddressCreateUpdateRequest) (*model.Address, error) {
	addressModel.Detail = addressUpdateRequest.Detail
	return addressModel, nil
}

func CategoryPopulator(categoryRequest *request.CategoryCreateUpdateRequest) (*model.Category, error) {
	return &model.Category{
		ID:   GenerateId(4, "category"),
		Name: categoryRequest.Name,
	}, nil
}

func CategoryUpdatePopulator(categoryModel *model.Category, categoryUpdateRequest *request.CategoryCreateUpdateRequest) (*model.Category, error) {
	categoryModel.Name = categoryUpdateRequest.Name
	return categoryModel, nil
}

func CategoryCreateUpdateResponse(categoryModel *model.Category) (*response.CategoryCreateUpdateResponse, error) {
	return &response.CategoryCreateUpdateResponse{
		ID:   categoryModel.ID,
		Name: categoryModel.Name,
	}, nil
}
