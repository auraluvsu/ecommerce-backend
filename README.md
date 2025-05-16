# 🛒 GoCommerce

A simple and scalable e-commerce backend built with Go (Golang).  
This project is part of my backend development learning journey, focused on building real-world, 
production-ready systems using clean architecture principles and Go’s powerful standard library.

## ✨ Features

- 📦 Product Management (CRUD)
- 🛍️ Shopping Cart System
- 👤 User Authentication (JWT planned)
- 🧠 In-memory caching (for performance)
- 🔍 RESTful API Design
- 🛠️ PostgreSQL integration

## 🧱 Tech Stack

- **Language:** Go
- **Database:** PostgreSQL (switchable)
- **Authentication:** JWT (coming soon)
- **API Style:** REST
- **Testing:** Built-in Go testing framework

## 🏗️ Project Structure

```bash
gocommerce/
├── cmd/                # Entry point of the application
├── internal/           # Business logic
│   ├── db/             # Database access and models
│   ├── handlers/       # User registry logic
│   ├── api/            # HTTP handlers and routing logic
│   ├── middleware/     # Authentication logic
│   └── models/         # Struct models
├── pkg/                # Shared Utilities
└── main.go             # Application bootstrapper
