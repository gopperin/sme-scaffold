//+build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"sme-scaffold/domain/base"
	"sme-scaffold/domain/product"
)

// InitBaseAPI init base api wire
func InitBaseAPI() base.API {
	wire.Build(base.ProvideAPI, base.ProvideService)
	return base.API{}
}

// InitBaseAPI init base api wire
func InitProductAPI(db *gorm.DB) product.API {
	wire.Build(product.ProvideRepository, product.ProvideAPI, product.ProvideService)
	return product.API{}
}
