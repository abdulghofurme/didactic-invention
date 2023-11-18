package controller

import (
	"net/http"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/model/web"
	"github.com/abdulghofurme/go-mkr/service"
	"github.com/julienschmidt/httprouter"
)

func NewPortfolioController(portfolioService service.PortfolioService) PortfolioController {
	return &PortfolioControllerImpl{
		PortfolioService: portfolioService,
	}
}

type PortfolioControllerImpl struct {
	PortfolioService service.PortfolioService
}

func (controller *PortfolioControllerImpl) Create(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	portfolioRequest := web.PortfolioCreateRequest{}
	helper.ReadFromRequestBody(request, &portfolioRequest)

	portfolioResponse := controller.PortfolioService.Create(
		request.Context(),
		portfolioRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   portfolioResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioControllerImpl) Update(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	portfolioRequest := web.PortfolioUpdateRequest{}
	helper.ReadFromRequestBody(request, &portfolioRequest)
	portfolioRequest.ID = params.ByName("id")

	portfolioResponse := controller.PortfolioService.Update(
		request.Context(),
		portfolioRequest,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioControllerImpl) Delete(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	portfolioId := params.ByName("id")

	portfolioResponse := controller.PortfolioService.Delete(
		request.Context(),
		portfolioId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioControllerImpl) FindById(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	portfolioId := params.ByName("id")

	portfolioResponse := controller.PortfolioService.FindByID(
		request.Context(),
		portfolioId,
	)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfolioResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PortfolioControllerImpl) FindAll(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params,
) {
	portfoliosResponse := controller.PortfolioService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   portfoliosResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
