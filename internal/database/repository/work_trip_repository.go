package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllWorkTrips(db *sql.DB) ([]models.WorkTrip, error) {
	rows, err := db.Query("SELECT * FROM work_trips")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	workTrips := []models.WorkTrip{}

	for rows.Next() {
		w := models.WorkTrip{}
		err := rows.Scan(&w.Id, &w.StartDateTime, &w.EndDateTime, &w.CropAmount, &w.MachineInvNumber, &w.MachineModelId, &w.EquipmentInvNumber, &w.EquipmentModelId)
		if err != nil {
			continue
		}
		workTrips = append(workTrips, w)
	}

	return workTrips, nil
}

func GetAllWorkTripsByWorkId(db *sql.DB, work_id uint64) ([]models.WorkTrip, error) {
	rows, err := db.Query("SELECT * FROM work_trips WHERE work_id = $1", work_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	workTrips := []models.WorkTrip{}

	for rows.Next() {
		w := models.WorkTrip{}
		err := rows.Scan(&w.Id, &w.StartDateTime, &w.EndDateTime, &w.CropAmount, &w.MachineInvNumber, &w.MachineModelId, &w.EquipmentInvNumber, &w.EquipmentModelId)
		if err != nil {
			continue
		}
		workTrips = append(workTrips, w)
	}

	return workTrips, nil
}

func CreateWorkTrip(db *sql.DB, workTrip *models.WorkTrip) (*models.WorkTrip, error) {
	row := db.QueryRow("INSERT INTO work_trips (start_date_time, end_date_time, crop_amount, work_id, machine_inv_number, machine_model_id, equipment_inv_number, equipment_model_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", workTrip.StartDateTime, workTrip.EndDateTime, workTrip.CropAmount, workTrip.MachineInvNumber, workTrip.MachineModelId, workTrip.EquipmentInvNumber, workTrip.EquipmentModelId)
	err := row.Scan(&workTrip.Id)
	if err != nil {
		return nil, err
	}
	return workTrip, nil
}

func DeleteWorkTrip(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM work_trips WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWorkTrip(db *sql.DB, worktrip *models.WorkTrip) (*models.WorkTrip, error) {

	result, err := db.Exec("UPDATE work_trips SET start_date_time = $1, end_date_time = $2, crop_amount = $3, machine_inv_number = $4, machine_model_id = $5, equipment_inv_number = $6, equipment_model_id = $7 WHERE id = $8", worktrip.StartDateTime, worktrip.EndDateTime, worktrip.CropAmount, worktrip.MachineInvNumber, worktrip.MachineModelId, worktrip.EquipmentInvNumber, worktrip.MachineModelId, worktrip.Id)
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

	return worktrip, nil
}
