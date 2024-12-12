package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllMachines(db *sql.DB) ([]models.Machine, error) {
	rows, err := db.Query("SELECT * FROM machines")
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

func GetAllMachinesByMachineTypeId(db *sql.DB, machineTypeId uint64) ([]models.Machine, error) {
	rows, err := db.Query("SELECT machines.inv_number, machines.machine_model_id, machines.status_id, machines.buy_date, machines.draw_down_date FROM machines LEFT JOIN machine_models ON machines.machine_model_id = machine_models.id WHERE machine_models.machine_type_id = $1", machineTypeId)
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

func GetAllMachinesByMachineModelId(db *sql.DB, machineModelId uint64) ([]models.Machine, error) {
	rows, err := db.Query("SELECT * FROM machines WHERE machine_model_id = $1", machineModelId)
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

func CreateMachine(db *sql.DB, machine *models.Machine) (*models.Machine, error) {
	row := db.QueryRow("SELECT inv_number FROM machines WHERE machine_model_id = $1 ORDER BY inv_number DESC LIMIT 1", machine.InvNumber)
	var inv_number uint64
	err := row.Scan(&inv_number)
	if err != nil {
		inv_number = 1
	} else {
		inv_number++
	}

	_, err = db.Exec("INSERT INTO machines (inv_number, machine_model_id, status_id, buy_date) VALUES ($1, $2, $3, $4)", &machine.InvNumber, &machine.MachineModelId, &machine.StatusId, &machine.BuyDate)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func DeleteMachine(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM machines WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMachine(db *sql.DB, machineModelId uint64, invNumber uint64, machine *models.Machine) (*models.Machine, error) {

	result, err := db.Exec("UPDATE machines SET status_id = $1, buy_date = $2, draw_down_date = $3 WHERE inv_number = $4 AND machine_model_id = $5", machine.StatusId, machine.BuyDate, machine.DrawDownDate, invNumber, machineModelId)
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

	return machine, nil
}

func ExistsMachine(db *sql.DB, machineModelId uint64, invNumber uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM machines WHERE machine_model_id = $1 AND inv_number = $2))", machineModelId, invNumber)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
