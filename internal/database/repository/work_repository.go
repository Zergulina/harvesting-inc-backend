package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllWorks(db *sql.DB) ([]models.Work, error) {
	rows, err := db.Query("SELECT * FROM works")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	works := []models.Work{}

	for rows.Next() {
		w := models.Work{}
		err := rows.Scan(&w.Id, &w.StartDate, &w.EndDate, &w.FieldId)
		if err != nil {
			continue
		}
		works = append(works, w)
	}

	return works, nil
}

func GetAllWorksByFieldId(db *sql.DB, field_id uint64) ([]models.Work, error) {
	rows, err := db.Query("SELECT * FROM works WHERE field_id = $1", field_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	works := []models.Work{}

	for rows.Next() {
		w := models.Work{}
		err := rows.Scan(&w.Id, &w.StartDate, &w.EndDate, &w.FieldId)
		if err != nil {
			continue
		}
		works = append(works, w)
	}

	return works, nil
}

func CreateWork(db *sql.DB, work *models.Work) (*models.Work, error) {
	row := db.QueryRow("INSERT INTO works VALUES ($1, $2, $3) RETURNING id", work.StartDate, work.EndDate, work.FieldId)
	err := row.Scan(&work.Id)
	if err != nil {
		return nil, err
	}
	return work, nil
}

func DeleteWork(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM works WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWork(db *sql.DB, work *models.Work) (*models.Work, error) {

	result, err := db.Exec("UPDATE statuses SET start_date = $1, end_date = $2 WHERE id = $3", work.StartDate, work.EndDate, work.Id)
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

	return work, nil
}
