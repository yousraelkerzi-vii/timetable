package events

import (
	"database/sql"
	"middleware/example/internal/models"
	eventsRepo "middleware/example/internal/repositories/events"
)

// GetAllEvents récupère tous les événements
func GetAllEvents(db *sql.DB) ([]models.Event, *models.ErrorResponse) {
	events, err := eventsRepo.GetAllEvents(db)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve events",
		}
	}

	// Retourner un tableau vide au lieu de nil si aucun événement
	if events == nil {
		events = []models.Event{}
	}

	return events, nil
}

// GetEventByID récupère un événement spécifique par son ID
func GetEventByID(db *sql.DB, id string) (*models.Event, *models.ErrorResponse) {
	event, err := eventsRepo.GetEventByID(db, id)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve event",
		}
	}

	if event == nil {
		return nil, &models.ErrorResponse{
			Message: "Event not found",
		}
	}

	return event, nil
}