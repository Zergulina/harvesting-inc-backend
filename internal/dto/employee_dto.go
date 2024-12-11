package dto

import "time"

type CreateEmployeeRequestDto struct {
	PostId         uint64    `json:"post_id"`
	EmploymentDate time.Time `json:"employment_date"`
	Salary         uint64    `json:"salary"`
}

type EmployeeDto struct {
	PeopleId       uint64     `json:"people_id"`
	PostId         uint64     `json:"post_id"`
	EmploymentDate time.Time  `json:"employment_date"`
	FireDate       *time.Time `json:"fire_date"`
	Salary         uint64     `json:"salary"`
}

type UpdateEmployeeRequestDto struct {
	EmploymentDate time.Time  `json:"employment_date"`
	FireDate       *time.Time `json:"fire_date"`
	Salary         uint64     `json:"salary"`
}
