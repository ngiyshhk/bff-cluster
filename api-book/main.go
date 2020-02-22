package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngiyshhk/bff-cluster/api-book/model"
)

func main() {
	r := gin.Default()

	book := r.Group("books")
	book.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"books": model.BooksMock,
		})
	})
	book.GET(":id", func(c *gin.Context) {
		id := c.Param("id")
		for _, b := range model.BooksMock {
			if b.ID == id {
				c.JSON(200, gin.H{
					"book": b,
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
