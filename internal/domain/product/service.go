package product

// Service Service
type Service struct {
	Repository Repository
}

// ProvideService ProvideService
func ProvideService(r Repository) Service {
	return Service{Repository: r}
}

// GetAll GetAll
func (s *Service) GetAll() []Product {
	return s.Repository.GetAll()
}

// GetByCode GetByCode
func (s *Service) GetByCode(code string) Product {
	return s.Repository.GetByCode(code)
}

// Save Save
func (s *Service) Save(product Product) Product {
	s.Repository.Save(product)
	return product
}

// Delete Delete
func (s *Service) Delete(code string) {
	s.Repository.Delete(code)
}
