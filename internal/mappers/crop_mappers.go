package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromCropToDto(crop *models.Crop) *dto.CropDto {
	cropDto := new(dto.CropDto)
	cropDto.Id = crop.Id
	cropDto.Name = crop.Name
	cropDto.CropTypeId = crop.CropTypeId
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
	}
	return crop
}

func FromUpdateRequestDtoToCrop(updateRequestDto *dto.UpdateCropRequestDto) *models.Crop {
	crop := new(models.Crop)
	crop.Name = updateRequestDto.Name
	if updateRequestDto.Description != nil {
		crop.Description.String = *updateRequestDto.Description
	}
	return crop
}
