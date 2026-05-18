package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sandu9618/food-ordering-backend/internal/app"
	"github.com/sandu9618/food-ordering-backend/internal/auth"
	"github.com/sandu9618/food-ordering-backend/internal/middleware"
)

func Register(engine *gin.Engine, a *app.App) {
	engine.Use(middleware.TenantMiddleware(a.TenantRepo))

	engine.GET("/health", func(c *gin.Context) {
		tenantID, _ := c.Get("tenant_id")
		c.JSON(200, gin.H{
			"message":   "Server running",
			"tenant_id": tenantID,
		})
	})

	authRoutes := engine.Group("/auth")
	{
		authRoutes.POST("/register", a.Auth.Register)
		authRoutes.POST("/login", a.Auth.Login)
	}

	protected := engine.Group("/api")
	protected.Use(auth.AuthMiddleware())

	protected.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"userId": c.MustGet("user_id"),
			"role":   c.MustGet("role"),
		})
	})

	admin := engine.Group("/admin")
	admin.Use(auth.AuthMiddleware())
	admin.Use(auth.RoleMiddleware("SHOP_ADMIN"))

	admin.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "admin access granted",
		})
	})

	categoryRoutes := admin.Group("/category")
	{
		categoryRoutes.POST("", a.Category.CreateCategory)
		categoryRoutes.GET("", a.Category.GetAllCategories)
	}

	productRoutes := admin.Group("/product")
	{
		productRoutes.POST("", a.Product.Createproduct)
		productRoutes.GET("", a.Product.GetAllProducts)
		productRoutes.GET("/:id", a.Product.GetProductById)
		productRoutes.PUT("/:id", a.Product.UpdateProduct)
		productRoutes.DELETE("/:id", a.Product.DeleteProduct)
	}

	admin.GET("/orders", a.Order.GetOrders)
	admin.PUT("/orders/:id/status", a.Order.UpdateStatus)

	customer := engine.Group("/customer")
	customer.Use(auth.AuthMiddleware())
	customer.Use(auth.RoleMiddleware("CUSTOMER"))

	cartRoutes := customer.Group("/cart")
	{
		cartRoutes.POST("/items", a.Cart.AddItem)
		cartRoutes.POST("", a.Cart.AddItem)
		cartRoutes.GET("", a.Cart.GetCart)
	}

	customer.POST("/checkout", a.Order.Checkout)
}
