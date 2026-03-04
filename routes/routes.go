package routes

import (
	"tacna-events-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, eventController *controllers.EventController) {
	// User routes
	r.GET("/users", userController.GetUsers)

	// Event routes
	r.GET("/events", eventController.GetEvents)
	r.GET("/events/:id", eventController.GetEvent)
	r.POST("/events", eventController.CreateEvent)
	r.PUT("/events/:id", eventController.UpdateEvent)
	r.DELETE("/events/:id", eventController.DeleteEvent)
}


