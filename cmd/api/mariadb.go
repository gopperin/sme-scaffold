package api

import (
	"time"

	"github.com/jinzhu/gorm"

	myconfig "sme-scaffold/internal/config"
	"sme-scaffold/internal/domain/product"
)

// InitDatabase InitDatabase
func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(myconfig.Database.Dialect, myconfig.Database.URL)
	if err != nil {
		return nil, err
	}

	db.LogMode(false)
	db.DB().SetMaxIdleConns(myconfig.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(myconfig.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	db.AutoMigrate(&product.Product{})

	return db, nil
}
