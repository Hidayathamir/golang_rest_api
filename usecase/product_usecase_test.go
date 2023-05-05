package usecase_test

import (
	"errors"
	"testing"

	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	mocks "github.com/Hidayathamir/golang_rest_api/mocks/repository"
	"github.com/Hidayathamir/golang_rest_api/repository"
	"github.com/Hidayathamir/golang_rest_api/usecase"
	"github.com/stretchr/testify/assert"
)

func Test_AddProduct_Success(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	productUsecase := usecase.NewProductUsecase(productRepository)

	var productInput dto.AddProductRequest
	productRepository.On("AddProduct", productInput).Return(dto.AddProductResponse{Name: "dummy_name"}, nil)
	product, _ := productUsecase.AddProduct(productInput)

	assert.Equal(t, "dummy_name", product.Name)
}

func Test_AddProduct_Error(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	productUsecase := usecase.NewProductUsecase(productRepository)

	var productInput dto.AddProductRequest
	productRepository.On("AddProduct", productInput).Return(dto.AddProductResponse{}, errors.New("dummy error"))
	_, err := productUsecase.AddProduct(productInput)

	assert.Equal(t, "usecase.AddProduct: dummy error", err.Error())
}

func Test_GetProducts_Success(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	productUsecase := usecase.NewProductUsecase(productRepository)

	var queryParam repository.GetProductsQueryParam
	productRepository.On("GetProducts", queryParam).Return([]dto.GetProductResponse{
		{Name: "dummy_name"},
		{Name: "dummy_name 2"},
	}, nil)
	products, _ := productUsecase.GetProducts(queryParam)

	assert.Equal(t, "dummy_name", products[0].Name)
	assert.Equal(t, "dummy_name 2", products[1].Name)
}

func Test_GetProducts_Error(t *testing.T) {
	productRepository := mocks.NewIProductRepository(t)
	productUsecase := usecase.NewProductUsecase(productRepository)

	var queryParam repository.GetProductsQueryParam
	productRepository.On("GetProducts", queryParam).Return([]dto.GetProductResponse{}, errors.New("dummy error"))
	_, err := productUsecase.GetProducts(queryParam)

	assert.Equal(t, "usecase.GetProducts: dummy error", err.Error())
}
