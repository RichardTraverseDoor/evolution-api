package note

import (
	"internal/model"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var newNote model.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Logik zum Speichern der Notiz
	c.JSON(200, gin.H{"message": "Notiz erstellt", "note": newNote})
}
