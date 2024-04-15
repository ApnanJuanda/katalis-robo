package controller

import (
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/service"
	"net/http"
)

type AddressController struct {
	AddressService service.AddressService
}

func NewAddressController(addressService service.AddressService) *AddressController {
	return &AddressController{
		AddressService: addressService,
	}
}

func (controller AddressController) Create(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	addressCreateRequest := request.AddressCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&addressCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "customer" {
		addressResponse := controller.AddressService.Create(authUser.UserEmail, &addressCreateRequest)
		if addressResponse != nil {
			ctx.JSON(http.StatusOK, addressResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": "Silahkan periksa kembali data product Anda",
			})
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"data": "Mohon Maaf, Anda tidak mempunyai akses",
		})
	}
}

func (controller AddressController) FindByIdAndEmail(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	addressCreateRequest := request.AddressCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&addressCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "customer" {
		addressId := ctx.Param("addressId")
		email := authUser.UserEmail
		addressResponse := controller.AddressService.FindByIdAndEmail(addressId, email)
		if addressResponse == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": "Mohon Maaf, produk tidak ditemukan",
			})
			return
		}
		ctx.JSON(http.StatusOK, addressResponse)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"data": "Mohon Maaf, Anda tidak mempunyai akses",
		})
	}
}

func (controller AddressController) FindByEmail(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	addressCreateRequest := request.AddressCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&addressCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "customer" {
		email := authUser.UserEmail
		addressResponses := controller.AddressService.FindByEmail(email)
		if addressResponses == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": "Mohon Maaf, produk tidak ditemukan",
			})
			return
		}
		ctx.JSON(http.StatusOK, addressResponses)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"data": "Mohon Maaf, Anda tidak mempunyai akses",
		})
	}
}

func (controller AddressController) Update(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	addressUpdateRequest := request.AddressCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&addressUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "customer" {
		addressId := ctx.Param("addressId")
		email := authUser.UserEmail
		productUpdateResponse := controller.AddressService.Update(addressId, email, &addressUpdateRequest)
		ctx.JSON(http.StatusOK, gin.H{
			"data": productUpdateResponse,
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}

func (controller AddressController) Delete(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	if authUser.Role == "customer" {
		addressId := ctx.Param("addressId")
		controller.AddressService.Delete(addressId, authUser.UserEmail)
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Alamat Anda berhasil dihapus",
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}
