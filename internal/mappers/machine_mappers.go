package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromMachineToDto(machine *models.Machine) *dto.MachineDto {
	machineDto := new(dto.MachineDto)
	machineDto.InvNumber = machine.InvNumber
	machineDto.MachineModelId = machine.MachineModelId
	machineDto.StatusId = machine.MachineModelId
	machineDto.BuyDate = machine.BuyDate
	if machine.DrawDownDate.Valid {
		machineDto.DrawDownDate = &machine.DrawDownDate.Time
	}
	return machineDto
}

func FromCreateRequestDtoToMachine(createDto *dto.CreateMachineRequestDto, machineModelId uint64) *models.Machine {
	machine := new(models.Machine)
	machine.MachineModelId = machineModelId
	machine.StatusId = createDto.StatusId
	machine.BuyDate = createDto.BuyDate
	return machine
}

func FromUpdateRequestDtoToMachine(updateDto *dto.UpdateMachineRequestDto) *models.Machine {
	machine := new(models.Machine)
	machine.StatusId = updateDto.StatusId
	machine.BuyDate = updateDto.BuyDate
	if updateDto.DrawDownDate != nil {
		machine.DrawDownDate.Time = *updateDto.DrawDownDate
		machine.DrawDownDate.Valid = true
	}
	return machine
}
