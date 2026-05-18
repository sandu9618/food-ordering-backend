package app

import (
	"github.com/sandu9618/food-ordering-backend/internal/cart"
	"github.com/sandu9618/food-ordering-backend/internal/category"
	"github.com/sandu9618/food-ordering-backend/internal/order"
	"github.com/sandu9618/food-ordering-backend/internal/product"
	"github.com/sandu9618/food-ordering-backend/internal/tenant"
	"github.com/sandu9618/food-ordering-backend/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&tenant.Tenant{},
		&user.User{},
		&category.Category{},
		&product.Product{},
		&cart.Cart{},
		&cart.CartItem{},
		&order.Order{},
		&order.OrderItem{},
	)
}
