package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) Login(ctx *gin.Context) {
	userLoginRequest := request.UserLoginRequest{}
	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Mohon, masukan Username atau Password yang benar",
		})
	}
	userLoginResponse, err := controller.UserService.Login(&userLoginRequest)
	data, err := helper.DecryptJWT(userLoginResponse.AccessToken)
	fmt.Println(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Username atau Password tidak valid",
		})
		return
	}
	ctx.JSON(http.StatusOK, userLoginResponse)
}
