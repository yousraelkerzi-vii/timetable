package models

import "time"

type Event struct {
    ID          string    `json:"id"`
    AgendaIDs   []string  `json:"agendaIds"`
    UID         string    `json:"uid"`
    Description string    `json:"description"`
    Name        string    `json:"name"`
    Start       time.Time `json:"start"`
    End         time.Time `json:"end"`
    Location    string    `json:"location"`
    LastUpdate  time.Time `json:"lastUpdate"`
}

// ErrorResponse représente une erreur retournée par l'API
type ErrorResponse struct {
	Message string `json:"message"`
}



