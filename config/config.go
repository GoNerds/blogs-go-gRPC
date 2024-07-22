package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config sturct {
	MongoURI string
	GRPCPort string
	LogLevel string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	return &Config{
		MongoURI: viper.GetString("mongo_uri"),
		GRPCPort: viper.GetString("grpc_port"),
		LogLevel: viper.GetString("log_level"),
	}
}