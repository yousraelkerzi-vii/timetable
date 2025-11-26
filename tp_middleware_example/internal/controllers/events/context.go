package events

import (
	"context"
	"net/http"
)

// EventContextMiddleware parse l'ID de l'événement depuis l'URL et l'ajoute au contexte
func EventContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'ID depuis l'URL
		eventId := r.PathValue("id")

		if eventId == "" {
			http.Error(w, "Event ID is required", http.StatusBadRequest)
			return
		}

		// Ajouter l'ID au contexte
		ctx := context.WithValue(r.Context(), "eventId", eventId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}