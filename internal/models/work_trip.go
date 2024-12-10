package models

import (
	"database/sql"
	"time"
)

type WorkTrip struct {
	Id                 uint64
	StartDateTime      time.Time
	EndDateTime        sql.NullTime
	CropAmount         uint64
	WorkId             uint64
	MachineInvNumber   uint64
	MachineModelId     uint64
	EquipmentInvNumber sql.NullInt64
	EquipmentModelId   sql.NullInt64
}
