package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllStatuses(db *sql.DB) ([]models.Status, error) {
	rows, err := db.Query("SELECT * FROM statuses")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	statuses := []models.Status{}

	for rows.Next() {
		s := models.Status{}
		err := rows.Scan(&s.Id, &s.Name, &s.IsAvailable)
		if err != nil {
			continue
		}
		statuses = append(statuses, s)
	}

	return statuses, nil
}

func CreateStatus(db *sql.DB, status *models.Status) (*models.Status, error) {
	row := db.QueryRow("INSERT INTO statuses VALUES ($1, $2) RETURNING id", status.Name, status.IsAvailable)
	err := row.Scan(&status.Id)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func DeleteStatus(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM statuses WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatus(db *sql.DB, status *models.Status) (*models.Status, error) {

	result, err := db.Exec("UPDATE statuses SET name = $1, is_available = $2 WHERE id = $3", status.Name, status.IsAvailable, status.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("")
	}

	return status, nil
}
