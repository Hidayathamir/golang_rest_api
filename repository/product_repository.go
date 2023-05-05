package repository

import (
	"fmt"

	"github.com/Hidayathamir/golang_rest_api/entity"
	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IProductRepository interface {
	AddProduct(productInput dto.AddProductRequest) (dto.AddProductResponse, error)
	GetProducts(queryParam GetProductsQueryParam) ([]dto.GetProductResponse, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) AddProduct(productInput dto.AddProductRequest) (dto.AddProductResponse, error) {
	var product entity.Product

	if err := copier.Copy(&product, &productInput); err != nil {
		logger.GetLog().Error(err)
		return dto.AddProductResponse{}, errors.Wrap(err, "repository.AddProduct")
	}

	if err := pr.db.Create(&product).Error; err != nil {
		logger.GetLog().Error(err)
		return dto.AddProductResponse{}, errors.Wrap(err, "repository.AddProduct")
	}

	var productResponse dto.AddProductResponse
	if err := copier.Copy(&productResponse, &product); err != nil {
		logger.GetLog().Error(err)
		return dto.AddProductResponse{}, errors.Wrap(err, "repository.AddProduct")
	}

	return productResponse, nil
}

type GetProductsQueryParam struct {
	SortBy string
	Sort   string
}

func (pr *productRepository) GetProducts(queryParam GetProductsQueryParam) ([]dto.GetProductResponse, error) {
	var products []entity.Product

	chainMethod := pr.db.Order(fmt.Sprintf("%s %s", queryParam.SortBy, queryParam.Sort))
	if err := chainMethod.Find(&products).Error; err != nil {
		logger.GetLog().Error(err)
		return []dto.GetProductResponse{}, errors.Wrap(err, "repository.GetProduct")
	}

	var productsResponse []dto.GetProductResponse
	if err := copier.Copy(&productsResponse, &products); err != nil {
		logger.GetLog().Error(err)
		return []dto.GetProductResponse{}, errors.Wrap(err, "repository.GetProduct")
	}

	return productsResponse, nil
}
