package main

import (
    "github.com/gin-gonic/gin"
    "tacna-events-backend/routes"
    "tacna-events-backend/db"
)

func main() {
    r := gin.Default()

    // conectar DB
    conn, err := db.ConnectDB()
    if err != nil {
        panic(err)
    }
    defer conn.Close(nil)

    routes.SetupRoutes(r)

    r.Run(":8001")
}