package mappers

import (
	"backend/internal/config"
	"backend/internal/dto"
	"backend/internal/helpers"
	"backend/internal/models"
)

func FromPeopleToDto(people *models.People) *dto.PeopleDto {
	peopleDto := new(dto.PeopleDto)
	peopleDto.Id = people.Id
	peopleDto.LastName = people.LastName
	peopleDto.FirstName = people.FirstName
	if people.MiddleName.Valid {
		peopleDto.MiddleName = &people.MiddleName.String
	}
	peopleDto.BirthDate = people.BirthDate
	return peopleDto
}

func FromRegisterDtoToPeople(registerRequestDto *dto.RegisterRequestDto) *models.People {
	people := new(models.People)
	people.FirstName = registerRequestDto.FirstName
	people.LastName = registerRequestDto.LastName
	if registerRequestDto.MiddleName != nil {
		people.MiddleName.String = *registerRequestDto.MiddleName
		people.MiddleName.Valid = true
	}
	people.BirthDate = registerRequestDto.BirthDate
	people.Login = registerRequestDto.Login
	people.PasswordHash = helpers.EncodeSha256(registerRequestDto.Password, config.DbSecretKey)
	return people
}

func FromUpdatePeopleRequestDtoToPeople(updateRequestDto *dto.UpdatePeopleRequestDto) *models.People {
	people := new(models.People)
	people.FirstName = updateRequestDto.FirstName
	people.LastName = updateRequestDto.LastName
	if updateRequestDto.MiddleName != nil {
		people.MiddleName.String = *updateRequestDto.MiddleName
		people.MiddleName.Valid = true
	}
	people.BirthDate = updateRequestDto.BirthDate
	return people
}
