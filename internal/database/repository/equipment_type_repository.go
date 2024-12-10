package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllEquipmentTypes(db *sql.DB) ([]models.EquipmentType, error) {
	rows, err := db.Query("SELECT * FROM equipment_types")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	equipment_types := []models.EquipmentType{}

	for rows.Next() {
		e := models.EquipmentType{}
		err := rows.Scan(&e.Id, &e.Name)
		if err != nil {
			continue
		}
		equipment_types = append(equipment_types, e)
	}

	return equipment_types, nil
}

func CreateEquipmentType(db *sql.DB, equipmentType *models.EquipmentType) (*models.EquipmentType, error) {
	_, err := db.Exec("INSERT INTO equipment_types VALUES ($1) RETURNING id", equipmentType.Name)
	if err != nil {
		return nil, err
	}
	return equipmentType, nil
}

func DeleteEquipmentType(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM equipment_types WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEquipmentType(db *sql.DB, equipmentType *models.EquipmentType) (*models.EquipmentType, error) {

	result, err := db.Exec("UPDATE equipment_types SET name = $1 WHERE id = $2", equipmentType.Name, equipmentType.Id)
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

	return equipmentType, nil
}
