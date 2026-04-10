Лабораторная работа №11
Контейнеризация мультиязычных приложений

Студент: Бондаренко Полина Кирилловна
Группа: 221331
Вариант: 10
Дата выполнения: 10.04.2026

Цель
Освоить упаковку в контейнеры приложений на Go и Rust, оптимизировать размер образов, использовать переменные окружения и кросс-платформенную сборку.

Выполненные задания

Средняя сложность
10. Использование переменных окружения для конфигурации (PORT, APP_NAME)
2. Многоэтапный Dockerfile для Go-приложения (scratch, 6.61 MB)
4. Сравнение размеров образов (simple: 23.6 MB, multistage: 6.61 MB)

Повышенная сложность
2. Rust-приложение с поддержкой musl (полностью статическая сборка, 2.59 MB)
4. Кросс-платформенная сборка через docker buildx (amd64/arm64)

Результаты оптимизации образов

| Сервис | До оптимизации | После оптимизации | Уменьшение |
|--------|----------------|-------------------|------------|
| Go (simple) | 23.6 MB | 6.61 MB | -72% |
| Rust (musl) | ~1.5 GB | 2.59 MB | -99% |

docker buildx

Образы собраны для платформ:
- linux/amd64 (go-app:amd64) — 6.61 MB
- linux/arm64 (go-app:arm64) — 7.85 MB

Тесты

| Компонент | Количество тестов |
|-----------|-------------------|
| Go (config) | 13 |
| Rust | 19 |
| Всего | 32 |

Быстрый старт

Go-приложение
cd go-app
go run ./cmd/server

Через Docker:
docker build -f Dockerfile.multistage -t go-app:multistage .
docker run -e PORT=9090 -e APP_NAME=test -p 9090:9090 go-app:multistage

Rust-приложение
cd rust-app
cargo run

Через Docker:
docker build -f Dockerfile.multistage -t rust-app:musl .
docker run -e PORT=9090 -e APP_NAME=test -p 9090:9090 rust-app:musl

Проверка работы
curl http://localhost:8080/health
curl http://localhost:8080/info

Структура проекта

11lab/
├── go-app/
│   ├── cmd/server/main.go
│   ├── internal/config/
│   │   ├── config.go
│   │   └── config_test.go
│   ├── Dockerfile
│   ├── Dockerfile.multistage
│   └── go.mod
├── rust-app/
│   ├── src/main.rs
│   ├── Dockerfile.multistage
│   └── Cargo.toml
├── benchmarks/
│   ├── image_sizes.txt
│   ├── rust_image_sizes.txt
│   └── buildx_results.txt
├── .gitignore
├── README.md
└── PROMPT_LOG.md

Ссылка на репозиторий
https://github.com/bpolina537/11lab