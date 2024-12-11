package dto

type CreateFieldRequestDto struct {
	Coords     string `json:"coords"`
	CustomerId uint64 `json:"customer_id"`
	CropId     uint64 `json:"crop_id"`
}

type FieldDto struct {
	Id         uint64 `json:"id"`
	Coords     string `json:"coords"`
	CustomerId uint64 `json:"customer_id"`
	CropId     uint64 `json:"crop_id"`
}

type UpdateFieldRequestDto struct {
	Coords     string `json:"coords"`
	CustomerId uint64 `json:"customer_id"`
	CropId     uint64 `json:"crop_id"`
}
