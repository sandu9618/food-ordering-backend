package order

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(order *Order) error {
	return r.DB.Create(order).Error
}

func (r *Repository) FindByTenant(tenantID uint) ([]Order, error) {
	var orders []Order

	err := r.DB.Preload("Items").Where("tenant_id = ?", tenantID).Find(&orders).Error

	return orders, err
}

func (r *Repository) UpdateStatus(orderID uint, tenantID uint, status string) error {
	return r.DB.Model(&Order{}).
		Where("id = ? AND tenant_id = ?", orderID, tenantID).
		Update("status", status).Error
}
