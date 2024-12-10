package models

import (
	"database/sql"
)

type Crop struct {
	Id          uint64
	Name        string
	CropTypeId  uint64
	Description sql.NullString
}
