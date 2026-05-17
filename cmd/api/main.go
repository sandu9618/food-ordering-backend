package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sandu9618/food-ordering-backend/internal/auth"
	"github.com/sandu9618/food-ordering-backend/internal/cart"
	"github.com/sandu9618/food-ordering-backend/internal/category"
	"github.com/sandu9618/food-ordering-backend/internal/database"
	"github.com/sandu9618/food-ordering-backend/internal/middleware"
	"github.com/sandu9618/food-ordering-backend/internal/product"
	"github.com/sandu9618/food-ordering-backend/internal/tenant"
	"github.com/sandu9618/food-ordering-backend/internal/user"
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
		&category.Category{},
		&product.Product{},
		&cart.Cart{},
		&cart.CartItem{},
	)
	if err != nil {
		log.Fatalf("Error migrating tenant table: %v", err)
	}

	router := gin.Default()

	tenantRepo := &tenant.Repository{DB: db}

	router.Use(middleware.TenantMiddleware(tenantRepo))

	userRepo := &user.Repository{DB: db}
	authService := &auth.Service{UserRepo: userRepo}
	authHandler := &auth.Handler{Service: authService}

	router.GET("/health", func(c *gin.Context) {
		tenantID, _ := c.Get("tenant_id")
		c.JSON(200, gin.H{
			"message":   "Server running",
			"tenant_id": tenantID,
		})
	})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	protected := router.Group("/api")
	protected.Use(auth.AuthMiddleware())

	protected.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"userId": c.MustGet("user_id"),
			"role":   c.MustGet("role"),
		})
	})

	admin := router.Group("/admin")

	admin.Use(auth.AuthMiddleware())
	admin.Use(auth.RoleMiddleware("SHOP_ADMIN"))

	admin.GET("/dashboard", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "admin access granted",
		})
	})

	categoryRepo := &category.Repository{DB: db}
	categoryHandler := &category.Handler{Repo: categoryRepo}

	categoryRoutes := admin.Group("/category")
	{
		categoryRoutes.POST("", categoryHandler.CreateCategory)
		categoryRoutes.GET("", categoryHandler.GetAllCategories)
	}

	productRepo := &product.Repository{DB: db}
	productService := &product.Service{Repo: productRepo}
	productHandler := &product.Handler{Service: productService}

	productRoutes := admin.Group("/product")
	{
		productRoutes.POST("", productHandler.Createproduct)
		productRoutes.GET("", productHandler.GetAllProducts)
		productRoutes.GET("/:id", productHandler.GetProductById)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	customer := router.Group("/customer")
	customer.Use(auth.AuthMiddleware())
	customer.Use(auth.RoleMiddleware("CUSTOMER"))

	cartRepo := &cart.Repository{DB: db}
	cartHandler := &cart.Handler{Repo: cartRepo}

	cartRoutes := customer.Group("/cart")
	{
		cartRoutes.POST("", cartHandler.AddItem)
		cartRoutes.GET("", cartHandler.GetCart)
	}

	router.Run(":8080")
}
