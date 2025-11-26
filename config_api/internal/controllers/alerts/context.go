package alerts

import (
	"context"
	"net/http"
)

// AlertContextMiddleware parse l'ID de l'alerte depuis l'URL et l'ajoute au contexte
func AlertContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'ID depuis l'URL
		alertId := r.PathValue("id")

		if alertId == "" {
			http.Error(w, "Alert ID is required", http.StatusBadRequest)
			return
		}

		// Ajouter l'ID au contexte
		ctx := context.WithValue(r.Context(), "alertId", alertId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
