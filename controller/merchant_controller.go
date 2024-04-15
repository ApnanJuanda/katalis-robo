package controller

import (
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/service"
	"net/http"
)

type MerchantController struct {
	MerchantService service.MerchantService
}

func NewMerchantController(merchantService service.MerchantService) *MerchantController {
	return &MerchantController{
		MerchantService: merchantService,
	}
}

func (controller MerchantController) Create(ctx *gin.Context) {
	merchantCreateRequest := request.MerchantCreateRequest{}
	if err := ctx.ShouldBindJSON(&merchantCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	controller.MerchantService.Save(&merchantCreateRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Selamat!, akun merchant Anda berhasil didaftarkan",
	})
}

func (controller MerchantController) FindAll(ctx *gin.Context) {
	merchantResponses := controller.MerchantService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"data": merchantResponses,
	})
}

func (controller MerchantController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	merchantResponse := controller.MerchantService.FindByEmail(email)
	if merchantResponse == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "Mohon maaf, customer tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": merchantResponse,
	})
}

func (controller MerchantController) Update(ctx *gin.Context) {
	merchantUpdateRequest := request.MerchantUpdateRequest{}
	if err := ctx.ShouldBindJSON(&merchantUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	email := ctx.Param("email")
	merchantUpdateResponse := controller.MerchantService.Update(email, &merchantUpdateRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"data": merchantUpdateResponse,
	})
}

func (controller MerchantController) Delete(ctx *gin.Context) {
	email := ctx.Param("email")
	controller.MerchantService.Delete(email)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Akun merchant Anda berhasil dihapus",
	})
}
