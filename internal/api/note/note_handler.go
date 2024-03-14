package note

import (
	"go-evolution-api/internal/db"
	"go-evolution-api/internal/model"
	"go-evolution-api/internal/repository"

	"github.com/gin-gonic/gin"
)

// Gibt eine Notiz zurück
func GetNote(c *gin.Context) {
	db := db.ConnectToDB()
	defer db.Close()

	id := c.Param("id")
	note, err := repository.GetNoteByID(db, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if note == nil {
		c.JSON(404, gin.H{"error": "Notiz nicht gefunden"})
		return
	}

	c.JSON(200, gin.H{"note": note})
}

// Erstellt neue Notiz
func CreateNote(c *gin.Context) {
	var newNote model.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database := db.ConnectToDB()
	defer database.Close()

	if err := repository.CreateNote(database, newNote); err != nil {
		c.JSON(500, gin.H{"error": "Fehler beim Einfügen der Notiz: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Notiz erstellt", "note": newNote})
}

// Löschen der Notiz
func DeleteNote(c *gin.Context) {
	db := db.ConnectToDB()
	defer db.Close()

	id := c.Param("id")
	err := repository.DeleteNoteByID(db, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Notiz erfolgreich gelöscht"})
}
