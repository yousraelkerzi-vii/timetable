package alerts

import (
	"database/sql"
	"middleware/example/internal/models"
	alertsRepo "middleware/example/internal/repositories/alerts"

	"github.com/google/uuid"
)

// GetAllAlerts récupère toutes les alertes
func GetAllAlerts(db *sql.DB) ([]models.Alert, *models.ErrorResponse) {
	alerts, err := alertsRepo.GetAllAlerts(db)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve alerts",
		}
	}

	if alerts == nil {
		alerts = []models.Alert{}
	}

	return alerts, nil
}

// GetAlertByID récupère une alerte par son ID
func GetAlertByID(db *sql.DB, id string) (*models.Alert, *models.ErrorResponse) {
	alert, err := alertsRepo.GetAlertByID(db, id)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve alert",
		}
	}

	if alert == nil {
		return nil, &models.ErrorResponse{
			Message: "Alert not found",
		}
	}

	return alert, nil
}

// CreateAlert crée une nouvelle alerte
func CreateAlert(db *sql.DB, alert *models.Alert) (*models.Alert, *models.ErrorResponse) {
	// Générer un UUID pour la nouvelle alerte
	alert.ID = uuid.New().String()

	// Valider les champs requis
	if alert.Email == "" || alert.AgendaID == "" {
		return nil, &models.ErrorResponse{
			Message: "email and agendaId are required",
		}
	}

	err := alertsRepo.CreateAlert(db, alert)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to create alert",
		}
	}

	return alert, nil
}

// UpdateAlert met à jour une alerte existante
func UpdateAlert(db *sql.DB, id string, alert *models.Alert) (*models.Alert, *models.ErrorResponse) {
	// Vérifier que l'alerte existe
	existing, err := alertsRepo.GetAlertByID(db, id)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve alert",
		}
	}
	if existing == nil {
		return nil, &models.ErrorResponse{
			Message: "Alert not found",
		}
	}

	// Mettre à jour l'ID
	alert.ID = id

	// Valider les champs requis
	if alert.Email == "" || alert.AgendaID == "" {
		return nil, &models.ErrorResponse{
			Message: "email and agendaId are required",
		}
	}

	err = alertsRepo.UpdateAlert(db, alert)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to update alert",
		}
	}

	return alert, nil
}

// DeleteAlert supprime une alerte
func DeleteAlert(db *sql.DB, id string) *models.ErrorResponse {
	err := alertsRepo.DeleteAlert(db, id)
	if err == sql.ErrNoRows {
		return &models.ErrorResponse{
			Message: "Alert not found",
		}
	}
	if err != nil {
		return &models.ErrorResponse{
			Message: "Failed to delete alert",
		}
	}

	return nil
}