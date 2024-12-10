package models

import (
	"database/sql"
	"time"
)

type Employee struct {
	PeopleId       uint64
	PostId         uint64
	EmploymentDate time.Time
	FireDate       sql.NullTime
	Salary         uint64
}
