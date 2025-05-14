package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Create /user endpoint returning JSON
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":    123,
			"name":  "Alice",
			"email": "alice@example.com",
		})
	})

	// Start the server
	r.Run() // Listens on 0.0.0.0:8080
}
