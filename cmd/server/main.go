package main

import (
	"go-evolution-api/internal/api/note"

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
	router.GET("/notes/:id", note.GetNote)

	router.POST("/notes", note.CreateNote)

	router.DELETE("/notes/:id", note.DeleteNote)

	router.Run(":8080")
}
