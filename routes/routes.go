package routes

import (
	"github.com/gin-gonic/gin"
	"tacna-events-backend/controllers"
)

func SetupRoutes(r gin.IRouter){
	r.GET("/users", controllers.GetUsers)
}
