# Basic CRUD Backend using gRPC, Go, and MongoDB

## 🚀 Project Overview

This project implements a **basic CRUD (Create, Read, Update, Delete) backend API** using **gRPC** in **Go** and **MongoDB** as the database. The backend provides fast, efficient, and type-safe communication between clients and the server, ideal for microservices architectures.

---

## 🎯 Features

* ✅ Create, Read, Update, and Delete operations
* ✅ gRPC for communication
* ✅ MongoDB as the persistent data store
* ✅ Clean and modular project structure
* ✅ Protobuf definitions for request/response schemas

---

## 📚 Tech Stack

* **Go** – Programming language
* **gRPC** – RPC framework
* **Protocol Buffers (Protobuf)** – Message serialization
* **MongoDB** – NoSQL database

---

## ⚡️ Prerequisites

* [Go](https://golang.org/dl/) installed (version >= 1.18)
* [MongoDB](https://www.mongodb.com/try/download/community)
* [protoc](https://grpc.io/docs/protoc-installation/)
* [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
* [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)

---

## 🏗️ Project Structure

```
.
├── proto/
│   └── user.proto          # Protobuf service & message definitions
├── server/
│   └── main.go             # gRPC server implementation
├── database/
│   └── mongo.go            # MongoDB connection logic
├── models/
│   └── user.go             # Data models
├── services/
│   └── user_service.go     # CRUD logic implementations
├── go.mod
└── README.md
```

---

## ⚙️ Setup Instructions

### 1️⃣ Clone the repository

```bash
git clone https://github.com/yourusername/grpc-go-mongodb-crud.git
cd grpc-go-mongodb-crud
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Generate Go code from Protobuf

```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

### 4️⃣ Configure MongoDB

Ensure MongoDB is running locally or remotely, and update the connection URI in `database/mongo.go`.

### 5️⃣ Run the gRPC server

```bash
go run server/main.go
```

---

## 📡 Example RPC Methods

* `CreateUser(User) returns (User)`
* `GetUser(UserID) returns (User)`
* `UpdateUser(User) returns (User)`
* `DeleteUser(UserID) returns (Empty)`

---

## ✅ Testing the API

You can use [grpcurl](https://github.com/fullstorydev/grpcurl) or a custom gRPC client to test the endpoints.

Example:

```bash
grpcurl -plaintext -d '{"id":"123"}' localhost:50051 your.package.UserService/GetUser
```

---

## 📚 References

* [gRPC Official Docs](https://grpc.io/docs/)
* [MongoDB Go Driver](https://go.mongodb.org/mongo-driver)

---

## 💡 Future Improvements

* Add authentication & authorization
* Implement pagination for list requests
* Add unit tests

---

## ⚡ License

This project is licensed under the MIT License.
