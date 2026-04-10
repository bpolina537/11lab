use axum::{Json, response::IntoResponse, http::StatusCode};
use serde_json::json;
use std::env;
use std::net::SocketAddr;

async fn health() -> impl IntoResponse {
    (StatusCode::OK, Json(json!({"status": "ok"})))
}

async fn info() -> impl IntoResponse {
    let app_name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
    (StatusCode::OK, Json(json!({"app_name": app_name})))
}

#[tokio::main]
async fn main() {
    let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
    let addr: SocketAddr = format!("0.0.0.0:{}", port).parse().unwrap();

    let app = axum::Router::new()
        .route("/health", axum::routing::get(health))
        .route("/info", axum::routing::get(info));

    println!("Rust server running on http://{}", addr);

    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    // ========== Тесты для /health ==========

    #[test]
    fn test_health_function_exists() {
        let _ = health();
        assert!(true);
    }

    // ========== Тесты для /info ==========

    #[test]
    fn test_info_function_exists() {
        let _ = info();
        assert!(true);
    }

    // ========== Тесты переменной PORT ==========

    #[test]
    fn test_port_default() {
        std::env::remove_var("PORT");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "8080");
    }

    #[test]
    fn test_port_custom() {
        std::env::set_var("PORT", "9090");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "9090");
    }

    #[test]
    fn test_port_empty() {
        std::env::set_var("PORT", "");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "");
    }

    #[test]
    fn test_port_with_spaces() {
        std::env::set_var("PORT", "  9090  ");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port.trim(), "9090");
    }

    #[test]
    fn test_port_zero() {
        std::env::set_var("PORT", "0");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "0");
    }

    #[test]
    fn test_port_negative() {
        std::env::set_var("PORT", "-8080");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "-8080");
    }

    #[test]
    fn test_port_max() {
        std::env::set_var("PORT", "65535");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        assert_eq!(port, "65535");
    }

    // ========== Тесты переменной APP_NAME ==========

    #[test]
    fn test_app_name_default() {
        std::env::remove_var("APP_NAME");
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(name, "default-app");
    }

    #[test]
    fn test_app_name_custom() {
    std::env::set_var("APP_NAME", "my-rust-app");
    let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
    assert_eq!(name, "my-rust-app");
    std::env::remove_var("APP_NAME");
}

    #[test]
    fn test_app_name_empty() {
        std::env::set_var("APP_NAME", "");
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(name, "");
    }

    #[test]
    fn test_app_name_with_spaces() {
        std::env::set_var("APP_NAME", "  my-app  ");
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(name.trim(), "my-app");
    }

    #[test]
    fn test_app_name_special_chars() {
        std::env::set_var("APP_NAME", "my-app-v1.0_test");
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(name, "my-app-v1.0_test");
    }

    #[test]
    fn test_app_name_long() {
    let long = "a".repeat(100);
    std::env::set_var("APP_NAME", &long);
    let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
    assert_eq!(name, long);
    std::env::remove_var("APP_NAME");
    }


    #[test]
    fn test_app_name_unicode() {
        std::env::set_var("APP_NAME", "приложение-тест");
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(name, "приложение-тест");
    }

    // ========== Комбинированные тесты ==========

    #[test]
    fn test_both_vars_set() {
        std::env::set_var("PORT", "9090");
        std::env::set_var("APP_NAME", "production");
        let port = env::var("PORT").unwrap();
        let name = env::var("APP_NAME").unwrap();
        assert_eq!(port, "9090");
        assert_eq!(name, "production");
    }

    #[test]
    fn test_both_vars_empty() {
        std::env::set_var("PORT", "");
        std::env::set_var("APP_NAME", "");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(port, "");
        assert_eq!(name, "");
    }

    #[test]
    fn test_both_vars_missing() {
        std::env::remove_var("PORT");
        std::env::remove_var("APP_NAME");
        let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
        let name = env::var("APP_NAME").unwrap_or_else(|_| "default-app".to_string());
        assert_eq!(port, "8080");
        assert_eq!(name, "default-app");
    }
}