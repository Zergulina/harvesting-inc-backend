package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromFieldToDto(field *models.Field) *dto.FieldDto {
	fieldDto := new(dto.FieldDto)
	fieldDto.Id = field.Id
	fieldDto.Coords = field.Coords
	fieldDto.CustomerId = field.CustomerId
	fieldDto.CropId = field.CropId
	return fieldDto
}

func FromCreateRequestToField(createDto *dto.CreateFieldRequestDto, customerId uint64) *models.Field {
	field := new(models.Field)
	field.Coords = createDto.Coords
	field.CustomerId = customerId
	field.CropId = createDto.CropId
	return field
}

func FromUpdateRequestToField(updateDto *dto.UpdateFieldRequestDto, customerId uint64) *models.Field {
	field := new(models.Field)
	field.Coords = updateDto.Coords
	field.CustomerId = customerId
	field.CropId = updateDto.CropId
	return field
}
