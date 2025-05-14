package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from Gin!")
	})

	// Create a GET endpoint that returns JSON data
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Data retrieved successfully",
			"data": gin.H{
				"id":    123,
				"name":  "Product",
				"price": 29.99,
			},
		})
	})

	r.Run()
}
