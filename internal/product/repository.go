package product

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(product *Product) error {
	return r.DB.Create(product).Error
}

func (r *Repository) FindAllByTenantId(tenantID uint) ([]Product, error) {
	var products []Product

	err := r.DB.Where("tenant_id = ?", tenantID).Find(&products).Error

	return products, err
}

func (r *Repository) FindById(id uint, tenantID uint) (*Product, error) {
	var product Product

	err := r.DB.Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repository) Update(product *Product) error {
	return r.DB.Save(product).Error
}

func (r *Repository) Delete(id uint, tenantID uint) error {
	return r.DB.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Product{}).Error
}
