package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromVacationToDto(vacation *models.Vacation) *dto.VacationDto {
	vacationDto := new(dto.VacationDto)
	vacationDto.PeopleId = vacation.PeopleId
	vacationDto.StartDate = vacation.StartDate
	vacationDto.EndDate = vacation.EndDate
	return vacationDto
}

func FromCreateRequestDtoToVacation(createDto *dto.CreateVacationRequestDto, peopleId uint64) *models.Vacation {
	vacation := new(models.Vacation)
	vacation.PeopleId = peopleId
	vacation.StartDate = createDto.StartDate
	vacation.EndDate = createDto.EndDate
	return vacation
}

func FromDeleteRequestDtoToVacation(deleteDto *dto.DeleteVacationRequestDto, peopleId uint64) *models.Vacation {
	vacation := new(models.Vacation)
	vacation.PeopleId = peopleId
	vacation.StartDate = deleteDto.StartDate
	return vacation
}
