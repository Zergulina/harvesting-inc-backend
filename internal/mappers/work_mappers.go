package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromWorkToDto(work *models.Work) *dto.WorkDto {
	workDto := new(dto.WorkDto)
	workDto.Id = work.Id
	workDto.StartDate = work.StartDate
	if work.EndDate.Valid {
		workDto.EndDate = &work.EndDate.Time
	}
	workDto.FieldId = work.FieldId
	return workDto
}

func FromCreateRequestDtoToWork(createDto *dto.CreateWorkRequestDto, fieldId uint64) *models.Work {
	work := new(models.Work)
	work.StartDate = createDto.StartDate
	work.FieldId = fieldId
	return work
}

func FromUpdateRequestDtoToWork(updateDto *dto.UpdateWorkRequestDto) *models.Work {
	work := new(models.Work)
	work.StartDate = updateDto.StartDate
	if updateDto.EndDate != nil {
		work.EndDate.Time = *updateDto.EndDate
		work.EndDate.Valid = true
	}
	return work
}
