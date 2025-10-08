# Go Restful API Template

A RESTful API built with Go for event management.

## Features

- 🎯 Event CRUD operations
- 👤 User authentication with JWT
- 🗄️ SQLite database with migrations
- 🔄 Hot reloading with Air
- 📊 Clean architecture

## Quick Start

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Run migrations: `go run cmd/migrate/main.go`
4. Start the server: `air` or `go run cmd/api/main.go`

## API Endpoints

- `GET /api/v1/events` - Get all events
- `POST /api/v1/events` - Create new event
- `GET /api/v1/events/:id` - Get event by ID
- `PUT /api/v1/events/:id` - Update event
- `DELETE /api/v1/events/:id` - Delete event

## Tech Stack

- **Language**: Go
- **Framework**: Gin
- **Database**: SQLite
- **Authentication**: JWT
- **Hot Reload**: Air
