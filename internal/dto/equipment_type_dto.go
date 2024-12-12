package dto

type CreateEquipmentTypeRequestDto struct {
	Name string `json:"name"`
}

type EquipmentTypeDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UpdateEquipmentTypeRequestDto struct {
	Name string `json:"name"`
}
