package repository

import (
	"database/sql"
	"go-evolution-api/internal/model"
)

func GetNoteByID(db *sql.DB, id string) (*model.Note, error) {
	var note model.Note
	query := "SELECT id, title, content FROM note WHERE id = ?"
	if err := db.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content); err != nil {
		if err == sql.ErrNoRows {
			// Kein Ergebnis ist kein Fehler in diesem Kontext
			return nil, nil
		}
		return nil, err
	}
	return &note, nil
}
func CreateNote(db *sql.DB, note model.Note) error {
	stmt, err := db.Prepare("INSERT INTO note (title, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(note.Title, note.Content)
	if err != nil {
		return err
	}

	return nil
}

// Löscht Notiz nach id
func DeleteNoteByID(db *sql.DB, id string) error {
	query := "DELETE FROM note WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
