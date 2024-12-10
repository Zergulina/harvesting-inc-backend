package models

import (
	"time"
)

type Vacation struct {
	PeopleId  uint64
	StartDate time.Time
	EndDate   time.Time
}
