package product

import (
	"fmt"
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
	_products := a.Service.GetAll()
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": _products})
}

// GetByCode GetByCode
func (a *API) GetByCode(c *gin.Context) {
	_code := c.Params.ByName("code")
	fmt.Println(_code)
	_prod := a.Service.GetByCode(_code)

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": _prod})
}

// Create Create
func (a *API) Create(c *gin.Context) {
	var _product Product
	err := c.BindJSON(&_product)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_createProduct := a.Service.Save(_product)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": _createProduct})
}

// Delete Delete
func (a *API) Delete(c *gin.Context) {
	_code := c.Params.ByName("code")
	a.Service.Delete(_code)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": "OK"})
}
