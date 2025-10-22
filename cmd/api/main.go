package main

import (
	"database/sql"
	"event-api-app/internal/database"
	"event-api-app/internal/env"
	"log"

	_ "event-api-app/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// @title Go Gin Rest API
// @version 1.0
// @description This is a sample server for a Go Gin Rest API application.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email brookhave.dev@gmail.com
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	user := env.GetEnvString("POSTGRES_USER", "postgres")
	password := env.GetEnvString("POSTGRES_PASSWORD", "your_password")
	dsn := env.GetEnvString("POSTGRES_DSN", "postgres://"+user+":"+password+"@localhost:5432/eventdb?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "mysecret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
