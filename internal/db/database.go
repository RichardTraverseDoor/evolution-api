package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	// Konfigurationen
	dsn := "root:Jirachi1.0@tcp(localhost:3306)/evolution"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Überprüfen der Verbindung
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Verbindung zur Datenbank erfolgreich!")
	return db
}
