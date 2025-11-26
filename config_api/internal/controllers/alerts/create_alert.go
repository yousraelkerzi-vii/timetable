package alerts

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/alerts"
	"net/http"
)

// CreateAlert godoc
// @Summary      Créer une alerte
// @Description  Crée une nouvelle alerte
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        alert  body      models.Alert  true  "Données de l'alerte (sans ID)"
// @Success      201    {object}  models.Alert
// @Failure      400    {object}  models.ErrorResponse
// @Failure      500    {object}  models.ErrorResponse
// @Router       /alerts [post]
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var alert models.Alert

	// Parser le body JSON
	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}

	db := helpers.OpenDatabase()
	defer db.Close()

	createdAlert, errResp := alerts.CreateAlert(db, &alert)
	if errResp != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAlert)
}