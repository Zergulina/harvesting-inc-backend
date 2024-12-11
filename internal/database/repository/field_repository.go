package repository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllFields(db *sql.DB) ([]models.Field, error) {
	rows, err := db.Query("SELECT * FROM fields")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fields := []models.Field{}

	for rows.Next() {
		f := models.Field{}
		err := rows.Scan(&f.Id, &f.Coords, &f.CustomerId, &f.CropId)
		if err != nil {
			continue
		}
		fields = append(fields, f)
	}

	return fields, nil
}

func GetAllFieldsByCustomerId(db *sql.DB, customerId uint64) ([]models.Field, error) {
	rows, err := db.Query("SELECT * FROM fields WHERE customer_id = $1", customerId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fields := []models.Field{}

	for rows.Next() {
		f := models.Field{}
		err := rows.Scan(&f.Id, &f.Coords, &f.CustomerId, &f.CropId)
		if err != nil {
			continue
		}
		fields = append(fields, f)
	}

	return fields, nil
}

func CreateField(db *sql.DB, field *models.Field) (*models.Field, error) {
	row := db.QueryRow("INSERT INTO fields (coords, customer_id, crop_id) VALUES ($1, $2, $3) RETURNING id", field.Coords, field.CustomerId, field.CropId)
	err := row.Scan(&field.Id)
	if err != nil {
		return nil, err
	}
	return field, nil
}

func DeleteField(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM fields WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateField(db *sql.DB, id uint64, field *models.Field) (*models.Field, error) {

	result, err := db.Exec("UPDATE fields SET coords = $1, crop_id = $2 WHERE id = $3", field.Coords, field.CropId, id)
	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}
	return field, nil
}

func ExistsField(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM fields WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
