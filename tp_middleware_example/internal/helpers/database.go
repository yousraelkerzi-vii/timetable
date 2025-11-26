package helpers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"  // Garder le _ devant
)

const databasePath = "timetable.db"

// OpenDatabase ouvre une connexion à la base de données SQLite
func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}