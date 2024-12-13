package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromEquipmentToDto(equipments *models.Equipment) *dto.EquipmentDto {
	equipmentDto := new(dto.EquipmentDto)
	equipmentDto.InvNumber = equipments.InvNumber
	equipmentDto.EquipmentModelId = equipments.EquipmentModelId
	equipmentDto.StatusId = equipments.EquipmentModelId
	equipmentDto.BuyDate = equipments.BuyDate
	if equipments.DrawDownDate.Valid {
		equipmentDto.DrawDownDate = &equipments.DrawDownDate.Time
	}
	return equipmentDto
}

func FromCreateRequestDtoToEquipment(createDto *dto.CreateEquipmentRequestDto, equipmentModelId uint64) *models.Equipment {
	equipments := new(models.Equipment)
	equipments.EquipmentModelId = equipmentModelId
	equipments.StatusId = createDto.StatusId
	equipments.BuyDate = createDto.BuyDate
	return equipments
}

func FromUpdateRequestDtoToEquipment(updateDto *dto.UpdateEquipmentRequestDto) *models.Equipment {
	equipments := new(models.Equipment)
	equipments.StatusId = updateDto.StatusId
	equipments.BuyDate = updateDto.BuyDate
	if updateDto.DrawDownDate != nil {
		equipments.DrawDownDate.Time = *updateDto.DrawDownDate
		equipments.DrawDownDate.Valid = true
	}
	return equipments
}
