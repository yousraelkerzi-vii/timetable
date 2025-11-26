package agendas

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/agendas"
	"net/http"
)

// DeleteAgenda godoc
// @Summary      Supprimer un agenda
// @Description  Supprime un agenda par son ID
// @Tags         agendas
// @Produce      json
// @Param        id   path      string  true  "Agenda ID"
// @Success      204  "No Content"
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /agendas/{id} [delete]
func DeleteAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaId, ok := ctx.Value("agendaId").(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Invalid agenda ID",
		})
		return
	}

	db := helpers.OpenDatabase()
	defer db.Close()

	errResp := agendas.DeleteAgenda(db, agendaId)
	if errResp != nil {
		if errResp.Message == "Agenda not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}