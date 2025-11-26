package agendas

import (
	"context"
	"net/http"
)

// AgendaContextMiddleware parse l'ID de l'agenda depuis l'URL et l'ajoute au contexte
func AgendaContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'ID depuis l'URL
		agendaId := r.PathValue("id")

		if agendaId == "" {
			http.Error(w, "Agenda ID is required", http.StatusBadRequest)
			return
		}

		// Ajouter l'ID au contexte
		ctx := context.WithValue(r.Context(), "agendaId", agendaId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
