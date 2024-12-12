package dto

type CreateEquipmentModelRequestDto struct {
	Name string `json:"name"`
}

type EquipmentModelDto struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	EquipmentTypeId uint64 `json:"equipment_type_id"`
}

type UpdateEquipmentModelRequestDto struct {
	Name string `json:"name"`
}
