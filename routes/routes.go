package routes

import (
	"github.com/gin-gonic/gin"
	"tacna-events-backend/controllers"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, eventController *controllers.EventController) {
	r.GET("/users", userController.GetUsers)
	r.GET("/events", eventController.GetEvents)
	r.POST("/users/create", userController.CreateUser)
}
