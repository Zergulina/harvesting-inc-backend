package dto

type CreateCropTypeRequestDto struct {
	Name string `json:"name"`
}

type CropTypeDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UpdateCropTypeRequestDto struct {
	Name string `json:"name"`
}
