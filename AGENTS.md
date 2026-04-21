# AGENTS.md

## 🧠 Overview

This project is a modular backend API written in Go.

Architecture style:

* Feature-based (modular)
* Each module is self-contained
* No global "shared business logic" dumping

Each module contains:

* HTTP handler
* business logic (service)
* data access (repository)
* models / DTOs

Database:
* SQLite in root with name "database.db"
* Connection descibed in (utils/dbConnect.go)

.ENV variables
* are storred in /.env
* handler for .env variables: (utils/loadEnvVariables.go)



---

## 📁 Project Structure

main.go
/internal

```
/auth
    handler.go
    service.go
    repository.go
    model.go

/ai
    service.go
    client.go
```

---


## 🧩 Module Generation Rules

To create a new module, follow this convention:

### Command

"Create module <name>"

Example:
"Create module user"

---

### Expected Structure

/internal/<name>/
handler.go
service.go
repository.go
model.go

---

### File Responsibilities

handler.go:

* HTTP layer
* request parsing
* response handling

service.go:

* business logic
* orchestration

repository.go:

* database access

model.go:

* domain models

dto.go:

* request/response structures

---

### Naming Conventions

* Module name: lowercase (user, listing, order)
* Structs:
  UserService
  UserHandler
  UserRepository

---

### Constructor Pattern

Each module must expose:

NewService(...)
NewHandler(...)
NewRepository(...)

---

### Routing

Handler must expose:

func (h *Handler) RegisterRoutes(router Router)

---

### Constraints

* No cross-module imports
* No business logic in handler
* No DB access outside repository

---



## 📦 Module Rules

Each module MUST:

* be isolated
* not depend on other modules directly
* communicate via interfaces only (if needed)

Allowed inside module:

* handler (HTTP layer)
* service (business logic)
* repository (DB layer)
* model (domain models)
* dto (request/response)

---

## 🔹 Handler Rules

Handlers must:

* only handle HTTP concerns
* parse request
* validate input (basic)
* call service
* return response

Handlers MUST NOT:

* contain business logic
* access database directly

---

## 🔹 Service Rules

Services:

* contain ALL business logic
* orchestrate repositories
* may call external APIs
* may use other services via interfaces

---

## 🔹 Repository Rules

Repositories:

* handle DB operations only
* no business logic
* return domain models

---

## 🔌 Dependency Rules

* No circular dependencies
* Modules should not import each other directly
* Shared logic goes to `/utils` or separate service

---

## 🤖 AI Integration Rules

* AI logic must live in `/internal/ai`
* Other modules call AI via interface
* Never call OpenAI directly from handlers

---

## ⚙️ Coding Guidelines

* Use constructor injection (NewService, NewHandler)
* Use interfaces for testability
* Keep functions small and explicit
* Prefer composition over inheritance

---

## 🚫 Anti-patterns

* fat handlers ❌
* direct DB access in handler ❌
* cross-module imports ❌
* global state ❌

---

## 📝 Notes

* This is an MVP — keep it simple
* Do not over-engineer
* Refactor when duplication becomes obvious
