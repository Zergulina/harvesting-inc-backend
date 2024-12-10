package models

import (
	"database/sql"
	"time"
)

type People struct {
	Id           uint64
	LastName     string
	FirstName    string
	MiddleName   sql.NullString
	BirthDate    time.Time
	Login        string
	PasswordHash string
}
