package app

import (
	"database/sql"
	"time"

	"github.com/abdulghofurme/go-mkr/config"
	"github.com/abdulghofurme/go-mkr/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", config.MyENV.DB_MYSQL_DSN)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(2)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	return db
}
