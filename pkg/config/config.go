package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
