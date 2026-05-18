package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sandu9618/food-ordering-backend/internal/app"
	"github.com/sandu9618/food-ordering-backend/internal/database"
	"github.com/sandu9618/food-ordering-backend/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := app.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	application := app.New(db)

	router := gin.Default()
	routes.Register(router, application)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server exited: %v", err)
	}
}
