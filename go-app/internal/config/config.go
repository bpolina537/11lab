package config

import (
    "os"
    "strconv"
)

type Config struct {
    Port    int
    AppName string
}

func Load() *Config {
    port := 8080
    if p := os.Getenv("PORT"); p != "" {
        if v, err := strconv.Atoi(p); err == nil {
            port = v
        }
    }

    appName := os.Getenv("APP_NAME")
    if appName == "" {
        appName = "default-app"
    }

    return &Config{
        Port:    port,
        AppName: appName,
    }
}