package report

import (
	"backend/internal/reports"
	"database/sql"
	"time"
)

func GetFieldHarvestingReport(db *sql.DB, start_period time.Time, end_period time.Time, customer_id uint64) ([]reports.FieldHarvestingReport, error) {
	rows, err := db.Query(`WITH
							harvested_on_field_at_day AS (SELECT works.field_id AS field_id, work_trips.start_date::DATE AS day, SUM(work_trips.crop_amount) AS crop_amount
								FROM (work_trips JOIN works ON work_trips.work_id = works.id) GROUP BY(work_trips.start_date::DATE, works.field_id)),
							customer_field AS (SELECT fields.id AS field_id, customers.id AS customer_id, fields.coords AS field_coords, customers.name AS customer_name, crop_types.name AS crop_type_name
								FROM 
									customers
										JOIN fields ON customers.id = fields.customer_id
										JOIN crops ON crops.id = fields.crop_id
										JOIN crop_types ON crop_types.id = crops.crop_type_id)
							SELECT customer_field.field_id, customer_field.field_coords, customer_field.crop_type_name, harvested_on_field_at_day.day, harvested_on_field_at_day.crop_amount
								FROM customer_field JOIN harvested_on_field_at_day JOIN ON customer_field.field_id = harvested_on_field_at_day.field_id
									WHERE harvested_on_field_at_day.day >= $1 AND harvested_on_field_at_day.day <= $2 AND customer_field.customer_id = $3`, start_period, end_period, customer_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	crop_types := []reports.FieldHarvestingReport{}

	for rows.Next() {
		f := reports.FieldHarvestingReport{}
		err := rows.Scan(&f.FieldId, &f.FieldCoords, &f.CropTypeName, &f.Day, &f.CropAmount)
		if err != nil {
			continue
		}
		crop_types = append(crop_types, f)
	}

	return crop_types, nil
}
