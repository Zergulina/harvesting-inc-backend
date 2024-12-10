package models

import (
	"database/sql"
	"time"
)

type People struct {
	Id           uint64         `json:"id"`
	LastName     string         `json:"lastname"`
	FirstName    string         `json:"firstname"`
	MiddleName   sql.NullString `json:"middlename"`
	BirthDate    time.Time      `json:"birthdate"`
	Login        string         `json:"login"`
	PasswordHash string         `json:"password"`
}
