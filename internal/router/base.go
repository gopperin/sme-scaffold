package router

import (
	"github.com/gin-gonic/gin"

	mywire "sme-scaffold/internal/wire"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {

	// initialize API
	_baseAPI := mywire.InitBaseAPI()

	r := g.Group("/")
	{
		r.GET("health", _baseAPI.Health)
		r.GET("release", _baseAPI.Release)
	}
}
