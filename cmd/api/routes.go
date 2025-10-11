package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		// Event routes
		v1.POST("/events", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)
		// Attendee routes
		v1.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)
		v1.GET("/events/:id/attendees", app.getAttendeesForEvent)
		v1.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent)
		v1.GET("/attendees/:userId/events", app.getEventsByAttendee)

		// User routes
		v1.POST("/auth/register", app.registerUser)

		// Authentication routes
		v1.POST("/auth/login", app.login)
	}

	return g
}
