package database

import (
	"fmt"

	"github.com/Hidayathamir/golang_rest_api/config"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() (*gorm.DB, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.GetLog().Error(err)
		return nil, errors.Wrap(err, "database.initDB")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.GetLog().Error(err)
		return nil, errors.Wrapf(err, "database.initDB please check your database config in config/config.json")
	}

	logger.GetLog().Info("database.initDB: init connection to database ", cfg.Database.DBName, "success")
	return db, nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return initDB()
	}
	return db, nil
}
