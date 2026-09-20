# Архитектура SelfDev

SelfDev — трекер привычек с Vue SPA, Go HTTP API, PostgreSQL, Redis и внешним OAuth/OIDC-провайдером Casdoor. Компоненты запускаются через Docker Compose и публикуются Traefik по разным хостам.

## Карта репозитория

- `frontend/` — Vue 3 + TypeScript SPA. Сборка выполняется Vite, production-статику раздаёт nginx.
  - `src/main.ts` — точка входа: создаёт приложение, подключает router и toast-уведомления.
  - `src/router/index.ts` — маршруты `/` (landing) и `/me/profile` (dashboard).
  - `src/views/` — страницы. `Dashboard.vue` также содержит клиентское состояние, HTTP-вызовы API и расчёт статистики/heatmap; это главный файл для изменений поведения dashboard.
  - `src/components/` — презентационные компоненты dashboard и двух header-вариантов.
  - `src/config/env.ts` — доступ к build-time переменным Vite.
  - `src/composables/useTheme.ts` — общее состояние светлой/тёмной темы.
  - `.env.development`, `.env.stage`, `.env.prod` — значения `VITE_*` для режимов сборки.
  - `public/`, `src/assets/` — статические ресурсы и глобальные стили.
- `backend/` — Go-модуль HTTP API.
  - `cmd/api/main.go` — исполняемая точка входа и composition root: создаёт конфигурацию, logger, подключения, repository/service/handler и HTTP server.
  - `internal/handler/` — маршрутизация, разбор HTTP-запросов и DTO.
  - `internal/service/` — сценарии аутентификации, пользователей и привычек.
  - `internal/repository/` — SQL-запросы к PostgreSQL и сессии в Redis.
  - `internal/model/` — доменные структуры `User`, `Habit` и отметки выполнения.
  - `internal/integrations/casdoor/` — адаптер OAuth token exchange и загрузки userinfo из Casdoor.
  - `internal/config/` — загрузка переменных окружения.
  - `database/` — создание клиентов PostgreSQL и Redis.
  - `middleware/` — CORS и проверка cookie-сессии.
  - `migrations/` — goose-миграции схемы приложения.
  - `logger/` — настройка zap для development/production.
- `infra/` — контейнеризация и окружение.
  - `docker-compose.dev.yaml`, `docker-compose.prod.yaml` — Traefik, frontend, backend, Casdoor, PostgreSQL, Redis и служебный контейнер миграций.
  - `Dockerfile.backend`, `Dockerfile.frontend` — сборка API и SPA.
  - `nginx.conf` — раздача SPA с fallback на `index.html`.
  - `auth/conf/` — конфигурация и начальные данные Casdoor.
  - `postgres/init.sql` — создаёт отдельную БД `casdoor`; таблицы приложения создают миграции.
  - `.env.example` — шаблон runtime/Compose/goose-переменных.
- `Taskfile.yaml` — команды запуска, миграций, форматирования и логов.

## Точки входа и поток зависимостей

Frontend начинается в `frontend/index.html` и `frontend/src/main.ts`; `App.vue` выводит текущий `RouterView`. Основные пользовательские сценарии сосредоточены в `Landing.vue` (переход в Casdoor) и `Dashboard.vue` (CRUD и отметки привычек через API).

Backend начинается в `backend/cmd/api/main.go` и собирается в одном направлении:

```text
HTTP request
  -> CORS / session middleware
  -> handler + DTO
  -> service
  -> repository / Casdoor adapter
  -> PostgreSQL, Redis или Casdoor
```

Связи задаются небольшими интерфейсами рядом с потребителем, а конкретные реализации связываются вручную в `main.go`. Поэтому для изменения HTTP-контракта нужно начинать с `internal/handler/<domain>/` и frontend `Dashboard.vue`; для бизнес-сценария — с `internal/service/<domain>/`; для хранения — с `internal/repository/<domain>/` и `migrations/`.

Основные домены:

