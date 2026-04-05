# Task Management Microservice

A high-performance task management service built with **Golang**, utilizing **gRPC** for low-latency communication and **gRPC-Gateway** for RESTful compatibility.

## Features
* **3-Layered Architecture:** Clear separation between Transport, Service, and Repository layers.
* **Dual Protocol Support:** gRPC (Port 50051) for internal service calls and REST (Port 8080) for web clients.
* **Persistent Storage:** PostgreSQL integration via GORM.
* **Automated Migrations:** Database schema is automatically updated on startup.
* **Containerized Environment:** Easy setup using Docker Compose.

---

## Tech Stack
* **Language:** Go 1.24+
* **Communication:** gRPC, Protocol Buffers v3
* **API Gateway:** gRPC-Gateway v2
* **Database:** PostgreSQL 15 (Alpine)
* **ORM:** GORM

---

## Getting Started

### 1. Prerequisites
* Docker & Docker Compose
* Go 1.24+
* `protoc` compiler (if modifying `.proto` files)

### 2. Infrastructure Setup
Spin up the PostgreSQL database:
```bash
docker-compose up -d
```

### 3. Generate Protobuf Code
If you have modified the `task.proto` file, regenerate the Go code:
```bash
make gen
```

### 4. Run the Service
```bash
go run cmd/server/main.go
```

---

## Testing the API

### REST Interface (JSON)
```bash
curl -X POST http://localhost:8080/v1/tasks \
     -H "Content-Type: application/json" \
     -d '{"title": "Complete README", "description": "Write documentation for the project"}'
```

### gRPC Interface
Using `grpcurl`:
```bash
grpcurl -plaintext -d '{"title": "gRPC Task", "description": "Testing binary speed"}' \
    localhost:50051 task.TaskService/CreateTask
```

---

## 📝 Database Management
You can inspect the saved tasks directly in the container:
```bash
docker exec -it task_db psql -U user -d task_management -c "SELECT * FROM tasks;"
```
