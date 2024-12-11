package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromCropTypeToDto(cropType *models.CropType) *dto.CropTypeDto {
	cropTypeDto := new(dto.CropTypeDto)
	cropTypeDto.Id = cropType.Id
	cropTypeDto.Name = cropType.Name
	return cropTypeDto
}

func FromCreateRequestDtoToCropType(createRequestDto *dto.CreateCropTypeRequestDto) *models.CropType {
	cropType := new(models.CropType)
	cropType.Name = createRequestDto.Name
	return cropType
}

func FromUpdateReqeustDtoToCropType(updateRequestDto *dto.UpdateCropTypeRequestDto) *models.CropType {
	cropType := new(models.CropType)
	cropType.Name = updateRequestDto.Name
	return cropType
}
