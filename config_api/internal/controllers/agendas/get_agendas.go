package agendas

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/agendas"
	"net/http"
)

// GetAgendas godoc
// @Summary      Récupérer tous les agendas
// @Description  Récupère la liste complète des agendas UCA
// @Tags         agendas
// @Produce      json
// @Success      200  {array}   models.Agenda
// @Failure      500  {object}  models.ErrorResponse
// @Router       /agendas [get]
func GetAgendas(w http.ResponseWriter, r *http.Request) {
	db := helpers.OpenDatabase()
	defer db.Close()

	agendasData, errResp := agendas.GetAllAgendas(db)
	if errResp != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agendasData)
}