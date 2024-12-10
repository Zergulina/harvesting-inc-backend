package models

import (
	"database/sql"
)

type Customer struct {
	Id   uint64
	Ogrn string
	Name string
	Logo sql.NullByte
}
