# 📋 Event Management REST API

A complete, production-ready REST API built with Go and Gin framework, featuring comprehensive authentication, authorization, and interactive API documentation.

## ✨ Key Features

🔐 **Authentication & Security**
- JWT-based authentication system
- Secure user registration and login
- Bearer token authorization
- Ownership-based access control

📚 **Interactive Documentation** 
- Complete Swagger/OpenAPI 3.0 documentation
- Interactive API testing interface
- Comprehensive endpoint documentation with examples

🎯 **Event Management**
- Full CRUD operations for events
- User ownership validation
- Attendee management system
- RESTful API design principles

🛠️ **Technical Stack**
- **Backend**: Go 1.24+ with Gin framework
- **Database**: SQLite with migration support
- **Authentication**: JWT tokens with HMAC-SHA256
- **Documentation**: Swagger/OpenAPI with Gin-Swagger
- **Development**: Air for hot reloading

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/brookhaven13/golang-event-api

# Install dependencies
go mod tidy

# Run the application
go run cmd/api/*.go

# Visit Swagger UI
http://localhost:8080/swagger/
```

## 📖 API Endpoints

### Public Endpoints
- `GET /api/v1/events` - List all events
- `POST /auth/register` - User registration
- `POST /auth/login` - User authentication
- `GET /events/{id}/attendees` - Get attendees for event
- `GET /users/{userId}/events` - Get events by attendee

### Protected Endpoints (Requires JWT)
- `POST /events` - Create new event
- `PUT /events/{id}` - Update event (owner and admin only)
- `DELETE /events/{id}` - Delete event (owner and admin only)
- `POST /events/{id}/attendees/{userId}` - Add attendee
- `DELETE /events/{id}/attendees/{userId}` - Remove attendee
- `PUT /auth/user` - Update user information (email, name, password)
- `DELETE /events/{id}/attendees/{userId}` - Remove attendee from event (admin or self)

## 🔧 Environment Configuration

```env
PORT=8080
JWT_SECRET=your_secure_jwt_secret
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DSN=postgres://your_username:your_password@localhost:5432/eventdb?sslmode=disable
```

## 🗄️ PostgreSQL 設定教學

1. 安裝 PostgreSQL 並啟動服務。
2. 建立資料庫：
	```sh
	createdb -U postgres -h localhost eventdb
	```
3. 建立/設定使用者與密碼（如需自訂）：
	```sh
	psql -U postgres -h localhost
	CREATE USER your_user WITH PASSWORD 'your_password';
	GRANT ALL PRIVILEGES ON DATABASE eventdb TO your_user;
	```
4. 在 `.env` 檔案設定：
	```env
	POSTGRES_USER=your_user
	POSTGRES_PASSWORD=your_password
	POSTGRES_DSN=postgres://your_user:your_password@localhost:5432/eventdb?sslmode=disable
	```

請勿將 `.env` 檔案加入版本控制，確保帳號密碼安全。

## 📈 Perfect for Learning

- Modern Go backend patterns
- RESTful API design
- JWT authentication implementation
- Database operations with migrations
- API documentation best practices
- Docker containerization ready

## 🏷️ GitHub Topics

```
golang, gin, rest-api, jwt, swagger, authentication, backend, api, go-gin, openapi, sqlite, migration, docker-ready, production-ready
```

## 💬 One-liner Description

```
Complete Go REST API with JWT auth, Swagger docs & event management - perfect for learning modern backend development! 🚀
```
