package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
	Email       struct {
		ProviderA string
		ProviderB string
	}
	Push struct {
		ProviderA string
		ProviderB string
	}
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
	return &cfg
}
