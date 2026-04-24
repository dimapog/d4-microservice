# D4 Microservice

A Go microservice built with Gin, GORM, SQLite, and OpenAI integration.
Swagger documentation: github.com/swaggo/swag/cmd/swag@latest

This project includes modular components for user management, authentication, AI-powered calculations, and high-load asynchronous CSV ingest.
This project is a placeholder for real-life chalanges.

## Project Structure

- `main.go` — application bootstrap and route registration.
- `internal/ai` — AI prompt handling and OpenAI response integration.
- `internal/auth` — JWT login service and auth routes.
- `internal/calculator` — user-specific BMI and heart rate zone calculations.
- `internal/csv` — asynchronous CSV upload and batch import into the `clients` table.
- `internal/user` — user repository, service, and API routes.
- `middleware` — authentication middleware for protected endpoints.
- `utils` — environment loading and database connection setup.

## Modules

### `internal/user`

Handles user storage and account operations.

Routes:
- `POST /user/` — create new user
- `GET /user/:id` — get user by ID (protected)
- `PATCH /user/` — update authenticated user (protected)
- `DELETE /user/:id` — delete user

### `internal/auth`

Routes:
- `POST /auth/login` — login and receive JWT

### `internal/calculator`

Placeholder functionality with descrbed business logic.
Provides authenticated user calculators.

Routes:
- `GET /calculator/bmi` — calculate BMI for authenticated user
- `GET /calculator/hrz` — calculate heart rate zones for authenticated user

### `internal/ai`

Integrates with OpenAI for enhanced personal health calculations.

Routes:
- `POST /ai/personal-calculation` — send health data and receive AI-generated analysis

### `internal/csv`

The CSV module is designed to process large CSV files asynchronously using channels and batch inserts.
Supports high-load async CSV import into a dedicated `clients` table.

Routes:
- `POST /csv/upload` — upload CSV file for async processing

## Middleware

- `middleware/auth.go` — bearer token JWT verification and request context user injection.

## Utils

- `utils/loadEnvVariables.go` — loads `.env` file via `godotenv`.
- `utils/dbConnect.go` — connects to SQLite using `DB_NAME` env.
- `utils/dbSync.go` — placeholder for DB sync logic.

## Environment

The service expects a `.env` file with at least:

```text
DB_NAME=database.db
SECRET_KEY=your-secret-key
OPENAI_API_KEY=your-openai-api-key
```

## Running

```bash
go run main.go
```

## Notes

- Module-specific migrations are now handled inside each module via `init()`.


## Swagger documentation

swag init -g main.go -o docs