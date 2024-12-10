package models

import (
	"database/sql"
	"time"
)

type Work struct {
	Id        uint64
	StartDate time.Time
	EndDate   sql.NullTime
	FieldId   uint64
}
