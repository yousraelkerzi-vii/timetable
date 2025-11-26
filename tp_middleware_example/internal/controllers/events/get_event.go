package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/events"
	"net/http"
)

// GetEvent godoc
// @Summary      Récupérer un événement
// @Description  Récupère un événement spécifique par son ID
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /events/{id} [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid event ID",
		})
		return
	}

	db := helpers.OpenDatabase()
	defer db.Close()

	event, errResp := events.GetEventByID(db, eventId)
	if errResp != nil {
		if errResp.Message == "Event not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}