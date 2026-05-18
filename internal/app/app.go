package app

import (
	"github.com/sandu9618/food-ordering-backend/internal/auth"
	"github.com/sandu9618/food-ordering-backend/internal/cart"
	"github.com/sandu9618/food-ordering-backend/internal/category"
	"github.com/sandu9618/food-ordering-backend/internal/order"
	"github.com/sandu9618/food-ordering-backend/internal/product"
	"github.com/sandu9618/food-ordering-backend/internal/tenant"
	"github.com/sandu9618/food-ordering-backend/internal/user"
	"gorm.io/gorm"
)

// App holds shared dependencies for HTTP handlers (composition root).
type App struct {
	DB         *gorm.DB
	TenantRepo *tenant.Repository
	Auth       *auth.Handler
	Category   *category.Handler
	Product    *product.Handler
	Cart       *cart.Handler
	Order      *order.Handler
}

func New(db *gorm.DB) *App {
	tenantRepo := &tenant.Repository{DB: db}

	userRepo := &user.Repository{DB: db}
	authService := &auth.Service{UserRepo: userRepo}
	authHandler := &auth.Handler{Service: authService}

	categoryRepo := &category.Repository{DB: db}
	categoryHandler := &category.Handler{Repo: categoryRepo}

	productRepo := &product.Repository{DB: db}
	productService := &product.Service{Repo: productRepo}
	productHandler := &product.Handler{Service: productService}

	cartRepo := &cart.Repository{DB: db}
	cartHandler := &cart.Handler{Repo: cartRepo}

	orderRepo := &order.Repository{DB: db}
	orderService := &order.Service{DB: db, OrderRepo: orderRepo}
	orderHandler := &order.Handler{Service: orderService}

	return &App{
		DB:         db,
		TenantRepo: tenantRepo,
		Auth:       authHandler,
		Category:   categoryHandler,
		Product:    productHandler,
		Cart:       cartHandler,
		Order:      orderHandler,
	}
}
