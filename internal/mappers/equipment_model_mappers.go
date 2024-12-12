package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromEquipmentModelToDto(equipmentModel *models.EquipmentModel) *dto.EquipmentModelDto {
	equipmentModelDto := new(dto.EquipmentModelDto)
	equipmentModelDto.Id = equipmentModel.Id
	equipmentModelDto.Name = equipmentModel.Name
	equipmentModelDto.EquipmentTypeId = equipmentModel.EquipmentTypeId
	return equipmentModelDto
}

func FromCreateRequestDtoToEquipmentModel(createDto *dto.CreateEquipmentModelRequestDto, equipmentTypeid uint64) *models.EquipmentModel {
	equipmentModel := new(models.EquipmentModel)
	equipmentModel.Name = createDto.Name
	equipmentModel.EquipmentTypeId = equipmentTypeid
	return equipmentModel
}

func FromUpdateRequestDtoToEquipmentModel(updateDto *dto.UpdateEquipmentModelRequestDto) *models.EquipmentModel {
	equipmentModel := new(models.EquipmentModel)
	equipmentModel.Name = updateDto.Name
	return equipmentModel
}
