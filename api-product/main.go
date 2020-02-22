package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngiyshhk/bff-cluster/api-product/model"
)

func main() {
	r := gin.Default()
	products := r.Group("/products")

	products.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": model.ProductsMock,
		})
	})
	products.GET(":id", func(c *gin.Context) {
		id := c.Param("id")
		for _, p := range model.ProductsMock {
			if p.ID == id {
				c.JSON(200, gin.H{
					"product": p,
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"error": "not found",
		})
	})
	r.Run()
}
