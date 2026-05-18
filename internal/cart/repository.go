package cart

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindOrCreateCart(
	userID uint,
	tenantID uint,
) (*Cart, error) {

	var cart Cart

	err := r.DB.
		Preload("Items").
		Where(
			"user_id = ? AND tenant_id = ?",
			userID,
			tenantID,
		).
		First(&cart).Error

	if err == nil {
		return &cart, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	cart = Cart{
		UserID:   userID,
		TenantID: tenantID,
	}

	err = r.DB.Create(&cart).Error

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *Repository) AddItem(
	item *CartItem,
) error {

	return r.DB.Create(item).Error
}

func (r *Repository) GetCart(
	userID uint,
	tenantID uint,
) (*Cart, error) {
	return r.FindOrCreateCart(userID, tenantID)
}

func (r *Repository) ClearCart(
	cartID uint,
) error {

	return r.DB.
		Where("cart_id = ?", cartID).
		Delete(&CartItem{}).Error
}
