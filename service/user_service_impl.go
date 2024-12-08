package service

import (
	"errors"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/repository"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	MerchantRepository repository.MerchantRepository
}

func NewUserService(customerRepository repository.CustomerRepository, merchantRepository repository.MerchantRepository) UserService {
	return &UserServiceImpl{
		CustomerRepository: customerRepository,
		MerchantRepository: merchantRepository,
	}
}

func (service UserServiceImpl) Login(userLoginRequest *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	var accessToken string
	var err error = nil

	email := userLoginRequest.Username
	encryptedPassword := userLoginRequest.EncryptedPassword
	customerModel := service.CustomerRepository.FindByEmail(email)

	if nil != customerModel && customerModel.Password == encryptedPassword {
		accessToken, err = helper.GenerateJWT(email, "customer")
	} else {
		merchantModel := service.MerchantRepository.FindByEmail(email)
		if nil != merchantModel && merchantModel.Password == encryptedPassword {
			accessToken, err = helper.GenerateJWT(email, "merchant")
		}
	}

	if nil == err && len(accessToken) > 0 {
		return &response.UserLoginResponse{
			AccessToken: accessToken,
			ExpiredIn:   strconv.FormatInt(time.Now().Add(24*time.Hour).Unix(), 10),
			TokenType:   "Bearer",
		}, nil
	} else {
		return nil, errors.New("Terjadi Kesalahan saat login")
	}
}
