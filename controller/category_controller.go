package controller

import (
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/service"
	"net/http"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (controller CategoryController) Create(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	categoryCreateRequest := request.CategoryCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&categoryCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "merchant" {
		categoryResponse := controller.CategoryService.Create(&categoryCreateRequest)
		if categoryResponse != nil {
			ctx.JSON(http.StatusOK, categoryResponse)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": "Silahkan periksa kembali data category Anda",
			})
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"data": "Mohon Maaf, Anda tidak mempunyai akses",
		})
	}
}

func (controller CategoryController) FindById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	categoryResponse := controller.CategoryService.FindById(categoryId)
	if categoryResponse == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "Mohon maaf, category tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": categoryResponse,
	})
}

func (controller CategoryController) Update(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	categoryUpdateRequest := request.CategoryCreateUpdateRequest{}
	if err := ctx.ShouldBindJSON(&categoryUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	if authUser.Role == "merchant" {
		categoryId := ctx.Param("categoryId")
		categoryUpdateResponse := controller.CategoryService.Update(categoryId, &categoryUpdateRequest)
		ctx.JSON(http.StatusOK, gin.H{
			"data": categoryUpdateResponse,
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}

func (controller CategoryController) Delete(ctx *gin.Context) {
	authUser := ctx.Request.Context().Value("authUser").(current.AuthUser)
	if authUser.Role == "merchant" {
		// recheck product's merchant
		categoryId := ctx.Param("categoryId")
		controller.CategoryService.Delete(categoryId)
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Category berhasil dihapus",
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"data": "Mohon Maaf, Anda tidak mempunyai akses",
	})
}
