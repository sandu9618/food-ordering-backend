package product

type Service struct {
	Repo *Repository
}

func (s *Service) Create(product *Product) error {
	return s.Repo.Create(product)
}

func (s *Service) FindAllByTenantId(tenantID uint) ([]Product, error) {
	return s.Repo.FindAllByTenantId(tenantID)
}

func (s *Service) FindById(id uint, tenantID uint) (*Product, error) {
	return s.Repo.FindById(id, tenantID)
}

func (s *Service) Update(product *Product) error {
	return s.Repo.Update(product)
}

func (s *Service) Delete(id uint, tenantID uint) error {
	return s.Repo.Delete(id, tenantID)
}