- `habit`: полный CRUD привычек и создание/удаление/чтение дат выполнения; цепочка `handler/habit` → `service/habit` → `repository/habit` → таблицы `habits`, `habits_completed`.
- `auth`: callback от Casdoor, получение профиля, поиск или создание локального пользователя и выпуск случайной session ID; цепочка `handler/auth` → `service/auth` → `integrations/casdoor`, `service/user`, `repository/auth`.
- `user`: модель и persistence пользователя используются аутентификацией; публичные user endpoints пока не зарегистрированы.

## HTTP и аутентификация

Маршруты регистрирует `backend/internal/handler/router.go` на стандартном `net/http.ServeMux`:

- публичный `GET /auth/callback` завершает OAuth flow;
- защищённые `/api/*` обслуживают список, создание, изменение и удаление привычек, а также `/api/habit/{id}/confirm` для отметок выполнения.

`main.go` помещает весь `/api/` subtree за `middleware/session.go`. Middleware читает cookie `session`, получает из Redis `user_id` и кладёт UUID в context запроса. Habit handler извлекает его из context, преобразует JSON через `internal/handler/habit/dto/` и вызывает service. CORS middleware оборачивает оба набора маршрутов и разрешает credentials для origin из `MIDDLEWARE`.

Frontend вызывает `${VITE_API_HOST}/api/...` с `credentials: 'include'`. Login/signup перенаправляют браузер на `${VITE_AUTH_HOST}/...`; Casdoor возвращает authorization code на backend callback, который устанавливает cookie и перенаправляет пользователя на `${REDIRECT_URI}/me/profile`.

## Данные и миграции

`database/postgres.go` создаёт и проверяет `pgxpool.Pool`; SQL написан непосредственно в repository. PostgreSQL хранит пользователей, привычки и историю их выполнения. `database/redis.go` создаёт go-redis client; `repository/auth` хранит ключи `session:<id>` с UUID пользователя и TTL 30 дней.

Схема последовательно описана в `backend/migrations/*.sql`. Новую миграцию следует создавать рядом с ними командой `task migration NAME=<name>`. Применение и откат: `task migrate`, `task migrate-one`, `task rollback-one` или `task rollback`; команды ожидают локальный `infra/.env` на основе `.env.example` и установленный `goose`.

## Конфигурация и запуск

Backend `internal/config/config.go` при старте обязательно загружает `.env` через `godotenv`, затем читает `POSTGRES_URL`, `REDIS_URL`, `AUTH_CLIENT_ID`, `AUTH_CLIENT_SECRET`, `REDIRECT_URI`, `COOKIE_DOMAIN`, `COOKIE_SECURE`, `MIDDLEWARE`, `ENV` и HTTP timeouts. Сервер фактически слушает `:8080`; поля `HOST` и `ADDR` сейчас не используются. В development пустой `COOKIE_DOMAIN` создаёт host-only cookie, а в production следует задать домен и `COOKIE_SECURE=true`. В Compose `infra/.env` монтируется как `/app/.env`.

Frontend получает конфигурацию при сборке из `VITE_AUTH_HOST`, `VITE_API_HOST`, `VITE_REDIRECT_URI`, `VITE_CASDOOR_CLIENT_ID`; `MODE` из Compose выбирает соответствующий Vite env-файл. Для development используются автоматически резолвящиеся хосты `*.self-dev.localhost`, поэтому ручная правка `/etc/hosts` не нужна. `task dev-up` поднимает локальный стек и применяет миграции, `task prod-up` — production Compose, `task dev-frontend` — только Vite dev server.

Development bootstrap Casdoor выполняется сервисом `auth-bootstrap`: он фиксирует OAuth client ID/secret и callback в локальной базе. Для production значения auth и cookie должны быть заменены на собственные.

## Тесты и проверки

Автоматических тестовых файлов в репозитории сейчас нет; frontend test-runner и npm-скрипт `test` также не настроены. При добавлении Go-тесты следует размещать рядом с пакетами как `*_test.go`, frontend-тесты — рядом с компонентами/модулями или в отдельном `frontend/src/**/__tests__/` после выбора test-runner.

Доступные проверки:

```sh
cd backend && go test ./...
cd frontend && npm ci                 # один раз установить зависимости
cd frontend && npm run type-check
cd frontend && npm run build
```

`task lint` форматирует и автоматически исправляет файлы (`go fmt`, `golangci-lint --fix`, Prettier), поэтому это не read-only проверка.
