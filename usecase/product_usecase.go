package usecase

import (
	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/Hidayathamir/golang_rest_api/repository"
	"github.com/pkg/errors"
)

type IProductUsecase interface {
	AddProduct(productInput dto.AddProductRequest) (dto.AddProductResponse, error)
	GetProducts(queryParam repository.GetProductsQueryParam) ([]dto.GetProductResponse, error)
}

type productUsecase struct {
	productRepository repository.IProductRepository
}

func NewProductUsecase(productRepository repository.IProductRepository) IProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (pu *productUsecase) AddProduct(productInput dto.AddProductRequest) (dto.AddProductResponse, error) {
	product, err := pu.productRepository.AddProduct(productInput)
	if err != nil {
		logger.GetLog().Error(err)
		return product, errors.Wrap(err, "usecase.AddProduct")
	}
	return product, nil
}

func (pu *productUsecase) GetProducts(queryParam repository.GetProductsQueryParam) ([]dto.GetProductResponse, error) {
	products, err := pu.productRepository.GetProducts(queryParam)
	if err != nil {
		logger.GetLog().Error(err)
		return products, errors.Wrap(err, "usecase.GetProducts")
	}
	return products, nil
}
