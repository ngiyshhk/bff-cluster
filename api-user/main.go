package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngiyshhk/bff-cluster/api-user/model"
)

func main() {
	r := gin.Default()

	user := r.Group("users")
	user.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": model.UsersMock,
		})
	})
	user.GET(":id", func(c *gin.Context) {
		id := c.Param("id")
		for _, u := range model.UsersMock {
			if u.ID == id {
				c.JSON(200, gin.H{
					"user": u,
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
