package main

import (
	"log"
	"net/http"

	eventsCtrl "middleware/example/internal/controllers/events"
	"middleware/example/internal/helpers"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "middleware/example/api"
)

func init() {
	// Initialisation de la base de données
	db := helpers.OpenDatabase()
	defer db.Close()

	// Création de la table events
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id TEXT PRIMARY KEY,
			agenda_ids TEXT NOT NULL,
			uid TEXT NOT NULL,
			description TEXT,
			name TEXT NOT NULL,
			start DATETIME NOT NULL,
			end DATETIME NOT NULL,
			location TEXT,
			last_update DATETIME NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to create events table:", err)
	}

	log.Println("Database initialized successfully")
}

func main() {
	// Création du routeur
	mux := http.NewServeMux()

	// Route pour la documentation Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Routes pour les événements
	mux.HandleFunc("GET /timetable_api/events", eventsCtrl.GetEvents)

	// Route avec middleware pour parser l'ID
	mux.Handle("GET /timetable_api/events/{id}",
		eventsCtrl.EventContextMiddleware(
			http.HandlerFunc(eventsCtrl.GetEvent),
		),
	)

	// Démarrage du serveur
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger documentation available at http://localhost%s/swagger/index.html", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}