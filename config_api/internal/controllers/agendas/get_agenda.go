package agendas

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/agendas"
	"net/http"
)

// GetAgenda godoc
// @Summary      Récupérer un agenda
// @Description  Récupère un agenda spécifique par son ID
// @Tags         agendas
// @Produce      json
// @Param        id   path      string  true  "Agenda ID"
// @Success      200  {object}  models.Agenda
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /agendas/{id} [get]
func GetAgenda(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agendaId, ok := ctx.Value("agendaId").(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid agenda ID",
		})
		return
	}

	db := helpers.OpenDatabase()
	defer db.Close()

	agenda, errResp := agendas.GetAgendaByID(db, agendaId)
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
	json.NewEncoder(w).Encode(agenda)
}