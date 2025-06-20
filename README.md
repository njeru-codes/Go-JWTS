# 🔐 Go JWT Authentication API

A lightweight and secure REST API built with **Go**, **Fiber**, and **GORM**, featuring JWT-based user authentication and PostgreSQL as the database.

---

## ✨ Features

- ✅ User Registration (with hashed passwords using bcrypt)
- ✅ User Login (with JWT token generation)
- ✅ JWT-protected `/profile` endpoint
- ✅ GORM-based PostgreSQL integration
- ✅ Modular route/middleware structure
- ✅ Docker-ready architecture (optional)

---

## 🧱 Tech Stack

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://github.com/golang-jwt/jwt)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

## ⚙️ Getting Started

1.  🔐 Environment Setup

Create a `.env` file in the project root:

```env
DATABASE_DSN=host=localhost user=postgres password=yourpassword dbname=jwt_demo port=5432 sslmode=disable
JWT_SECRET=your-secret-key
```

2. Install Dependencies
```
go mod tidy
```

3. 🧪 Run the App
```
go run main.go
```
The API will run at: http://localhost:3000

📌 API Endpoints
| Method | Endpoint                | Auth | Description               |
| ------ | ----------------------- | ---- | ------------------------- |
| POST   | `/api/v1/auth/register` | ❌    | Register a new user       |
| POST   | `/api/v1/auth/login`    | ❌    | Login and receive a token |
| GET    | `/api/v1/auth/profile`  | ✅    | Get user profile          |

## 🔐 JWT Authentication

For protected routes, pass the token in the Authorization header:
```
Authorization: Bearer <your_token>
```

## PROJECT STRUCTURE
```
├── config/         # Env and config loader
├── database/       # DB connection and migration
├── dto/            # Request/response structs
├── handlers/       # Business logic for routes
├── middleware/     # JWT auth middleware
├── models/         # GORM models
├── routes/         # Route registration
├── utils/          # Helpers (JWT, password hashing)
├── main.go         # Entry point
└── go.mod          # Go module file

```
## ✍️ Author

Built by @njeru-codes

