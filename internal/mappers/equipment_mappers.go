package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromEquipmentToDto(equipment *models.Equipment) *dto.EquipmentDto {
	equipmentDto := new(dto.EquipmentDto)
	equipmentDto.InvNumber = equipment.InvNumber
	equipmentDto.EquipmentModelId = equipment.EquipmentModelId
	equipmentDto.StatusId = equipment.EquipmentModelId
	equipmentDto.BuyDate = equipment.BuyDate
	if equipment.DrawDownDate.Valid {
		equipmentDto.DrawDownDate = &equipment.DrawDownDate.Time
	}
	return equipmentDto
}

func FromCreateRequestDtoToEquipment(createDto *dto.CreateEquipmentRequestDto, equipmentModelId uint64) *models.Equipment {
	equipment := new(models.Equipment)
	equipment.EquipmentModelId = equipmentModelId
	equipment.StatusId = createDto.StatusId
	equipment.BuyDate = createDto.BuyDate
	return equipment
}

func FromUpdateRequestDtoToEquipment(updateDto *dto.UpdateEquipmentRequestDto) *models.Equipment {
	equipment := new(models.Equipment)
	equipment.StatusId = updateDto.StatusId
	equipment.BuyDate = updateDto.BuyDate
	if updateDto.DrawDownDate != nil {
		equipment.DrawDownDate.Time = *updateDto.DrawDownDate
		equipment.DrawDownDate.Valid = true
	}
	return equipment
}
