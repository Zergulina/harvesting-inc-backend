package models

import (
	"database/sql"
	"time"
)

type Machine struct {
	InvNumber      uint64
	MachineModelId uint64
	StatusId       uint64
	BuyDate        time.Time
	DrawDownDate   sql.NullTime
}
