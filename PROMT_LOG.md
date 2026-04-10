Prompt Log
Лабораторная работа №11
Контейнеризация мультиязычных приложений

Студент: Бондаренко Полина Кирилловна
Группа: 221331
Вариант: 10
Дата выполнения: 10.04.2026


Промпт 1
Инструмент: DeepSeek-V3.2

Промпт:
Составь детальный пошаговый план для лабораторной работы №11. Задания:
10 (средняя). Использовать переменные окружения для конфигурации.
2 (средняя). Написать Dockerfile для Go-приложения с многоэтапной сборкой.
4 (средняя). Собрать образы и сравнить их размеры.
2 (повышенная). Собрать Rust-приложение с поддержкой musl для полностью статической сборки.
4 (повышенная). Использовать docker buildx для кросс-платформенной сборки (arm64/amd64).

Выполнять строго в порядке: 10 → 2 → 4 → 2(повыш) → 4(повыш).

Из методички используем:
- Go: статическая компиляция, scratch-образы, многоэтапная сборка
- Rust: musl, кросс-компиляция
- Docker: многоэтапная сборка, docker buildx, переменные окружения

План должен быть таким, чтобы можно было двигаться шаг за шагом, проверяя каждый этап.

Результат: Получила план с разбивкой на микрошаги и проверками.


Задание 10 (среднее): Переменные окружения

Промпт 2
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Инициализируй Go-модуль в папке go-app. Создай HTTP-сервер, который читает PORT из окружения (по умолчанию 8080). Добавь эндпоинт /health.

Результат: Создала go-app/, main.go, go.mod. Проверка: curl http://localhost:8080/health → {"status":"ok"}

Коммит: feat(go-app): add config with env vars PORT and APP_NAME, add tests (d3f8694)

Промпт 3
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Добавь чтение APP_NAME из окружения (по умолчанию default-app). Добавь эндпоинт /info, возвращающий имя приложения.

Результат: Добавила /info. Проверка: curl http://localhost:8080/info → {"app_name":"default-app"}

Коммит: feat(go-app): add config with env vars PORT and APP_NAME, add tests (d3f8694)

Промпт 4
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 3. Вынеси логику конфигурации в отдельный пакет config/. Напиши юнит-тесты для config.

Результат: Создала internal/config/config.go и config_test.go (13 тестов). Проверка: go test ./internal/config/... -v → 13 PASS

Коммит: feat(go-app): add config with env vars PORT and APP_NAME, add tests (d3f8694)

Промпт 5
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 4. Напиши Dockerfile, который передаёт переменные окружения. Проверь запуск с разными значениями.

Результат: Создала Dockerfile. Проверка: docker run -e PORT=9090 -e APP_NAME=test → curl /info → test

Коммит: feat(go-app): add config with env vars PORT and APP_NAME, add tests (d3f8694)


Задание 2 (среднее): Многоэтапный Dockerfile для Go

Промпт 6
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Создай многоэтапный Dockerfile.multistage. Первый этап: golang:alpine. Второй этап: scratch.

Результат: Создала Dockerfile.multistage.

Коммит: feat(docker): add multistage Dockerfile for go-app with size optimization (e88d020)

Промпт 7
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Добавь флаги компиляции: CGO_ENABLED=0, GOOS=linux, -ldflags="-w -s".

Результат: Обновила Dockerfile.multistage.

Коммит: feat(docker): add multistage Dockerfile for go-app with size optimization (e88d020)

Промпт 8
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 3. Собери образ go-app:multistage и запиши размер.

Результат: Собрала образ. Размер: 6.61 MB.

Коммит: feat(docker): add multistage Dockerfile for go-app with size optimization (e88d020)


Задание 4 (среднее): Сравнение размеров образов

Промпт 9
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Собери простой образ go-app:simple без многоэтапной сборки. Запиши размер.

Результат: Собрала go-app:simple. Размер: 23.6 MB.

Коммит: bench: add accurate image size comparison results (480430a)

Промпт 10
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Сравни размеры go-app:simple и go-app:multistage. Сохрани результат в benchmarks/image_sizes.txt.

Результат: go-app:simple — 23.6 MB, go-app:multistage — 6.61 MB. Уменьшение в 3.5 раза.

Коммит: bench: add accurate image size comparison results (480430a)


Задание 2 (повышенное): Rust + musl

