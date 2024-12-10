package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllMachineModels(db *sql.DB) ([]models.MachineModel, error) {
	rows, err := db.Query("SELECT * FROM machine_models")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	machineModels := []models.MachineModel{}

	for rows.Next() {
		m := models.MachineModel{}
		err := rows.Scan(&m.Id, &m.Name, &m.MachineTypeId)
		if err != nil {
			continue
		}
		machineModels = append(machineModels, m)
	}

	return machineModels, nil
}

func GetAllMachineModelsByMachineTypeId(db *sql.DB, machineTypeId uint64) ([]models.MachineModel, error) {
	rows, err := db.Query("SELECT * FROM machine_models WHERE machine_type_id = $1", machineTypeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	machineModels := []models.MachineModel{}

	for rows.Next() {
		c := models.MachineModel{}
		err := rows.Scan(&c.Id, &c.Name, &c.MachineTypeId)
		if err != nil {
			continue
		}
		machineModels = append(machineModels, c)
	}

	return machineModels, nil
}

func CreateMachineModel(db *sql.DB, machineModel *models.MachineModel) (*models.MachineModel, error) {
	row := db.QueryRow("INSERT INTO machine_models VALUES ($1, $2) RETURNING id", machineModel.Name, machineModel.MachineTypeId)
	err := row.Scan(&machineModel.Id)
	if err != nil {
		return nil, err
	}
	return machineModel, nil
}

func DeleteMachineModel(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM machine_models WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMachineModel(db *sql.DB, machineModel *models.MachineModel) (*models.MachineModel, error) {

	result, err := db.Exec("UPDATE machine_models SET name = $1 WHERE id = $3", machineModel.Name, machineModel.Id)
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

	return machineModel, nil
}
