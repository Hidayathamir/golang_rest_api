package main

import (
	"github.com/Hidayathamir/golang_rest_api/database"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/Hidayathamir/golang_rest_api/migration"
	"github.com/Hidayathamir/golang_rest_api/router"
	"github.com/pkg/errors"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		logger.GetLog().Error(err)
		panic(errors.Wrap(err, "main.main"))
	}
	migration.Automigrate(db)
	r := router.GetRouter()
	r.Run()
}
