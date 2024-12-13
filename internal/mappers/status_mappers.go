package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromStatusToDto(status *models.Status) *dto.StatusDto {
	statusDto := new(dto.StatusDto)
	statusDto.Id = status.Id
	statusDto.Name = status.Name
	statusDto.IsAvailable = status.IsAvailable
	return statusDto
}

func FromCreateRequestDtoToStatus(createDto *dto.CreateStatusRequestDto) *models.Status {
	status := new(models.Status)
	status.Name = createDto.Name
	status.IsAvailable = createDto.IsAvailable
	return status
}

func FromUpdateRequestDtoToStatus(updateDto *dto.UpdateStatusRequestDto) *models.Status {
	status := new(models.Status)
	status.Name = updateDto.Name
	status.IsAvailable = updateDto.IsAvailable
	return status
}
