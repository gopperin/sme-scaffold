package product

import (
	"fmt"
	"gorm.io/gorm"
)

// Repository Repository
type Repository struct {
	DB *gorm.DB
}

// ProvideRepository ProvideRepository
func ProvideRepository(DB *gorm.DB) Repository {
	return Repository{DB: DB}
}

// GetAll GetAll
func (r *Repository) GetAll() []Product {
	var objs []Product
	r.DB.Find(&objs)
	return objs
}

// GetByCode GetByCode
func (r *Repository) GetByCode(code string) Product {
	fmt.Println(code)
	var obj Product
	r.DB.Where("code = ?", code).First(&obj)
	fmt.Println(obj)
	return obj
}

// Save Save
func (r *Repository) Save(prod Product) Product {
	r.DB.Save(&prod)
	return prod
}

// Delete Delete
func (r *Repository) Delete(code string) error {
	return r.DB.Unscoped().Where("code=?", code).Delete(Product{}).Error
}