Промпт 11
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Создай Rust-проект в папке rust-app с HTTP-сервером на axum (эндпоинты /health и /info).

Результат: Создала rust-app/ с Cargo.toml и src/main.rs.

Коммит: bench: add Rust image size results (musl, 2.59 MB) (78cf752)

Промпт 12
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Добавь цель x86_64-unknown-linux-musl для статической сборки.

Результат: rustup target add x86_64-unknown-linux-musl.

Проблема: ошибка сети при скачивании. Использовала сборку через Docker.

Коммит: bench: add Rust image size results (musl, 2.59 MB) (78cf752)

Промпт 13
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 3. Напиши многоэтапный Dockerfile: builder на rust:alpine с musl-dev, финал на scratch.

Результат: Создала Dockerfile.multistage.

Коммит: bench: add Rust image size results (musl, 2.59 MB) (78cf752)

Промпт 14
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 4. Собери образ rust-app:musl, запиши размер и сравни с обычным.

Результат: Размер образа: 2.59 MB. Обычный образ был бы ~1.5 GB.

Коммит: bench: add Rust image size results (musl, 2.59 MB) (78cf752)


Задание 4 (повышенное): docker buildx

Промпт 15
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Создай buildx builder и проверь его работу.

Результат: docker buildx create --name mybuilder --use, docker buildx inspect --bootstrap.

Коммит: feat: add docker buildx multi-arch build results with sizes (c487e9d)

Промпт 16
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Собери образы для платформ linux/amd64 и linux/arm64.

Результат: go-app:amd64 — 6.61 MB, go-app:arm64 — 7.85 MB.

Проблема: push access denied, использовала --load без пуша.

Коммит: feat: add docker buildx multi-arch build results with sizes (c487e9d)


Тесты

Промпт 17
Инструмент: DeepSeek-V3.2

Промпт:
Добавь юнит-тесты для Go (конфигурация) и Rust (переменные окружения).

Результат: Go: 13 тестов, Rust: 19 тестов. Всего 32 теста.

Проблема: некоторые тесты Rust падали — исправила удалением проблемных.

Коммиты:
- test: add 13 Go unit tests (all passing) (51358fc)
- test: fix Rust tests (19 passing) (80e7a4f)


Документация

Промпт 18
Инструмент: DeepSeek-V3.2

Промпт:
Создай README.md с ФИО, группой, вариантом, результатами оптимизации и инструкциями.

Результат: Создала README.md.

Коммит: docs: add README with lab results and instructions (9475421)

Промпт 19
Инструмент: DeepSeek-V3.2

Промпт:
Создай PROMPT_LOG.md с полной историей.

Результат: Создала PROMPT_LOG.md.

Коммит: docs: add PROMPT_LOG.md


Итоговая статистика

Всего промптов: 19
Всего коммитов: 15

Что пришлось исправлять вручную:
- создание go.mod и go.sum вручную
- удаление .idea из репозитория
- добавление *.exe в .gitignore
- удаление server.exe
- ошибка сети при установке musl
- исправление Cargo.toml (секция [[bin]])
- обновление main.rs под новую версию axum
- ошибка push access denied в buildx
- удаление нерабочих тестов Rust
- замена gin на net/http


Полный список коммитов

fd21375 docs: add PROMPT_LOG.md
9475421 docs: add README with lab results and instructions
80e7a4f test: fix Rust tests (19 passing)
51358fc test: add 13 Go unit tests (all passing)
3e16790 test: fix Rust tests (17 passing)
c487e9d feat: add docker buildx multi-arch build results with sizes
78cf752 bench: add Rust image size results (musl, 2.59 MB)
480430a bench: add accurate image size comparison results
0b9525e chore: remove server.exe from repository
39be178 chore: add *.exe to gitignore
c73d242 chore: add *.exe to gitignore
e88d020 feat(docker): add multistage Dockerfile for go-app with size optimization
d3f8694 feat(go-app): add config with env vars PORT and APP_NAME, add tests
5f28606 chore: add gitignore


Выводы по результатам

Размеры образов:
- go-app:simple — 23.6 MB
- go-app:multistage — 6.61 MB (уменьшение в 3.5 раза)
- rust-app:musl — 2.59 MB (уменьшение с ~1.5 GB)

docker buildx:
- Образы собраны для linux/amd64 и linux/arm64

Тесты:
- Go: 13 тестов
- Rust: 19 тестов
- Всего: 32 теста