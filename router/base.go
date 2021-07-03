package router

import (
	"github.com/gin-gonic/gin"

	myhandler "sme-stage/handler"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {
	r0 := g.Group("/")
	{
		r0.GET("release", myhandler.Release)
	}
}
