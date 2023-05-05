package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	"github.com/Hidayathamir/golang_rest_api/handler"
	mocks "github.com/Hidayathamir/golang_rest_api/mocks/usecase"
	"github.com/Hidayathamir/golang_rest_api/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_AddProduct_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	addProductReq := dto.AddProductRequest{
		Name:        "dummy product name",
		Price:       123456,
		Description: "dummy product desc",
		Quantity:    1111,
	}
	req, _ := json.Marshal(addProductReq)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(req))

	productUsecase := mocks.NewIProductUsecase(t)
	productHandler := handler.NewProductHandler(productUsecase)

	productUsecase.On("AddProduct", addProductReq).Return(dto.AddProductResponse{
		Name:        "dummy product name res",
		Price:       123456,
		Description: "dummy product desc res",
		Quantity:    1111,
	}, nil)
	productHandler.AddProduct(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_AddProduct_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	addProductReq := dto.AddProductRequest{
		Name:        "dummy product name",
		Price:       123456,
		Description: "dummy product desc",
		Quantity:    1111,
	}
	req, _ := json.Marshal(addProductReq)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(req))

	productUsecase := mocks.NewIProductUsecase(t)
	productHandler := handler.NewProductHandler(productUsecase)

	productUsecase.On("AddProduct", addProductReq).Return(dto.AddProductResponse{}, errors.New("dummy error"))
	productHandler.AddProduct(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_AddProduct_Error_MissingRequiredRequestBody(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	addProductReq := dto.AddProductRequest{
		Name:        "dummy product name",
		Description: "dummy product desc",
		Quantity:    1111,
	}
	req, _ := json.Marshal(addProductReq)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(req))

	productUsecase := mocks.NewIProductUsecase(t)
	productHandler := handler.NewProductHandler(productUsecase)

	productHandler.AddProduct(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_GetProducts_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	c.Request = httptest.NewRequest(http.MethodGet, "/products", nil)

	productUsecase := mocks.NewIProductUsecase(t)
	productHandler := handler.NewProductHandler(productUsecase)

	queryParam := repository.GetProductsQueryParam{
		SortBy: "updated_at",
		Sort:   "desc",
	}
	productUsecase.On("GetProducts", queryParam).Return([]dto.GetProductResponse{}, nil)
	productHandler.GetProducts(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_GetProducts_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	c.Request = httptest.NewRequest(http.MethodGet, "/products", nil)

	productUsecase := mocks.NewIProductUsecase(t)
	productHandler := handler.NewProductHandler(productUsecase)

	queryParam := repository.GetProductsQueryParam{
		SortBy: "updated_at",
		Sort:   "desc",
	}
	productUsecase.On("GetProducts", queryParam).Return([]dto.GetProductResponse{}, errors.New("dummy error"))
	productHandler.GetProducts(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
