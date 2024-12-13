package dto

import "time"

type PeoplePostsDto struct {
	Id         uint64    `json:"id"`
	LastName   string    `json:"lastname"`
	FirstName  string    `json:"firstname"`
	MiddleName *string   `json:"middlename"`
	BirthDate  time.Time `json:"birthdate"`
	Posts      []string  `json:"posts"`
}
