package alerts

import (
	"database/sql"
	"middleware/example/internal/models"
)

// GetAllAlerts récupère toutes les alertes
func GetAllAlerts(db *sql.DB) ([]models.Alert, error) {
	rows, err := db.Query(`
		SELECT id, email, agenda_id 
		FROM alerts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var alert models.Alert
		err := rows.Scan(
			&alert.ID,
			&alert.Email,
			&alert.AgendaID,
		)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, rows.Err()
}

// GetAlertByID récupère une alerte par son ID
func GetAlertByID(db *sql.DB, id string) (*models.Alert, error) {
	var alert models.Alert
	err := db.QueryRow(`
		SELECT id, email, agenda_id 
		FROM alerts 
		WHERE id = ?
	`, id).Scan(
		&alert.ID,
		&alert.Email,
		&alert.AgendaID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

// CreateAlert crée une nouvelle alerte
func CreateAlert(db *sql.DB, alert *models.Alert) error {
	_, err := db.Exec(`
		INSERT INTO alerts (id, email, agenda_id)
		VALUES (?, ?, ?)
	`, alert.ID, alert.Email, alert.AgendaID)
	return err
}

// UpdateAlert met à jour une alerte existante
func UpdateAlert(db *sql.DB, alert *models.Alert) error {
	result, err := db.Exec(`
		UPDATE alerts 
		SET email = ?, agenda_id = ?
		WHERE id = ?
	`, alert.Email, alert.AgendaID, alert.ID)

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

// DeleteAlert supprime une alerte
func DeleteAlert(db *sql.DB, id string) error {
	result, err := db.Exec(`DELETE FROM alerts WHERE id = ?`, id)
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