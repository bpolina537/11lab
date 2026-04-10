package main

import (
    "fmt"
    "log"
    "net/http"
    "go-app/internal/config"
)

func main() {
    cfg := config.Load()

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok"}`))
    })

    http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(fmt.Sprintf(`{"app_name":"%s"}`, cfg.AppName)))
    })

    addr := fmt.Sprintf(":%d", cfg.Port)
    log.Printf("Server starting on %s", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}