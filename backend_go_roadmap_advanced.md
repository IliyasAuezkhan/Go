# 🐹 Роудмап 2.0: Go-бэкенд — уровень Middle

> **Уровень:** после базового роудмапа · **Длительность:** 15 дней · **Язык:** Go (Golang)
> Продолжение [backend_go_roadmap.md](.) — фокус на архитектуре, продакшене и распределённых системах.

---

## 📌 Легенда

| Тег | Описание |
|-----|----------|
| 📘 теория | Концепции и паттерны |
| 🛠 практика | Код, упражнения |
| 🚀 проект | Самостоятельная разработка |

---

## Неделя 1 — архитектура и надёжность

### День 1 — Чистая архитектура
> 📘 теория

- Слоистая архитектура: `handler → service → repository`
- Dependency Injection без магии (руками, через интерфейсы)
- Разделение на пакеты: `internal/`, `pkg/`, `cmd/`
- Domain-driven подход: сущности, DTO, мапперы

---

### День 2 — Конфигурация и структурные логи
> 🛠 практика

- `viper` или `envconfig` для конфигов (env / yaml)
- `slog` по-серьёзному: structured logging, контекст, уровни
- Паттерн graceful shutdown (`context.Context` + `signal.NotifyContext`)
- `errors.Is`/`errors.As`, обёртывание ошибок (`fmt.Errorf("%w")`)

---

### День 3 — Контекст и конкурентность
> 📘 теория

- `context.Context`: отмена, таймауты, значения (и почему их лучше не класть)
- Горутины и каналы: паттерны `worker pool`, `fan-in/fan-out`
- `sync.WaitGroup`, `sync.Mutex`, `errgroup`
- Гонки данных: `go run -race`

---

### День 4 — PostgreSQL глубже
> 🛠 практика

- Индексы: B-tree, составные, `EXPLAIN ANALYZE`
- Миграции через `golang-migrate` (вместо `AutoMigrate`)
- Пул соединений `pgxpool`, настройка `max_conns`
- N+1 проблема и её решения (батчинг, `sqlc` вместо ORM)

---

### День 5 — Redis и кэширование
> 📘 теория

- Паттерны кэша: cache-aside, write-through
- Redis как хранилище сессий и rate-limiter
- TTL, инвалидация кэша — «две главные проблемы информатики»
- Distributed lock через Redis (`SETNX`)

---

### День 6 — Асинхронность через очереди
> 🛠 практика

- Kafka: продюсеры/консьюмеры на Go (`segmentio/kafka-go` или `confluent-kafka-go`)
- Идемпотентность обработки сообщений
- Consumer groups, партиции, offset-менеджмент
- Паттерн Outbox для надёжной публикации событий

*(раз ты уже работаешь с Kafka UI и топиком `RECEIPTS_EXPORT` — тут будет полезно закрепить теорию под практику)*

---

### День 7 — Мини-проект: событийный сервис
> 🚀 проект

Сервис обработки заказов:

- HTTP API создаёт заказ → пишет в Postgres
- Публикует событие в Kafka (`order.created`)
- Отдельный консьюмер обрабатывает событие асинхронно
- Redis кэширует статус заказа для быстрого чтения

**Стек:** Gin + PostgreSQL + Kafka + Redis

---

## Неделя 2 — продакшен и масштаб

### День 8 — Аутентификация через OAuth2/OIDC
> 📘 теория

- OAuth2 flows: authorization code, client credentials
- OIDC поверх OAuth2: id_token, JWKS, валидация подписи
- Интеграция с Keycloak как Identity Provider из Go-сервиса
- Middleware для проверки токена и ролей (RBAC)

*(отлично ляжет на твой опыт с Keycloak на `logistic-admin.skiftrade.kz`)*

---

### День 9 — gRPC
> 🛠 практика

- Protocol Buffers: `.proto`, генерация кода (`protoc`)
- Unary и streaming RPC
- gRPC vs REST: когда что выбирать
- Interceptors (аналог middleware в gRPC)

---

### День 10 — Наблюдаемость (observability)
> 📘 теория

- Метрики: Prometheus + `client_golang`, `/metrics` эндпоинт
- Дашборды в Grafana
- Трейсинг: OpenTelemetry, распределённый трейс между сервисами
- Health checks: `/healthz`, `/readyz`

---

### День 11 — Тестирование по-взрослому
> 🛠 практика

- Table-driven тесты, `testify`
- Интеграционные тесты с `testcontainers-go` (реальный Postgres/Kafka в Docker)
- Моки через `mockery` или `gomock`
- Контрактное тестирование API

---

### День 12 — CI/CD
> 📘 теория

- GitHub Actions: lint → test → build → deploy
- `golangci-lint`, статический анализ
- Сборка multi-arch Docker-образов
- Секреты и переменные окружения в CI

---

### День 13 — Kubernetes: основы для бэкендера
> 🛠 практика

- Deployment, Service, ConfigMap, Secret — минимум для запуска
- Readiness/liveness probes (связь с Днём 10)
- Horizontal Pod Autoscaler — идея, не глубина
- `kubectl` базовые команды, локальный кластер (`kind`/`minikube`)

---

### День 14 — Итоговый проект
> 🚀 проект

Продакшен-готовый сервис:

- Чистая архитектура (слои + DI)
- OIDC-аутентификация через Keycloak
- gRPC + REST API одновременно
- Kafka для событий, Redis для кэша
- Метрики + трейсинг
- Docker Compose → манифесты для Kubernetes
- CI-пайплайн с тестами и линтом

---

## День 15 — Ретроспектива и специализация

> 🚀 проект

- Выбор направления для углубления: высоконагруженные системы / платформенная инженерия / DevOps
- Профилирование (`pprof`) — поиск узких мест в своём проекте
- Код-ревью своего итогового проекта (можно попросить ИИ или коллегу)
- План на следующий месяц: конкретный pet-проект или вклад в open source

---

## 📚 Полезные ресурсы

| Ресурс | Ссылка |
|--------|--------|
| Effective Go | https://go.dev/doc/effective_go |
| Go Concurrency Patterns | https://go.dev/blog/pipelines |
| sqlc | https://sqlc.dev |
| golang-migrate | https://github.com/golang-migrate/migrate |
| Kafka Go client | https://github.com/segmentio/kafka-go |
| OpenTelemetry Go | https://opentelemetry.io/docs/languages/go |
| testcontainers-go | https://golang.testcontainers.org |
| Kubernetes basics | https://kubernetes.io/docs/concepts |

---

*Темп — 1–2 часа в день, но День 6, 9 и 14 потребуют больше времени на практику.*
