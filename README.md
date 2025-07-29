# 🧩 mini-gRPC

<a href="https://ibb.co/vt7fFyZ">
  <img src="https://i.ibb.co/LGLBFjn/IMG-20250729-145240.png" alt="IMG-20250729-145240" border="0">
</a>

**mini-gRPC** is a lightweight gRPC-based backend service designed to manage **Product** and **Category** data.

---

## ✨ Features

- ✅ Full CRUD operations for **Product** and **Category**
- ⚙️ Native SQL queries using `sqlc`
- 🛢️ PostgreSQL database integration
- 🚦 Built-in request limiter using **LRU cache**
- 🌐 **CORS** support for cross-origin requests
- 🔌 gRPC APIs using modern `connectRPC`

---

## 🛠️ Tools Used

| Tool         | Description                                                              |
|--------------|---------------------------------------------------------------------------|
| **connectRPC** | Simple and fast gRPC implementation for Go                              |
| **buf**         | Modern tooling for managing Protobuf files                             |
| **sqlc**        | Generates type-safe Go code from raw SQL queries                       |
| **protoc**      | Protocol Buffers compiler for generating gRPC server/client interfaces |
| **PostgreSQL**  | Relational database used as primary data store

---

## 🚀 Installation

### Prerequisites

Make sure the following are installed:

- [Go](https://go.dev/dl/) ≥ 1.23
- [buf](https://docs.buf.build/installation)
- [protoc](https://grpc.io/docs/protoc-installation/)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [PostgreSQL](https://www.postgresql.org/download/) ≥ 10

### 1. Clone the repository

```bash
git clone https://github.com/Kalveir/minigRPC.git && \
cd minigRPC
```

### 2. Install Go dependencies
```bash
go mod tidy
```
### 3.Generate Protobuf files
```bash
buf generate
```
### 4.Generate Go code from SQL queries
```bash
sqlc generate
```
### 5. Set Local Environtment
create ``.env`` file : 
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=<yourpassword>
DB_NAME=<your_db>

PORT_SERVICE=6363
RATE=100
BURST=30
LRU=200
```
## ⚙️ Development
To start the development server:
```bash 
go run cmd/main.go
```
## 📦 Deployment
### 1. Build the binary
```bash
make build
```
### 2. Set Production Environtment
```bash
mv env.prod .env
```
### 3. Run Docker Compose
```bash
docker compose up -d
```
### 4. Check Docker logs
```bash
docker logs myrpc
```

