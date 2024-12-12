package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromEquipmentTypeToDto(machineType *models.EquipmentType) *dto.EquipmentTypeDto {
	equipmentTypeDto := new(dto.EquipmentTypeDto)
	equipmentTypeDto.Id = machineType.Id
	equipmentTypeDto.Name = machineType.Name
	return equipmentTypeDto
}

func FromCreateRequestDtoToEquipmentType(createRequestDto *dto.CreateEquipmentTypeRequestDto) *models.EquipmentType {
	equipmentType := new(models.EquipmentType)
	equipmentType.Name = createRequestDto.Name
	return equipmentType
}

func FromUpdateRequestDtoToEquipmentType(updateRequestDto *dto.UpdateEquipmentTypeRequestDto) *models.EquipmentType {
	equipmentType := new(models.EquipmentType)
	equipmentType.Name = updateRequestDto.Name
	return equipmentType
}
