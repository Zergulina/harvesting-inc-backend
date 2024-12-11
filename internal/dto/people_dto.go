package dto

import (
	"time"
)

type RegisterRequestDto struct {
	LastName   string    `json:"lastname"`
	FirstName  string    `json:"firstname"`
	MiddleName *string   `json:"middlename"`
	BirthDate  time.Time `json:"birthdate"`
	Login      string    `json:"login"`
	Password   string    `jsom:"password"`
}

type LoginRequestDto struct {
	Login    string `json:"login"`
	Password string `jsom:"password"`
}

type PeopleDto struct {
	Id         uint64    `json:"id"`
	LastName   string    `json:"lastname"`
	FirstName  string    `json:"firstname"`
	MiddleName *string   `json:"middlename"`
	BirthDate  time.Time `json:"birthdate"`
}

type UpdatePeopleRequestDto struct {
	LastName   string    `json:"lastname"`
	FirstName  string    `json:"firstname"`
	MiddleName *string   `json:"middlename"`
	BirthDate  time.Time `json:"birthdate"`
}
