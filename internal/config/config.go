package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Driver   string
	Dbfolder string
	Username string
	Password string
}

func GetConfig(filepath string) (*Config, error) {
	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func InitConfig(filepath string) *Config {
	cfg, err := GetConfig(filepath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфиг файла: %s", err)
	}
	return cfg
}
