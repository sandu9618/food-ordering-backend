package tenant

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindBySubdomain(subdomain string) (*Tenant, error) {
	var tenant Tenant

	err := r.DB.Where("subdomain = ?", subdomain).First(&tenant).Error

	if err != nil {
		return nil, err
	}

	return &tenant, nil
}
