package models

import(
	time
)

// Agenda représente un agenda UCA
type Agenda struct {
	ID       string `json:"id"`
	UcaID    string `json:"ucaId"`    // ID pour récupérer le fichier iCal
	Name     string `json:"name"`
	
	

// Alert représente une alerte configurée
type Alert struct {
	ID        string `json:"id"`
	Email     string `json:"email"`     // Destinataire
	AgendaID  string `json:"agendaId"`  // Agenda concerné
	
	

// ErrorResponse représente une erreur retournée par l'API
type ErrorResponse struct {
	Message string `json:"message"`
}

