package controller

import (
	"github.com/julienschmidt/httprouter"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
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

func (controller CustomerController) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	// read request
	customerRequest := request.CustomerCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &customerRequest)

	// write response
	controller.CustomerService.Create(&customerRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Selamat!, akun Anda berhasil didaftarkan",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CustomerController) FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	customerResponses := controller.CustomerService.FindAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CustomerController) FindByEmail(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	email := params.ByName("email")
	customerResponse := controller.CustomerService.FindByEmail(email)
	if nil != customerResponse {
		webResponse := response.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   customerResponse,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (controller CustomerController) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	customerUpdateRequest := request.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &customerUpdateRequest)

	email := params.ByName("email")
	customerUpdateResponse := controller.CustomerService.Update(email, &customerUpdateRequest)
	if nil != customerUpdateResponse {
		webResponse := response.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   customerUpdateResponse,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (controller CustomerController) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	email := params.ByName("email")
	controller.CustomerService.Delete(email)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Akun Anda berhasil dihapus",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
