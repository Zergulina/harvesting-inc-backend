package repository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllCropTypes(db *sql.DB) ([]models.CropType, error) {
	rows, err := db.Query("SELECT * FROM crop_types")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	crop_types := []models.CropType{}

	for rows.Next() {
		c := models.CropType{}
		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			continue
		}
		crop_types = append(crop_types, c)
	}

	return crop_types, nil
}

func GetCropTypeById(db *sql.DB, id uint64) (*models.CropType, error) {
	cropType := new(models.CropType)
	row := db.QueryRow("SELECT * FROM crop_types WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&cropType.Id, &cropType.Name)
	if err != nil {
		return nil, err
	}
	return cropType, nil
}

func CreateCropType(db *sql.DB, post *models.CropType) (*models.CropType, error) {
	row := db.QueryRow("INSERT INTO crop_types (name) VALUES ($1) RETURNING id", post.Name)
	err := row.Scan(&post.Id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func DeleteCropType(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM crop_types WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCropType(db *sql.DB, id uint64, post *models.CropType) (*models.CropType, error) {
	result, err := db.Exec("UPDATE crop_types SET name = $1 WHERE id = $2", post.Name, id)
	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ExistsCropType(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM crop_types WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
