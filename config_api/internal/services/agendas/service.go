package agendas

import (
	"database/sql"
	"middleware/example/internal/models"
	agendasRepo "middleware/example/internal/repositories/agendas"

	"github.com/google/uuid"
)

// GetAllAgendas récupère tous les agendas
func GetAllAgendas(db *sql.DB) ([]models.Agenda, *models.ErrorResponse) {
	agendas, err := agendasRepo.GetAllAgendas(db)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve agendas",
		}
	}

	if agendas == nil {
		agendas = []models.Agenda{}
	}

	return agendas, nil
}

// GetAgendaByID récupère un agenda par son ID
func GetAgendaByID(db *sql.DB, id string) (*models.Agenda, *models.ErrorResponse) {
	agenda, err := agendasRepo.GetAgendaByID(db, id)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve agenda",
		}
	}

	if agenda == nil {
		return nil, &models.ErrorResponse{
			Message: "Agenda not found",
		}
	}

	return agenda, nil
}

// CreateAgenda crée un nouvel agenda
func CreateAgenda(db *sql.DB, agenda *models.Agenda) (*models.Agenda, *models.ErrorResponse) {
	// Générer un UUID pour le nouvel agenda
	agenda.ID = uuid.New().String()

	// Valider les champs requis
	if agenda.UcaID == "" || agenda.Name == "" {
		return nil, &models.ErrorResponse{
			Message: "ucaId and name are required",
		}
	}

	err := agendasRepo.CreateAgenda(db, agenda)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to create agenda",
		}
	}

	return agenda, nil
}

// UpdateAgenda met à jour un agenda existant
func UpdateAgenda(db *sql.DB, id string, agenda *models.Agenda) (*models.Agenda, *models.ErrorResponse) {
	// Vérifier que l'agenda existe
	existing, err := agendasRepo.GetAgendaByID(db, id)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to retrieve agenda",
		}
	}
	if existing == nil {
		return nil, &models.ErrorResponse{
			Message: "Agenda not found",
		}
	}

	// Mettre à jour l'ID
	agenda.ID = id

	// Valider les champs requis
	if agenda.UcaID == "" || agenda.Name == "" {
		return nil, &models.ErrorResponse{
			Message: "ucaId and name are required",
		}
	}

	err = agendasRepo.UpdateAgenda(db, agenda)
	if err != nil {
		return nil, &models.ErrorResponse{
			Message: "Failed to update agenda",
		}
	}

	return agenda, nil
}

// DeleteAgenda supprime un agenda
func DeleteAgenda(db *sql.DB, id string) *models.ErrorResponse {
	err := agendasRepo.DeleteAgenda(db, id)
	if err == sql.ErrNoRows {
		return &models.ErrorResponse{
			Message: "Agenda not found",
		}
	}
	if err != nil {
		return &models.ErrorResponse{
			Message: "Failed to delete agenda",
		}
	}

	return nil
}