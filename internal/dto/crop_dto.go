package dto

type CreateCropRequestDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CropDto struct {
	Id           uint64  `json:"id"`
	Name         string  `json:"name"`
	CropTypeId   uint64  `json:"crop_type_id"`
	CropTypeName string  `json:"crop_type_name"`
	Description  *string `json:"description"`
}

type UpdateCropRequestDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
