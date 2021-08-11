package product

import (
	"gorm.io/gorm"
)

// Product Product
type Product struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}
