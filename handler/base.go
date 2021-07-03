package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Release Release
func Release(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"release": "2021.07.03", "version": "v20210703"})
}
