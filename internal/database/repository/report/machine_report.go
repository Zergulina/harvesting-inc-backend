package report

import (
	"backend/internal/reports"
	"database/sql"
)

func GetMachineReport(db *sql.DB) ([]reports.MachineReport, error) {
	rows, err := db.Query(`SELECT machines.machine_model_id, machine_types.name, machines.inv_number, machine_models.name, statuses.name, machines.buy_date FROM machines JOIN machine_models ON machine_models.id = machines.machine_model_id JOIN statuses ON statuses.id = machines.status_id JOIN machine_types ON machine_types.id = machine_models.machine_type_id WHERE machines.draw_down_date IS NULL`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	mReports := []reports.MachineReport{}

	for rows.Next() {
		m := reports.MachineReport{}
		err := rows.Scan(&m.ModelId, &m.TypeName, &m.InvNumber, &m.ModelName, &m.Status, &m.BuyDate)
		if err != nil {
			continue
		}
		mReports = append(mReports, m)
	}

	return mReports, nil
}
