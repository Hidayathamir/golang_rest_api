package config

import (
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type configuration struct {
	Database databaseConfiguration
}

type databaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var cfg *configuration

func initConfig() (*configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		logger.GetLog().Error(err)
		return nil, errors.Wrap(err, "config.initConfig")
	}

	cfg = &configuration{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.GetLog().Error(err)
		return nil, errors.Wrap(err, "config.initConfig")
	}

	logger.GetLog().Info("config.initConfig: init config success")
	return cfg, nil
}

func GetConfig() (*configuration, error) {
	if cfg == nil {
		return initConfig()
	}
	return cfg, nil
}
