# students-api

![Go Version](https://img.shields.io/badge/Go-1.25.4-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square)

A lightweight, production-oriented **REST API** built in **Go (Golang)** for managing student records. This project follows **clean architecture** principles with **dependency injection**, **SQLite** persistence, **structured logging** via **slog**, and **request validation** using **go-playground/validator**.

---

## Repository Description
students-api is a cleanly-structured Go REST service that exposes CRUD operations for student records using `net/http` (`ServeMux`) and a SQLite-backed storage layer.

---

## Suggested GitHub Topics
- go
- golang
- rest-api
- sqlite
- net-http
- clean-architecture
- dependency-injection
- slog
- validation
- students

---

## Project Structure

```text
students-api/
├─ cmd/
│  └─ students-api/
│     └─ main.go
├─ internal/
│  ├─ config/
│  │  └─ config.go
│  ├─ http/
│  │  └─ handlers/
│  │     └─ student/
│  │        └─ student.go
│  ├─ storage/
│  │  ├─ storage.go
│  │  └─ sqlite/
│  │     └─ sqlite.go
│  ├─ types/
│  │  └─ types.go
│  └─ utils/
│     └─ response/
│        └─ response.go
├─ go.mod
└─ go.sum
```

---

## Tech Stack
- **Go (Golang)**
- **SQLite**
- **net/http (ServeMux)**
- **go-playground/validator**
- **slog** (structured logging)

---

## Architecture (Clean Separation of Concerns)

**Handler → Storage Interface → SQLite Implementation**

- **HTTP layer (Handlers)**: request parsing/decoding, request validation, and response formatting  
- **Storage interface**: abstracts persistence operations from HTTP concerns  
- **SQLite implementation**: provides the concrete repository using `modernc.org/sqlite`  
- **Dependency injection**: wiring happens at startup; handlers receive a storage implementation

---

## Installation & Setup

### 1) Prerequisites
- Go **1.25+**
- No external DB server required (SQLite file-based storage)

### 2) Clone
```bash
git clone https://github.com/Anandusanthosh123/students-api.git
cd students-api
```

### 3) Run
```bash
go run cmd/students-api/main.go -config config/local.yaml
```

### 4) Verify
The API exposes student CRUD endpoints under the `/api/students` base path.

---

## API Endpoints

Base path: `/api/students`

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/students` | Create a new student |
| `GET` | `/api/students` | Get all students |
| `GET` | `/api/students/{id}` | Get a student by ID |
| `PUT` | `/api/students/{id}` | Update student details |
| `DELETE` | `/api/students/{id}` | Delete a student by ID |

### Response behavior
- Input is validated using `go-playground/validator`
- JSON responses are returned with appropriate HTTP status codes

---

## Sample `curl` Request

### Create a student
```bash
curl -s -X POST "http://localhost:8080/api/students" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }'
```

Example response:
```json
{ "id": 1 }
```

---

## Design Principles

- **Clean Architecture**
  - Handlers do not depend on storage implementation details
  - Storage is accessed via an interface (`internal/storage/storage.go`)
- **Dependency Injection**
  - The startup wiring injects the SQLite repository into handlers
- **Structured Logging**
  - Uses `log/slog` for production-friendly, key/value observability
- **Request Validation**
  - Enforces constraints at the handler layer using `go-playground/validator`
  - Returns `400 Bad Request` on validation failures
- **SQLite Persistence**
  - Uses `modernc.org/sqlite` (no external database server required)

---

## Future Improvements

- **Graceful shutdown**
  - Add full `context`-based graceful termination and request draining
- **Consistent error envelope**
  - Standardize error responses across endpoints (code/message/details)
- **Pagination & filtering**
  - Support query parameters for paging, sorting, and searching
- **Database performance**
  - Add indexes (e.g., on `email`) and improve query efficiency
- **Request correlation**
  - Add request IDs and propagate them through structured logs
- **Testing**
  - Unit tests with storage mocks + integration tests for SQLite

---

## Owner

- **GitHub**: https://github.com/Anandusanthosh123  
- **Username**: Anandusanthosh123
