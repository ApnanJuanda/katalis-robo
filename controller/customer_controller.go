package controller

import (
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/service"
	"net/http"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) *CustomerController {
	return &CustomerController{
		CustomerService: customerService,
	}
}

func (controller CustomerController) Create(ctx *gin.Context) {
	customerRequest := request.CustomerCreateRequest{}
	if err := ctx.ShouldBindJSON(&customerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	controller.CustomerService.Create(&customerRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Selamat!, akun Anda berhasil didaftarkan",
	})
}

func (controller CustomerController) FindAll(ctx *gin.Context) {
	customerResponses := controller.CustomerService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"data": customerResponses,
	})
}

func (controller CustomerController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	customerResponse := controller.CustomerService.FindByEmail(email)
	if nil == customerResponse {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "Mohon maaf, customer tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})
}

func (controller CustomerController) Update(ctx *gin.Context) {
	customerUpdateRequest := request.CustomerUpdateRequest{}
	if err := ctx.ShouldBindJSON(&customerUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	email := ctx.Param("email")
	customerUpdateResponse := controller.CustomerService.Update(email, &customerUpdateRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"data": customerUpdateResponse,
	})
}

func (controller CustomerController) Delete(ctx *gin.Context) {
	email := ctx.Param("email")
	controller.CustomerService.Delete(email)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Akun Anda berhasil dihapus",
	})
}
