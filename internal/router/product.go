package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	mywire "sme-scaffold/internal/wire"
)

// SetupProductRouter SetupProductRouter
func SetupProductRouter(g *gin.Engine, db *gorm.DB) {

	// initialize API
	productAPI := mywire.InitProductAPI(db)

	r := g.Group("/api/v1/products")
	{
		r.GET("", productAPI.GetAll)
		r.GET("/:code", productAPI.GetByCode)
		r.POST("/", productAPI.Create)
		r.DELETE("/:code", productAPI.Delete)
	}
}
