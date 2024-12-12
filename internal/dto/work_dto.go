package dto

import "time"

type CreateWorkRequestDto struct {
	StartDate time.Time `json:"start_date"`
}

type WorkDto struct {
	Id        uint64     `json:"id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	FieldId   uint64     `json:"field_id"`
}

type UpdateWorkRequestDto struct {
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}
