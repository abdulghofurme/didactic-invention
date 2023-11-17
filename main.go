package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/abdulghofurme/go-mkr/app"
	"github.com/abdulghofurme/go-mkr/config"
	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	app.NewDB()

	server := http.Server{
		Addr:    config.MyENV.SERVER_ADDRESS,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
