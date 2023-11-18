package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/abdulghofurme/go-mkr/app"
	"github.com/abdulghofurme/go-mkr/config"
	"github.com/abdulghofurme/go-mkr/controller"
	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/abdulghofurme/go-mkr/repository"
	"github.com/abdulghofurme/go-mkr/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	db := app.NewDB()
	validate := validator.New()

	houseRepository := repository.NewHouseRepository()
	houseService := service.NewHouseService(
		houseRepository,
		db,
		validate,
	)
	houseController := controller.NewHouseController(houseService)

	portfolioRepository := repository.NewPortfolioRepository()
	portfolioService := service.NewPortfolioService(
		portfolioRepository,
		db,
		validate,
	)
	portfolioController := controller.NewPortfolioController(portfolioService)

	portfolioHouseRepository := repository.NewPortfolioHouseRepository()
	portfolioHouseService := service.NewPortfolioHouseService(
		portfolioHouseRepository,
		db,
		validate,
	)
	portfolioHouseController := controller.NewPortfolioHouseController(portfolioHouseService)

	router.POST("/api/houses", houseController.Create)
	router.GET("/api/houses", houseController.FindAll)
	router.GET("/api/houses/:id", houseController.FindById)
	router.PUT("/api/houses/:id", houseController.Update)
	router.DELETE("/api/houses/:id", houseController.Delete)

	router.POST("/api/portfolios", portfolioController.Create)
	router.GET("/api/portfolios", portfolioController.FindAll)
	router.GET("/api/portfolios/:id", portfolioController.FindById)
	router.PUT("/api/portfolios/:id", portfolioController.Update)
	router.DELETE("/api/portfolios/:id", portfolioController.Delete)

	router.POST("/api/portfolio-houses", portfolioHouseController.Create)
	router.GET("/api/portfolio-houses", portfolioHouseController.FindAll)
	router.GET("/api/portfolio-houses/:id", portfolioHouseController.FindById)
	router.PUT("/api/portfolio-houses/:id", portfolioHouseController.Update)
	router.DELETE("/api/portfolio-houses/:id", portfolioHouseController.Delete)

	server := http.Server{
		Addr:    config.MyENV.SERVER_ADDRESS,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
