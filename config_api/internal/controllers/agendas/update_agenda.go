package agendas

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/agendas"
	"net/http"
)

// UpdateAgenda godoc
// @Summary      Mettre à jour un agenda
// @Description  Met à jour un agenda existant par son ID
// @Tags         agendas
// @Accept       json
// @Produce      json
// @Param        id      path      string         true  "Agenda ID"
// @Param        agenda  body      models.Agenda  true  "Nouvelles données de l'agenda"
// @Success      200     {object}  models.Agenda
// @Failure      400     {object}  models.ErrorResponse
// @Failure      404     {object}  models.ErrorResponse
// @Failure      500     {object}  models.ErrorResponse
// @Router       /agendas/{id} [put]
func UpdateAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaId, ok := ctx.Value("agendaId").(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Invalid agenda ID",
		})
		return
	}

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

	updatedAgenda, errResp := agendas.UpdateAgenda(db, agendaId, &agenda)
	if errResp != nil {
		if errResp.Message == "Agenda not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAgenda)
}