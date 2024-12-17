package dto

type CreateMachineModelRequestDto struct {
	Name string `json:"name"`
}

type MachineModelDto struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	MachineTypeId   uint64 `json:"machine_type_id"`
	MachineTypeName string `json:"machine_type_name"`
}

type UpdateMachineModelRequestDto struct {
	Name string `json:"name"`
}
