package models

import (
	"database/sql"
	"time"
)

type Equipment struct {
	InvNumber        uint64
	EquipmentModelId uint64
	StatusId         uint64
	BuyDate          time.Time
	DrawDownDate     sql.NullTime
}
