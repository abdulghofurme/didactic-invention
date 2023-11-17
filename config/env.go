package config

import (
	"os"

	"github.com/abdulghofurme/go-mkr/helper"
	"github.com/joho/godotenv"
)

type ENV struct {
	SERVER_ADDRESS string
	DB_MYSQL_DSN   string
}

var MyENV *ENV

func init() {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	MyENV = &ENV{
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		DB_MYSQL_DSN:   os.Getenv("DB_MYSQL_DSN"),
	}
}
