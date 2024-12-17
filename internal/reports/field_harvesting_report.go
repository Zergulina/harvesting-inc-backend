package reports

import "time"

type FieldHarvestingReport struct {
	FieldId      uint64    `json:"field_id"`
	FieldCoords  string    `json:"field_coords"`
	CropTypeName string    `json:"crop_type_name"`
	Day          time.Time `json:"day"`
	CropAmount   uint64    `json:"crop_amount"`
}
