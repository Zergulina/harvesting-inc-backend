package repository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllCrops(db *sql.DB) ([]models.Crop, error) {
	rows, err := db.Query("SELECT * FROM crops")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	crops := []models.Crop{}

	for rows.Next() {
		c := models.Crop{}
		err := rows.Scan(&c.Id, &c.Name, &c.CropTypeId, &c.Description)
		if err != nil {
			continue
		}
		crops = append(crops, c)
	}

	return crops, nil
}

func GetAllCropsByCropTypeId(db *sql.DB, cropTypeId uint64) ([]models.Crop, error) {
	rows, err := db.Query("SELECT * FROM crops WHERE crop_type_id = $1", cropTypeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	crops := []models.Crop{}

	for rows.Next() {
		c := models.Crop{}
		err := rows.Scan(&c.Id, &c.Name, &c.CropTypeId, &c.Description)
		if err != nil {
			continue
		}
		crops = append(crops, c)
	}

	return crops, nil
}

func CreateCrop(db *sql.DB, crop *models.Crop) (*models.Crop, error) {
	row := db.QueryRow("INSERT INTO crops (name, crop_type_id, description) VALUES ($1, $2, $3) RETURNING id", crop.Name, crop.CropTypeId, crop.Description)
	err := row.Scan(&crop.Id)
	if err != nil {
		return nil, err
	}
	return crop, nil
}

func DeleteCrop(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM crops WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCrop(db *sql.DB, id uint64, crop *models.Crop) (*models.Crop, error) {

	result, err := db.Exec("UPDATE crops SET name = $1, description = $2 WHERE id = $3", crop.Name, crop.Description, id)
	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}
	return crop, nil
}

func ExistsCrop(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM crops WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
