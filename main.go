package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"katalisRobo/component-store/app"
	"katalisRobo/component-store/controller"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/repository"
	"katalisRobo/component-store/service"
	"net/http"
	"os"
)

func main() {
	// make sure success load env
	err := godotenv.Load("config/.env")
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)

	merchantRepository := repository.NewMerchantRepository(db)
	merchantService := service.NewMerchantService(merchantRepository, validate)
	merchantController := controller.NewMerchantController(merchantService)

	router := httprouter.New()

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

	fmt.Println("My Application is running")
	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
