package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	mywire "sme-stage/wire"
)

// SetupProductRouter SetupProductRouter
func SetupProductRouter(g *gin.Engine, db *gorm.DB) {

	// initialize API
	_productAPI := mywire.InitProductAPI(db)

	r := g.Group("/api/v1/products")
	{
		r.GET("/", _productAPI.GetAll)
		r.GET("/:code", _productAPI.GetByCode)
		r.POST("/", _productAPI.Create)
		r.DELETE("/:code", _productAPI.Delete)
	}
}
