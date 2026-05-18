package order

import (
	"errors"
	"fmt"

	"github.com/sandu9618/food-ordering-backend/internal/cart"
	"github.com/sandu9618/food-ordering-backend/internal/product"
	"gorm.io/gorm"
)

var ErrEmptyCart = errors.New("cart is empty")

type Service struct {
	DB        *gorm.DB
	OrderRepo *Repository
}

var allowedStatuses = map[string]struct{}{
	"PENDING":    {},
	"CONFIRMED":  {},
	"PREPARING":  {},
	"READY":      {},
	"DELIVERED":  {},
	"CANCELLED":  {},
}

func (s *Service) CreateOrder(userID uint, tenantID uint) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		orderRepo := &Repository{DB: tx}
		cartRepo := &cart.Repository{DB: tx}
		productRepo := &product.Repository{DB: tx}

		userCart, err := cartRepo.FindOrCreateCart(userID, tenantID)
		if err != nil {
			return err
		}

		if len(userCart.Items) == 0 {
			return ErrEmptyCart
		}

		var total float64
		var orderItems []OrderItem

		for _, line := range userCart.Items {
			if line.Quantity <= 0 {
				return fmt.Errorf("invalid quantity for product %d", line.ProductID)
			}

			p, err := productRepo.FindById(line.ProductID, tenantID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("product %d not found or not in tenant", line.ProductID)
				}
				return err
			}

			subtotal := p.Price * float64(line.Quantity)
			total += subtotal

			orderItems = append(orderItems, OrderItem{
				ProductID:   p.ID,
				ProductName: p.Name,
				Price:       p.Price,
				Quantity:    line.Quantity,
				Subtotal:    subtotal,
			})
		}

		newOrder := &Order{
			TenantID:    tenantID,
			UserID:      userID,
			Status:      "PENDING",
			TotalAmount: total,
			Items:       orderItems,
		}

		if err := orderRepo.Create(newOrder); err != nil {
			return err
		}

		return cartRepo.ClearCart(userCart.ID)
	})
}

func (s *Service) OrdersForTenant(tenantID uint) ([]Order, error) {
	return s.OrderRepo.FindByTenant(tenantID)
}

func (s *Service) SetOrderStatus(orderID, tenantID uint, status string) error {
	if _, ok := allowedStatuses[status]; !ok {
		return fmt.Errorf("invalid status: %s", status)
	}
	return s.OrderRepo.UpdateStatus(orderID, tenantID, status)
}
