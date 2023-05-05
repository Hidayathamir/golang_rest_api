package main

import (
	"github.com/Hidayathamir/golang_rest_api/database"
	"github.com/Hidayathamir/golang_rest_api/handler"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/Hidayathamir/golang_rest_api/migration"
	"github.com/Hidayathamir/golang_rest_api/repository"
	"github.com/Hidayathamir/golang_rest_api/router"
	"github.com/Hidayathamir/golang_rest_api/usecase"
	"github.com/pkg/errors"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		logger.GetLog().Error(err)
		panic(errors.Wrap(err, "main.main"))
	}
	migration.Automigrate(db)

	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	allHandler := handler.AllHandler{
		ProductHandler: productHandler,
	}

	r := router.GetRouter(allHandler)
	r.Run()
}
