package migration

import (
	"github.com/Hidayathamir/golang_rest_api/entity"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Product{})
	logger.GetLog().Info("migration.Automigrate: automigrate success")
}
