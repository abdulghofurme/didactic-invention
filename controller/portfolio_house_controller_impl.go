package controller

import (
	"net/http"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/web"
	"github.com/abdulghofurme/go-mkr/service"
	"github.com/julienschmidt/httprouter"
)

func NewPortfolioHouseController(portfolioHouseService service.PortfolioHouseService) PortfolioHouseController {
	return &PortfolioHouseControllerImpl{
		PortfolioHouseService: portfolioHouseService,
	}
}

type PortfolioHouseControllerImpl struct {
	PortfolioHouseService service.PortfolioHouseService
}

func (controller *PortfolioHouseControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	portfolioHouseRequest := web.PortfolioHouseCreateRequest{}
	helper.ReadFromRequestBody(request, &portfolioHouseRequest)

	portfolioHouseResponse := controller.PortfolioHouseService.Create(
		request.Context(),
		portfolioHouseRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   portfolioHouseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioHouseControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	portfolioHouseRequest := web.PortfolioHouseUpdateRequest{}
	helper.ReadFromRequestBody(request, &portfolioHouseRequest)
	portfolioHouseRequest.ID = params.ByName("id")

	portfolioHouseResponse := controller.PortfolioHouseService.Update(
		request.Context(),
		portfolioHouseRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioHouseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioHouseControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	portfolioHouseId := params.ByName("id")

	portfolioHouseResponse := controller.PortfolioHouseService.Delete(
		request.Context(),
		portfolioHouseId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioHouseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioHouseControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	portfolioHouseId := params.ByName("id")

	portfolioHouseResponse := controller.PortfolioHouseService.FindByID(
		request.Context(),
		portfolioHouseId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioHouseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioHouseControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	portfolioHousesResponse := controller.PortfolioHouseService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioHousesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
