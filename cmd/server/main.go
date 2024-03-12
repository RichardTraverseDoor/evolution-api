package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Grundlegende Route zum Testen, ob der Server läuft
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running!",
		})
	})

	router.Run()
}
