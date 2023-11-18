package controller

import (
	"net/http"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/web"
	"github.com/abdulghofurme/go-mkr/service"
	"github.com/julienschmidt/httprouter"
)

func NewHouseController(houseService service.HouseService) HouseController {
	return &HouseControllerImpl{
		HouseService: houseService,
	}
}

type HouseControllerImpl struct {
	HouseService service.HouseService
}

func (controller *HouseControllerImpl) Create(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	houseRequest := web.HouseCreateRequest{}
	helper.ReadFromRequestBody(request, &houseRequest)

	houseResponse := controller.HouseService.Create(
		request.Context(),
		houseRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   houseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HouseControllerImpl) Update(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	houseRequest := web.HouseUpdateRequest{}
	helper.ReadFromRequestBody(request, &houseRequest)
	houseRequest.ID = params.ByName("id")

	houseResponse := controller.HouseService.Update(
		request.Context(),
		houseRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   houseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HouseControllerImpl) Delete(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	houseId := params.ByName("id")

	controller.HouseService.Delete(
		request.Context(),
		houseId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *HouseControllerImpl) FindById(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	houseId := params.ByName("id")

	houseResponse := controller.HouseService.FindByID(
		request.Context(),
		houseId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   houseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *HouseControllerImpl) FindAll(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	housesResponse := controller.HouseService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   housesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
