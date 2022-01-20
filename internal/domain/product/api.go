package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API API
type API struct {
	Service Service
}

// ProvideAPI ProvideAPI
func ProvideAPI(service Service) API {
	return API{Service: service}
}

// GetAll GetAll
func (a *API) GetAll(c *gin.Context) {
	products := a.Service.GetAll()
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": products})
}

// GetByCode GetByCode
func (a *API) GetByCode(c *gin.Context) {
	code := c.Params.ByName("code")
	prod := a.Service.GetByCode(code)

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": prod})
}

// Create Create
func (a *API) Create(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	createProduct := a.Service.Save(product)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": createProduct})
}

// Delete Delete
func (a *API) Delete(c *gin.Context) {
	code := c.Params.ByName("code")
	a.Service.Delete(code)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": "OK"})
}
