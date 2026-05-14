package main

import (
	"log"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sandu9618/food-ordering-backend/internal/database"
	"github.com/sandu9618/food-ordering-backend/internal/middleware"
	"github.com/sandu9618/food-ordering-backend/internal/tenant"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(
		&tenant.Tenant{},
		&user.User{},
	)
	if err != nil {
		log.Fatalf("Error migrating tenant table: %v", err)
	}

	router := gin.Default()

	tenantRepo := &tenant.Repository{DB: db}

	router.Use(middleware.TenantMiddleware(tenantRepo))

	router.GET("/health", func(c *gin.Context) {
		tenantID, _ := c.Get("tenant_id")
		c.JSON(200, gin.H{
			"message":   "Server running",
			"tenant_id": tenantID,
		})
	})

	router.Run(":8080")
}
