package config

import (
    "log"

    "github.com/joho/godotenv"
)

func LoadConfig(configPath string) {
    if configPath == "" {
        configPath = ".env"
    }

    err := godotenv.Load(configPath)
    if err != nil {
        log.Fatalf("Ошибка загрузки файла конфигурации: %v", err)
    }
}
