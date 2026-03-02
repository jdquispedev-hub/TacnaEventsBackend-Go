package main

import (
    "github.com/gin-gonic/gin"
    "tacna-events-backend/routes"
    "tacna-events-backend/db"
    "tacna-events-backend/controllers"
)

func main() {
    r := gin.Default()

    // conectar DB con pool
    pool, err := db.ConnectDB()
    if err != nil {
        panic(err)
    }
    defer pool.Close()
    
    // Asignar pool global
    db.DB = pool

    // Inyectar DB en controladores
    userController := controllers.NewUserController(pool)
    eventController := controllers.NewEventController(pool)

    routes.SetupRoutes(r, userController, eventController)

    r.Run(":8001")
}