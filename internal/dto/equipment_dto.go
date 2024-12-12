package dto

import "time"

type CreateEquipmentRequestDto struct {
	StatusId uint64    `json:"status_id"`
	BuyDate  time.Time `json:"buy_date"`
}

type EquipmentDto struct {
	InvNumber        uint64     `json:"inv_number"`
	EquipmentModelId uint64     `json:"equipment_model_id"`
	StatusId         uint64     `json:"status_id"`
	BuyDate          time.Time  `json:"buy_date"`
	DrawDownDate     *time.Time `json:"draw_down_date"`
}

type UpdateEquipmentRequestDto struct {
	StatusId     uint64     `json:"status_id"`
	BuyDate      time.Time  `json:"buy_date"`
	DrawDownDate *time.Time `json:"draw_down_date"`
}
