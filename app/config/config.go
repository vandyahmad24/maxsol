package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Rest struct {
		Port    int    `mapstructure:"port"`
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"rest"`
	Db struct {
		Address  string `mapstructure:"address"`
		DbName   string `mapstructure:"dbname"`
		Port     int    `mapstructure:"port"`
		Ssl      string `mapstructure:"ssl"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"db"`
}

func LoadConfig() (*Config, error) {
	// set config file name and location
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return &cfg, nil
}
