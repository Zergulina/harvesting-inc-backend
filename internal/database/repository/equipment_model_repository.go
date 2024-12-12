package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllEquipmentModels(db *sql.DB) ([]models.EquipmentModel, error) {
	rows, err := db.Query("SELECT * FROM equipment_models")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	equipmentModels := []models.EquipmentModel{}

	for rows.Next() {
		e := models.EquipmentModel{}
		err := rows.Scan(&e.Id, &e.Name, &e.EquipmentTypeId)
		if err != nil {
			continue
		}
		equipmentModels = append(equipmentModels, e)
	}

	return equipmentModels, nil
}

func GetAllEquipmentModelsByEquipmentTypeId(db *sql.DB, equipmentTypeId uint64) ([]models.EquipmentModel, error) {
	rows, err := db.Query("SELECT * FROM equipment_models WHERE equipment_type_id = $1", equipmentTypeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	machineModels := []models.EquipmentModel{}

	for rows.Next() {
		c := models.EquipmentModel{}
		err := rows.Scan(&c.Id, &c.Name, &c.EquipmentTypeId)
		if err != nil {
			continue
		}
		machineModels = append(machineModels, c)
	}

	return machineModels, nil
}

func CreateEquipmentModel(db *sql.DB, equipmentModel *models.EquipmentModel) (*models.EquipmentModel, error) {
	row := db.QueryRow("INSERT INTO equipment_models (name, equipment_type_id) VALUES ($1, $2) RETURNING id", equipmentModel.Name, equipmentModel.EquipmentTypeId)
	err := row.Scan(&equipmentModel.Id)
	if err != nil {
		return nil, err
	}
	return equipmentModel, nil
}

func DeleteEquipmentModel(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM equipment_models WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEquipmentModel(db *sql.DB, id uint64, equipmentModel *models.EquipmentModel) (*models.EquipmentModel, error) {

	result, err := db.Exec("UPDATE equipment_models SET name = $1 WHERE id = $2", equipmentModel.Name, id)
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

	return equipmentModel, nil
}

func ExistsEquipmentModel(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM equipment_models WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
