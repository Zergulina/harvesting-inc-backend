package models

import "database/sql"

type Customer struct {
	Id            uint64
	Ogrn          string
	Name          string
	Logo          []byte
	LogoExtension sql.NullString
}
