package handler

import (
	"net/http"

	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	"github.com/Hidayathamir/golang_rest_api/handler/helper"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/Hidayathamir/golang_rest_api/repository"
	"github.com/Hidayathamir/golang_rest_api/usecase"
	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	AddProduct(c *gin.Context)
	GetProducts(c *gin.Context)
}

type productHandler struct {
	productUsecase usecase.IProductUsecase
}

func NewProductHandler(productUsecase usecase.IProductUsecase) IProductHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (ph *productHandler) AddProduct(c *gin.Context) {
	var productInput dto.AddProductRequest

	if err := c.ShouldBindJSON(&productInput); err != nil {
		helper.WriteRespon(c, http.StatusBadRequest, err.Error(), nil)
		logger.GetLog().Error(err)
		return
	}

	product, err := ph.productUsecase.AddProduct(productInput)
	if err != nil {
		helper.WriteRespon(c, http.StatusBadRequest, err.Error(), nil)
		logger.GetLog().Error(err)
		return
	}

	helper.WriteRespon(c, http.StatusOK, "add product success", product)
}

func (ph *productHandler) GetProducts(c *gin.Context) {
	queryParam := repository.GetProductsQueryParam{
		SortBy: c.DefaultQuery("sort_by", "updated_at"),
		Sort:   c.DefaultQuery("sort", "desc"),
	}

	products, err := ph.productUsecase.GetProducts(queryParam)
	if err != nil {
		helper.WriteRespon(c, http.StatusBadRequest, err.Error(), nil)
		logger.GetLog().Error(err)
		return
	}

	helper.WriteRespon(c, http.StatusOK, "get products success", products)
}
