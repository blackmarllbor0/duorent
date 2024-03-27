package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type ConfigService interface {
	LoadConfig() error
	GetAppConfig() AppConfig
	GetServerConfig() ServerConfig
	GetDBConfig() DBConfig
	GetLocalHash() LocalHash
}

type configService struct {
	app AppConfig
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (cs *configService) LoadConfig() error {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "" {
		runMode = "dev"
	}

	viper.SetConfigName("config." + runMode)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cfg: error on load configuration, err: %v", err)
	}

	var appConfig AppConfig

	if err := viper.UnmarshalKey("App", &appConfig); err != nil {
		return fmt.Errorf("cfg: error on unmarsahl app structure, err: %v", err)
	}

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
