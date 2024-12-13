package dto

type CreateStatusRequestDto struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}

type StatusDto struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}

type UpdateStatusRequestDto struct {
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}
