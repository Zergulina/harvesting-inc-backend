package reports

import "time"

type FieldHarvestingReport struct {
	FieldId      uint64
	FieldCoords  string
	CropTypeName string
	Day          time.Time
	CropAmount   uint64
}
