package agendas

import (
	"database/sql"
	"middleware/example/internal/models"
)

// GetAllAgendas récupère tous les agendas
func GetAllAgendas(db *sql.DB) ([]models.Agenda, error) {
	rows, err := db.Query(`
		SELECT id, uca_id, name 
		FROM agendas
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda
	for rows.Next() {
		var agenda models.Agenda
		err := rows.Scan(
			&agenda.ID,
			&agenda.UcaID,
			&agenda.Name,
		)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, agenda)
	}

	return agendas, rows.Err()
}

// GetAgendaByID récupère un agenda par son ID
func GetAgendaByID(db *sql.DB, id string) (*models.Agenda, error) {
	var agenda models.Agenda
	err := db.QueryRow(`
		SELECT id, uca_id, name 
		FROM agendas 
		WHERE id = ?
	`, id).Scan(
		&agenda.ID,
		&agenda.UcaID,
		&agenda.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &agenda, nil
}

// CreateAgenda crée un nouvel agenda
func CreateAgenda(db *sql.DB, agenda *models.Agenda) error {
	_, err := db.Exec(`
		INSERT INTO agendas (id, uca_id, name)
		VALUES (?, ?, ?)
	`, agenda.ID, agenda.UcaID, agenda.Name)
	return err
}

// UpdateAgenda met à jour un agenda existant
func UpdateAgenda(db *sql.DB, agenda *models.Agenda) error {
	result, err := db.Exec(`
		UPDATE agendas 
		SET uca_id = ?, name = ?
		WHERE id = ?
	`, agenda.UcaID, agenda.Name, agenda.ID)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// DeleteAgenda supprime un agenda
func DeleteAgenda(db *sql.DB, id string) error {
	result, err := db.Exec(`DELETE FROM agendas WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
