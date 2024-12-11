package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromMachineTypeToDto(machineType *models.MachineType) *dto.MachineTypeDto {
	machineTypeDto := new(dto.MachineTypeDto)
	machineTypeDto.Id = machineType.Id
	machineTypeDto.Name = machineType.Name
	return machineTypeDto
}

func FromCreateRequestDtoToMachineType(createRequestDto *dto.CreateMachineTypeRequestDto) *models.MachineType {
	machineType := new(models.MachineType)
	machineType.Name = createRequestDto.Name
	return machineType
}

func FromUpdateRequestDtoToMachineType(updateRequestDto *dto.UpdateMachineTypeRequestDto) *models.MachineType {
	machineType := new(models.MachineType)
	machineType.Name = updateRequestDto.Name
	return machineType
}
