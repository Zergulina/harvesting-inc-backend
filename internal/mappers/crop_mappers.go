package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromCropToDto(crop *models.Crop, cropType *models.CropType) *dto.CropDto {
	cropDto := new(dto.CropDto)
	cropDto.Id = crop.Id
	cropDto.Name = crop.Name
	cropDto.CropTypeId = cropType.Id
	cropDto.CropTypeName = cropType.Name
	if crop.Description.Valid {
		cropDto.Description = &crop.Description.String
	}
	return cropDto
}

func FromCreateRequestDtoToCrop(createRequestDto *dto.CreateCropRequestDto, cropTypeId uint64) *models.Crop {
	crop := new(models.Crop)
	crop.Name = createRequestDto.Name
	crop.CropTypeId = cropTypeId
	if createRequestDto.Description != nil {
		crop.Description.String = *createRequestDto.Description
		crop.Description.Valid = true
	}
	return crop
}

func FromUpdateRequestDtoToCrop(updateRequestDto *dto.UpdateCropRequestDto) *models.Crop {
	crop := new(models.Crop)
	crop.Name = updateRequestDto.Name
	if updateRequestDto.Description != nil {
		crop.Description.String = *updateRequestDto.Description
		crop.Description.Valid = true
	}
	return crop
}
