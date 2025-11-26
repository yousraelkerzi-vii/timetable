package events

import (
	"database/sql"
	"encoding/json"
	"middleware/example/internal/models"
)

// GetAllEvents récupère tous les événements de la base de données
func GetAllEvents(db *sql.DB) ([]models.Event, error) {
	rows, err := db.Query(`
		SELECT id, agenda_ids, uid, description, name, start, end, location, last_update 
		FROM events
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var agendaIdsJSON string

		err := rows.Scan(
			&event.ID,
			&agendaIdsJSON,
			&event.UID,
			&event.Description,
			&event.Name,
			&event.Start,
			&event.End,
			&event.Location,
			&event.LastUpdate,
		)
		if err != nil {
			return nil, err
		}

		// Désérialiser agendaIds depuis JSON
		if err := json.Unmarshal([]byte(agendaIdsJSON), &event.AgendaIDs); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// GetEventByID récupère un événement spécifique par son ID
func GetEventByID(db *sql.DB, id string) (*models.Event, error) {
	var event models.Event
	var agendaIdsJSON string

	err := db.QueryRow(`
		SELECT id, agenda_ids, uid, description, name, start, end, location, last_update 
		FROM events 
		WHERE id = ?
	`, id).Scan(
		&event.ID,
		&agendaIdsJSON,
		&event.UID,
		&event.Description,
		&event.Name,
		&event.Start,
		&event.End,
		&event.Location,
		&event.LastUpdate,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Aucun événement trouvé
	}
	if err != nil {
		return nil, err
	}

	// Désérialiser agendaIds depuis JSON
	if err := json.Unmarshal([]byte(agendaIdsJSON), &event.AgendaIDs); err != nil {
		return nil, err
	}

	return &event, nil
}