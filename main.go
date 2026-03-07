package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tacna-events-backend/controllers"
	"tacna-events-backend/routes"
)

func main() {
	r := gin.Default()

	// Conectar a PostgreSQL con GORM
	dsn := "host=localhost user=postgres password=jesus dbname=tacna_events_go port=5432 sslmode=disable TimeZone=America/Lima"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate (crea tablas si no existen)


	fmt.Println("✅ Conectado a PostgreSQL con GORM")

	// Inyectar DB en controladores
	userController := controllers.NewUserController(db)
	eventController := controllers.NewEventController(db)
	categoryController := controllers.NewCategoryController(db)
	
	routes.SetupRoutes(r, userController, eventController, categoryController)

	r.Run(":8001")
}
