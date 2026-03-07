package routes

import (
	"tacna-events-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, eventController *controllers.EventController, categoryController *controllers.CategoryController) {

	// User routes

	r.GET("/users", userController.GetUsers)

	// Event routes

	r.GET("/events", eventController.GetEvents)

	r.GET("/events/:id", eventController.GetEvent)

	r.POST("/events", eventController.CreateEvent)

	r.PUT("/events/:id", eventController.UpdateEvent)

	// Category routes

	r.GET("/categories", categoryController.GetCategories)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.GetCategoryByID)

}
