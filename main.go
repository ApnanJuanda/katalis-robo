package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"katalisRobo/component-store/app"
	"katalisRobo/component-store/controller"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/middleware"
	"katalisRobo/component-store/repository"
	"katalisRobo/component-store/service"
	"os"
)

func main() {
	// make sure success load env
	err := godotenv.Load("config/.env")
	helper.PanicIfError(err)

	router := gin.Default()

	db := app.NewDB()
	validate := validator.New()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)

	merchantRepository := repository.NewMerchantRepository(db)
	merchantService := service.NewMerchantService(merchantRepository, validate)
	merchantController := controller.NewMerchantController(merchantService)

	userService := service.NewUserService(customerRepository, merchantRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	groupRepository := repository.NewGroupRepository(db)
	productService := service.NewProductService(productRepository, categoryRepository, merchantRepository, groupRepository, validate)
	productController := controller.NewProductController(productService)

	addressRepository := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepository, customerRepository, validate)
	addressController := controller.NewAddressController(addressService)

	// Customer
	router.POST("/api/customers", customerController.Create)
	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:email", customerController.FindByEmail)
	router.PUT("/api/customers/:email", customerController.Update)
	router.DELETE("/api/customers/:email", customerController.Delete)

	// Merchant
	router.POST("/api/merchants", merchantController.Create)
	router.GET("/api/merchants", merchantController.FindAll)
	router.GET("/api/merchants/:email", merchantController.FindByEmail)
	router.PUT("/api/merchants/:email", merchantController.Update)
	router.DELETE("/api/merchants/:email", merchantController.Delete)

	// User
	router.POST("/api/login", userController.Login)

	// Product
	router.POST("/api/products", middleware.WithAuth(), productController.Create)
	router.GET("/api/products/:productId", productController.FindById)
	router.GET("/api/products/merchant/:merchantId", productController.FindByMerchantId)
	router.PUT("/api/products/:productId", middleware.WithAuth(), productController.Update)
	router.DELETE("/api/products/:productId", middleware.WithAuth(), productController.Delete)

	// Address
	router.POST("/api/address", middleware.WithAuth(), addressController.Create)
	router.GET("/api/address/:addressId", middleware.WithAuth(), addressController.FindByIdAndEmail)
	router.GET("/api/address", middleware.WithAuth(), addressController.FindByEmail)
	router.PUT("/api/address/:addressId", middleware.WithAuth(), addressController.Update)
	router.DELETE("/api/address/:addressId", middleware.WithAuth(), addressController.Delete)

	fmt.Println("My Application is running")
	router.Run(":" + os.Getenv("PORT"))
}
