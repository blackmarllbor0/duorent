package config

import (
	"duorent.ru/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
)

type ConfigService interface {
	LoadConfig(runMode string) error
	GetAppConfig() AppConfig
	GetServerConfig() ServerConfig
	GetDBConfig() DBConfig
	GetLocalHash() LocalHash
}

type configService struct {
	app        AppConfig
	logService logger.LoggerService
}

func NewConfigService(logService logger.LoggerService) ConfigService {
	return &configService{logService: logService}
}

func (cs *configService) LoadConfig(runMode string) error {
	viper.SetConfigName("config." + runMode)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("cfg: error on load configuration, err: %v", err)
		cs.logService.Error(err.Error())

		return err
	}

	go cs.logService.Info("Configuration successfully loaded")

	var appConfig AppConfig

	if err := viper.UnmarshalKey("App", &appConfig); err != nil {
		err = fmt.Errorf("cfg: error on unmarsahl app structure, err: %v", err)
		cs.logService.Error(err.Error())

		return err
	}

	go cs.logService.Info("Configuration successfully unmarshalling")

	cs.app = appConfig

	return nil
}

func (cs *configService) GetAppConfig() AppConfig {
	return cs.app
}

func (cs *configService) GetServerConfig() ServerConfig {
	return cs.app.Server
}

func (cs *configService) GetDBConfig() DBConfig {
	return cs.app.DB
}

func (cs *configService) GetLocalHash() LocalHash {
	return cs.app.LocalHash
}
