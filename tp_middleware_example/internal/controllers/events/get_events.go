package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/events"
	"net/http"
)

// GetEvents godoc
// @Summary      Récupérer tous les événements
// @Description  Récupère la liste complète des événements du calendrier
// @Tags         events
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  models.ErrorResponse
// @Router       /events [get]
func GetEvents(w http.ResponseWriter, r *http.Request) {
	db := helpers.OpenDatabase()
	defer db.Close()

	eventsData, errResp := events.GetAllEvents(db)
	if errResp != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventsData)
}