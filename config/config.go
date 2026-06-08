package config

import (
    "errors"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    AppEnv string
    Port   string

    DatabaseURL string

    AWSRegion string
    CURBucket string
}

func Load() (*Config, error) {
    _ = godotenv.Load()

    cfg := &Config{
        AppEnv:      os.Getenv("APP_ENV"),
        Port:        os.Getenv("PORT"),
        DatabaseURL: os.Getenv("DATABASE_URL"),
        AWSRegion:   os.Getenv("AWS_REGION"),
        CURBucket:   os.Getenv("CUR_BUCKET"),
    }

    if cfg.Port == "" {
        return nil, errors.New("PORT is required")
    }

    if cfg.DatabaseURL == "" {
        return nil, errors.New("DATABASE_URL is required")
    }

    return cfg, nil
}