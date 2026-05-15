package category

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(category *Category) error {
	return r.DB.Create(category).Error
}

func (r *Repository) FindAllByTenant(tenantID uint) ([]Category, error) {
	var categories []Category

	err := r.DB.Where("tenant_id = ?", tenantID).Find(&categories).Error

	return categories, err
}
