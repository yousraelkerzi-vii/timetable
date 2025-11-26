package agendas

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/agendas"
	"net/http"
)

// CreateAgenda godoc
// @Summary      Créer un agenda
// @Description  Crée un nouvel agenda UCA
// @Tags         agendas
// @Accept       json
// @Produce      json
// @Param        agenda  body      models.Agenda  true  "Données de l'agenda (sans ID)"
// @Success      201     {object}  models.Agenda
// @Failure      400     {object}  models.ErrorResponse
// @Failure      500     {object}  models.ErrorResponse
// @Router       /agendas [post]
func CreateAgenda(w http.ResponseWriter, r *http.Request) {
	var agenda models.Agenda

	// Parser le body JSON
	err := json.NewDecoder(r.Body).Decode(&agenda)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}

	db := helpers.OpenDatabase()
	defer db.Close()

	createdAgenda, errResp := agendas.CreateAgenda(db, &agenda)
	if errResp != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAgenda)
}