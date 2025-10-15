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
git clone https://github.com/brookhaven13/golang-api-template

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

### Protected Endpoints (Requires JWT)
- `POST /events` - Create new event
- `PUT /events/{id}` - Update event (owner only)
- `DELETE /events/{id}` - Delete event (owner only)
- `POST /events/{id}/attendees/{userId}` - Add attendee
- `DELETE /events/{id}/attendees/{userId}` - Remove attendee

## 🔧 Environment Configuration

```bash
PORT=8080
JWT_SECRET=your_secure_jwt_secret
```

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