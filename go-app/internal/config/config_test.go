package config

import (
    "os"
    "testing"
)

func TestLoadDefaultPort(t *testing.T) {
    os.Unsetenv("PORT")
    os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.Port != 8080 {
        t.Errorf("Expected 8080, got %d", cfg.Port)
    }
    if cfg.AppName != "default-app" {
        t.Errorf("Expected default-app, got %s", cfg.AppName)
    }
}

func TestLoadCustomPort(t *testing.T) {
    os.Setenv("PORT", "9090")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != 9090 {
        t.Errorf("Expected 9090, got %d", cfg.Port)
    }
}

func TestLoadCustomAppName(t *testing.T) {
    os.Setenv("APP_NAME", "my-app")
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.AppName != "my-app" {
        t.Errorf("Expected my-app, got %s", cfg.AppName)
    }
}

func TestLoadInvalidPort(t *testing.T) {
    os.Setenv("PORT", "invalid")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != 8080 {
        t.Errorf("Expected fallback to 8080, got %d", cfg.Port)
    }
}