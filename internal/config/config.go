package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port uint `yaml:"Port"`
}

type DBConfig struct {
	ConnString string `yaml:"ConnString"`
}

type ConfigService interface {
	LoadConfig() error
	GetAppConfig() AppConfig
	GetServerConfig() ServerConfig
	GetDBConfig() DBConfig
}

type configService struct {
	app AppConfig
}

func NewConfigService() ConfigService {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "" {
		runMode = "dev"
	}

	viper.SetConfigName("config." + runMode)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	return &configService{}
}

func (cs *configService) LoadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cfg: error on load configuration, err: %v", err)
	}

	var appConfig AppConfig

	if err := viper.Unmarshal(&appConfig); err != nil {
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
