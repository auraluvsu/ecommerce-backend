# 🛒 GoCommerce

A simple and scalable e-commerce backend built with Go (Golang).  
This project is part of my backend development learning journey, focused on building real-world, production-ready systems using clean architecture principles and Go’s powerful standard library.

## ✨ Features

- 📦 Product Management (CRUD)
- 🛍️ Shopping Cart System
- 👤 User Authentication (OAuth planned)
- 💳 Order Processing (basic)
- 🧠 In-memory caching (for performance)
- 🔍 RESTful API Design
- 🛠️ SQLite or Prisma integration (TBD)

## 🧱 Tech Stack

- **Language:** Go
- **Database:** SQLite (switchable)
- **Authentication:** OAuth (coming soon)
- **API Style:** REST
- **Monitoring (planned):** Prometheus + Grafana
- **Testing:** Built-in Go testing framework

## 🏗️ Project Structure

```bash
gocommerce/
├── cmd/                # Entry point of the application
├── internal/           # Business logic
│   ├── products/       # Product domain logic
│   ├── users/          # User domain logic
│   └── orders/         # Order domain logic
├── pkg/                # Shared utilities
├── api/                # HTTP handlers and routing
├── db/                 # Database access and models
└── main.go             # Application bootstrapper
