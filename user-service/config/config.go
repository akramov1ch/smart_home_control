package config

import (
    "os"
)

type Config struct {
    MongoURI   string
    RedisAddr  string
    RabbitMQURI string
}

func LoadConfig() *Config {
    return &Config{
        MongoURI:   os.Getenv("MONGO_URI"),
        RedisAddr:  os.Getenv("REDIS_ADDR"),
        RabbitMQURI: os.Getenv("RABBITMQ_URI"),
    }
}
