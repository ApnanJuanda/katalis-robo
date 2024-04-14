package service

import (
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
)

type UserService interface {
	Login(userLoginRequest *request.UserLoginRequest) (*response.UserLoginResponse, error)
}
