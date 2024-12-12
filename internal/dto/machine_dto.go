package dto

import "time"

type CreateMachineRequestDto struct {
	StatusId uint64    `json:"status_id"`
	BuyDate  time.Time `json:"buy_date"`
}

type MachineDto struct {
	InvNumber      uint64     `json:"inv_number"`
	MachineModelId uint64     `json:"machine_model_id"`
	StatusId       uint64     `json:"status_id"`
	BuyDate        time.Time  `json:"buy_date"`
	DrawDownDate   *time.Time `json:"draw_down_date"`
}

type UpdateMachineRequestDto struct {
	StatusId     uint64     `json:"status_id"`
	BuyDate      time.Time  `json:"buy_date"`
	DrawDownDate *time.Time `json:"draw_down_date"`
}
