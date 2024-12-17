package reports

import "time"

type MachineReport struct {
	ModelId   uint64    `json:"model_id"`
	TypeName  string    `json:"type_name"`
	InvNumber uint64    `json:"inv_number"`
	ModelName string    `json:"model_name"`
	Status    string    `json:"status"`
	BuyDate   time.Time `json:"buy_date"`
}
