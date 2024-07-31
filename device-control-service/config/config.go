package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort string `mapstructure:"GRPC_PORT"`
	MongoURI string `mapstructure:"MONGO_URI"`
	RedisURI string `mapstructure:"REDIS_URI"`
	RabbitMQURI string `mapstructure:"RABBITMQ_URI"`
}

func LoadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error while unmarshaling config %s", err)
	}

	return config
}
