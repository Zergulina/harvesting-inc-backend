package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllMachineTypes(db *sql.DB) ([]models.MachineType, error) {
	rows, err := db.Query("SELECT * FROM machine_types")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	machine_types := []models.MachineType{}

	for rows.Next() {
		m := models.MachineType{}
		err := rows.Scan(&m.Id, &m.Name)
		if err != nil {
			continue
		}
		machine_types = append(machine_types, m)
	}

	return machine_types, nil
}

func GetMachineTypeById(db *sql.DB, id uint64) (*models.MachineType, error) {
	machineType := new(models.MachineType)
	row := db.QueryRow("SELECT * FROM machine_types WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&machineType.Id, &machineType.Name)
	if err != nil {
		return nil, err
	}
	return machineType, nil
}

func CreateMachineType(db *sql.DB, machineType *models.MachineType) (*models.MachineType, error) {
	_, err := db.Exec("INSERT INTO machine_types (name) VALUES ($1) RETURNING id", machineType.Name)
	if err != nil {
		return nil, err
	}
	return machineType, nil
}

func DeleteMachineType(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM machine_types WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMachineType(db *sql.DB, id uint64, machineType *models.MachineType) (*models.MachineType, error) {

	result, err := db.Exec("UPDATE machine_types SET name = $1 WHERE id = $2", machineType.Name, id)
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

	return machineType, nil
}

func ExistsMachineType(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM machine_types WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
