package controller

import (
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/service"
	"net/http"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (controller ProductController) Create(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	productCreateRequest := request.ProductCreateRequest{}
	if err := ctx.ShouldBindJSON(&productCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "merchant" {
		productResponse := controller.ProductService.Create(&productCreateRequest, &authUser)
		if productResponse != nil {
			ctx.JSON(http.StatusOK, productResponse)
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

func (controller ProductController) FindById(ctx *gin.Context) {
	productId := ctx.Param("productId")
	productResponse := controller.ProductService.FindById(productId)
	if productResponse == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "Mohon Maaf, produk tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, productResponse)
}

func (controller ProductController) FindByMerchantId(ctx *gin.Context) {
	merchantId := ctx.Param("merchantId")
	productResponses := controller.ProductService.FindByMerchantId(merchantId)
	if productResponses == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "Mohon Maaf, produk tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, productResponses)
}

func (controller ProductController) Update(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	productUpdateRequest := request.ProductUpdateRequest{}
	if err := ctx.ShouldBindJSON(&productUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "merchant" {
		productId := ctx.Param("productId")
		productUpdateResponse := controller.ProductService.Update(productId, &productUpdateRequest)
		ctx.JSON(http.StatusOK, gin.H{
			"data": productUpdateResponse,
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}

func (controller ProductController) Delete(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	if authUser.Role == "merchant" {
		// recheck product's merchant
		productId := ctx.Param("productId")
		controller.ProductService.Delete(productId)
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Product Anda berhasil dihapus",
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}
