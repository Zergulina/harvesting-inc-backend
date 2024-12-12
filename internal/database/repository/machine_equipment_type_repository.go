package repository

import (
	"backend/internal/models"
	"database/sql"
)

func CreateMachineEquipmentType(db *sql.DB, machine_type_id uint64, equipment_type_id uint64) error {
	_, err := db.Exec("INSERT INTO machine_equipment_types (machine_type_id, equipment_type_id) VALUES($1, $2)", machine_type_id, equipment_type_id)

	return err
}

func DeleteMachineEquipmentType(db *sql.DB, machine_type_id uint64, equipment_type_id uint64) error {
	_, err := db.Exec("DELETE FROM machine_equipment_types WHERE machine_type_id = $1 AND equipment_type_id = $2", machine_type_id, equipment_type_id)

	return err
}

func GetAllMachineTypesByEquipmentTypeId(db *sql.DB, equipment_type_id uint64) ([]models.MachineType, error) {
	rows, err := db.Query("SELECT machine_types.id, machine_types.name FROM machine_types LEFT JOIN machine_equipment_types ON machine_types.id = machine_equipment_types.machine_type_id")
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

func GetAllEquipmentTypesByMachineTypeId(db *sql.DB, machine_type_id uint64) ([]models.EquipmentType, error) {
	rows, err := db.Query("SELECT equipment_types.id, equipment_types.name FROM machine_types LEFT JOIN machine_equipment_types ON equipment_types.id = machine_equipment_types.equipment_type_id")
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

func ExistsMachineEquipmentType(db *sql.DB, machine_type_id uint64, equipment_type_id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM machine_equipment_types WHERE machine_type_id = $1 AND equipment_type_id = $2))", machine_type_id, equipment_type_id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
