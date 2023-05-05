package repository

import (
	"github.com/Hidayathamir/golang_rest_api/entity"
	"github.com/Hidayathamir/golang_rest_api/entity/dto"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IProductRepository interface {
	AddProduct(productInput dto.Product) (entity.Product, error)
	GetProducts() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) AddProduct(productInput dto.Product) (entity.Product, error) {
	var product entity.Product
	if err := copier.Copy(&product, &productInput); err != nil {
		logger.GetLog().Error(err)
		return entity.Product{}, errors.Wrap(err, "repository.AddProduct")
	}
	if err := pr.db.Create(&product).Error; err != nil {
		logger.GetLog().Error(err)
		return entity.Product{}, errors.Wrap(err, "repository.AddProduct")
	}
	return product, nil
}

func (pr *productRepository) GetProducts() ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Find(&products).Error; err != nil {
		logger.GetLog().Error(err)
		return []entity.Product{}, errors.Wrap(err, "repository.GetProduct")
	}
	return products, nil
}
