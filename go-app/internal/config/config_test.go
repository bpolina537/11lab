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

func TestLoadEmptyPort(t *testing.T) {
    os.Setenv("PORT", "")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != 8080 {
        t.Errorf("Expected fallback to 8080, got %d", cfg.Port)
    }
}

func TestLoadEmptyAppName(t *testing.T) {
    os.Setenv("APP_NAME", "")
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.AppName != "default-app" {
        t.Errorf("Expected default-app, got %s", cfg.AppName)
    }
}

func TestLoadPortZero(t *testing.T) {
    os.Setenv("PORT", "0")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != 0 {
        t.Errorf("Expected 0, got %d", cfg.Port)
    }
}

func TestLoadPortNegative(t *testing.T) {
    os.Setenv("PORT", "-8080")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != -8080 {
        t.Errorf("Expected -8080, got %d", cfg.Port)
    }
}

func TestLoadPortVeryLarge(t *testing.T) {
    os.Setenv("PORT", "65535")
    defer os.Unsetenv("PORT")
    cfg := Load()
    if cfg.Port != 65535 {
        t.Errorf("Expected 65535, got %d", cfg.Port)
    }
}

func TestLoadAppNameSpecialChars(t *testing.T) {
    os.Setenv("APP_NAME", "my-app-v1.0")
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.AppName != "my-app-v1.0" {
        t.Errorf("Expected my-app-v1.0, got %s", cfg.AppName)
    }
}

func TestLoadAppNameLong(t *testing.T) {
    longName := "this-is-a-very-long-application-name-for-testing-purposes"
    os.Setenv("APP_NAME", longName)
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.AppName != longName {
        t.Errorf("Expected %s, got %s", longName, cfg.AppName)
    }
}

func TestLoadAppNameUnicode(t *testing.T) {
    os.Setenv("APP_NAME", "приложение-тест")
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.AppName != "приложение-тест" {
        t.Errorf("Expected приложение-тест, got %s", cfg.AppName)
    }
}

func TestLoadConfigBothSet(t *testing.T) {
    os.Setenv("PORT", "9090")
    os.Setenv("APP_NAME", "production-app")
    defer os.Unsetenv("PORT")
    defer os.Unsetenv("APP_NAME")
    cfg := Load()
    if cfg.Port != 9090 || cfg.AppName != "production-app" {
        t.Errorf("Expected 9090/production-app, got %d/%s", cfg.Port, cfg.AppName)
    }
}