package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	g := gin.Default()
	g.Use(CORSMiddleware())

	v1 := g.Group("/api/v1")
	{
		// Event routes
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)

		// Attendee routes
		v1.GET("/events/:id/attendees", app.getAttendeesForEvent)
		v1.GET("/attendees/:userId/events", app.getEventsByAttendee)

		// User routes
		v1.POST("/auth/register", app.registerUser)

		// Authentication routes
		v1.POST("/auth/login", app.login)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		// 受保護的路由

		// Event routes
		authGroup.POST("/events", RequireVerifiedUser(), app.createEvent)
		authGroup.PUT("/events/:id", RequireVerifiedUser(), app.updateEvent)
		authGroup.DELETE("/events/:id", RequireVerifiedUser(), app.deleteEvent)

		// Attendee routes
		authGroup.POST("/events/:id/attendees/:userId", RequireVerifiedUser(), app.addAttendeeToEvent)
		authGroup.DELETE("/events/:id/attendees/:userId", RequireVerifiedUser(), app.deleteAttendeeFromEvent)

		// User update route
		authGroup.PUT("/auth/user", app.updateUser)
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

	return g
}
