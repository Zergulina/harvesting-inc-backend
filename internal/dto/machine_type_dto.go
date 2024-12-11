package dto

type CreateMachineTypeRequestDto struct {
	Name string `json:"name"`
}

type MachineTypeDto struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UpdateMachineTypeRequestDto struct {
	Name string `json:"name"`
}
