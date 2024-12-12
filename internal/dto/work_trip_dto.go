package dto

import "time"

type CreateWorkTripRequestDto struct {
	StartDateTime      time.Time `json:"start_date_time"`
	MachineInvNumber   uint64    `json:"machine_inv_number"`
	MachineModelId     uint64    `json:"machine_model_id"`
	EquipmentInvNumber *uint64   `json:"equipment_inv_number"`
	EquipmentModelId   *uint64   `json:"equipment_model_id"`
}

type WorkTripDto struct {
	Id                 uint64     `json:"id"`
	StartDateTime      time.Time  `json:"start_date_time"`
	EndDateTime        *time.Time `json:"end_date_time"`
	CropAmount         uint64     `json:"crop_amount"`
	WorkId             uint64     `json:"work_id"`
	MachineInvNumber   uint64     `json:"machine_inv_number"`
	MachineModelId     uint64     `json:"machine_model_id"`
	EquipmentInvNumber *uint64    `json:"equipment_inv_number"`
	EquipmentModelId   *uint64    `json:"equipment_model_id"`
}

type UpdateWorkTripRequestDto struct {
	StartDateTime time.Time  `json:"start_date"`
	EndDateTime   *time.Time `json:"end_date_time"`
	CropAmount    uint64     `json:"crop_amount"`
}
