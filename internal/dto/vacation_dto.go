package dto

import "time"

type CreateVacationRequestDto struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type VacationDto struct {
	PeopleId  uint64    `json:"people_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type DeleteVacationRequestDto struct {
	StartDate time.Time `json:"start_date"`
}
