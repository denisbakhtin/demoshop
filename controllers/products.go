package controllers

import (
	"github.com/denisbakhtin/demoshop/models"
	"github.com/denisbakhtin/demoshop/system"
	"github.com/gin-gonic/gin"
)

//ProductsGet handles GET /api/products route
func ProductsGet(c *gin.Context) {
	db := models.GetDB()
	domain := system.GetConfig().Domain

	var list []models.Product
	db.Preload("Vendor").Order("id").Find(&list)
	for i := 0; i < len(list); i++ {
		list[i].ImageURL = fullImageURL(domain, list[i].ImageURL)
		list[i].Vendor.ImageURL = fullImageURL(domain, list[i].Vendor.ImageURL)
	}
	c.JSON(200, list)
}

//ProductGet handles /api/products/:id route
func ProductGet(c *gin.Context) {
	db := models.GetDB()
	domain := system.GetConfig().Domain
	id := c.Param("id")

	product := &models.Product{}
	db.Preload("Vendor").First(product, id)
	if product.ID == 0 {
		c.JSON(404, nil)
		return
	}
	product.ImageURL = fullImageURL(domain, product.ImageURL)
	product.Vendor.ImageURL = fullImageURL(domain, product.Vendor.ImageURL)
	c.JSON(200, product)
}

func fullImageURL(domain string, url string) string {
	return "http://" + domain + "/public" + url
}
