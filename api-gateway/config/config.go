package config

import (
    "os"
)

type Config struct {
    Port                string
    UserServiceAddress  string
    DeviceServiceAddress string
    JWTSecret           string
    MongoURI            string
    RabbitMQURI         string
}

func GetConfig() *Config {
    return &Config{
        Port:                os.Getenv("PORT"),
        UserServiceAddress:  os.Getenv("USER_SERVICE_ADDRESS"),
        DeviceServiceAddress: os.Getenv("DEVICE_SERVICE_ADDRESS"),
        JWTSecret:           os.Getenv("JWT_SECRET"),
        MongoURI:            os.Getenv("MONGO_URI"),
        RabbitMQURI:         os.Getenv("RABBITMQ_URI"),
    }
}
