package controller

import (
	"github.com/julienschmidt/httprouter"
	"katalisRobo/component-store/data/request"
	"katalisRobo/component-store/data/response"
	"katalisRobo/component-store/helper"
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

func (controller MerchantController) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	merchantCreateRequest := request.MerchantCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &merchantCreateRequest)

	controller.MerchantService.Save(&merchantCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Selamat!, akun supplier Anda berhasil didaftarkan",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller MerchantController) FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	merchantResponses := controller.MerchantService.FindAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   merchantResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller MerchantController) FindByEmail(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	email := params.ByName("email")
	merchantResponse := controller.MerchantService.FindByEmail(email)
	if nil != merchantResponse {
		webResponse := response.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   merchantResponse,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (controller MerchantController) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	merchantUpdateRequest := request.MerchantUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &merchantUpdateRequest)

	email := params.ByName("email")
	merchantUpdateResponse := controller.MerchantService.Update(email, &merchantUpdateRequest)
	if nil != merchantUpdateResponse {
		webResponse := response.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   merchantUpdateResponse,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (controller MerchantController) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	email := params.ByName("email")
	controller.MerchantService.Delete(email)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Akun Anda berhasil dihapus",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
