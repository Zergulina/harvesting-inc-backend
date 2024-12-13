package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllEquipment(db *sql.DB) ([]models.Machine, error) {
	rows, err := db.Query("SELECT * FROM equipment")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	machines := []models.Machine{}

	for rows.Next() {
		m := models.Machine{}
		err := rows.Scan(&m.InvNumber, &m.MachineModelId, &m.StatusId, &m.BuyDate, &m.DrawDownDate)
		if err != nil {
			continue
		}
		machines = append(machines, m)
	}

	return machines, nil
}

func GetAllEquipmentsByEquipmentTypeId(db *sql.DB, equipmentTypeId uint64) ([]models.Equipment, error) {
	rows, err := db.Query("SELECT equipments.inv_number, equipments.machine_model_id, equipments.status_id, equipments.buy_date, equipments.draw_down_date FROM equipment LEFT JOIN equipment_models ON equipments.equipment_model_id = equipment_models.id WHERE equipment_models.equipment_type_id = $1", equipmentTypeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	equipment := []models.Equipment{}

	for rows.Next() {
		e := models.Equipment{}
		err := rows.Scan(&e.InvNumber, &e.EquipmentModelId, &e.StatusId, &e.BuyDate, &e.DrawDownDate)
		if err != nil {
			continue
		}
		equipment = append(equipment, e)
	}

	return equipment, nil
}

func GetAllEquipmentsByEquipmentModelId(db *sql.DB, equipmentModelId uint64) ([]models.Equipment, error) {
	rows, err := db.Query("SELECT * FROM equipments WHERE equipment_model_id = $1", equipmentModelId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	equipment := []models.Equipment{}

	for rows.Next() {
		e := models.Equipment{}
		err := rows.Scan(&e.InvNumber, &e.EquipmentModelId, &e.StatusId, &e.BuyDate, &e.DrawDownDate)
		if err != nil {
			continue
		}
		equipment = append(equipment, e)
	}

	return equipment, nil
}

func CreateEquipment(db *sql.DB, equipment *models.Equipment) (*models.Equipment, error) {
	row := db.QueryRow("SELECT inv_number FROM equipments WHERE equipment_model_id = $1 ORDER BY inv_number DESC LIMIT 1", equipment.InvNumber)
	var inv_number uint64
	err := row.Scan(&inv_number)
	if err != nil {
		inv_number = 1
	} else {
		inv_number++
	}

	_, err = db.Exec("INSERT INTO equipments (inv_number, equipment_model_id, status_id, buy_date) VALUES ($1, $2, $3, $4)", &equipment.InvNumber, &equipment.EquipmentModelId, &equipment.StatusId, &equipment.BuyDate)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}

func DeleteEquipment(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM equipments WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEquipment(db *sql.DB, equipmentModelId uint64, invNumber uint64, equipment *models.Equipment) (*models.Equipment, error) {
	result, err := db.Exec("UPDATE equipments SET status_id = $1, buy_date = $2, draw_down_date = $3 WHERE inv_number = $4 AND equipment_model_id = $5", equipment.StatusId, equipment.BuyDate, equipment.DrawDownDate, invNumber, equipmentModelId)
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

	return equipment, nil
}

func ExistsEquipment(db *sql.DB, equipmentModelId uint64, invNumber uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM equipments WHERE equipment_model_id = $1 AND inv_number = $2))", equipmentModelId, invNumber)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
