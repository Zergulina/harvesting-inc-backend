package repository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllVacations(db *sql.DB) ([]models.Vacation, error) {
	rows, err := db.Query("SELECT * FROM vacations")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	vacations := []models.Vacation{}

	for rows.Next() {
		v := models.Vacation{}
		err := rows.Scan(&v.PeopleId, &v.StartDate, &v.EndDate)
		if err != nil {
			continue
		}
		vacations = append(vacations, v)
	}

	return vacations, nil
}

func GetAllVacationsByPeopleId(db *sql.DB, peopleId uint64) ([]models.Vacation, error) {
	rows, err := db.Query("SELECT * FROM vacations WHERE people_id = $1", peopleId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	vacations := []models.Vacation{}

	for rows.Next() {
		v := models.Vacation{}
		err := rows.Scan(&v.PeopleId, &v.StartDate, &v.EndDate)
		if err != nil {
			continue
		}
		vacations = append(vacations, v)
	}

	return vacations, nil
}

func CreateVacation(db *sql.DB, vacation *models.Vacation) (*models.Vacation, error) {
	_, err := db.Exec("INSERT INTO vacations (people_id, start_date) VALUES ($1, $2, $3) RETURNING id", vacation.PeopleId, vacation.StartDate, vacation.EndDate)
	if err != nil {
		return nil, err
	}
	return vacation, nil
}

func DeleteVacation(db *sql.DB, vacation *models.Vacation) error {
	_, err := db.Exec("DELETE FROM vacations WHERE people_id = $1, start_date = $2", vacation.PeopleId, vacation.StartDate)
	if err != nil {
		return err
	}
	return nil
}

func ExistsStatus(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM statuses WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
