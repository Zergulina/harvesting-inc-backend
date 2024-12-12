package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromMachineModelToDto(machineModel *models.MachineModel) *dto.MachineModelDto {
	machineModelDto := new(dto.MachineModelDto)
	machineModelDto.Id = machineModel.Id
	machineModelDto.Name = machineModel.Name
	machineModelDto.MachineTypeId = machineModel.MachineTypeId
	return machineModelDto
}

func FromCreateRequestDtoToMachineModel(createDto *dto.CreateMachineModelRequestDto, machineTypeid uint64) *models.MachineModel {
	machineModel := new(models.MachineModel)
	machineModel.Name = createDto.Name
	machineModel.MachineTypeId = machineTypeid
	return machineModel
}

func FromUpdateRequestDtoToMachineModel(updateDto *dto.UpdateMachineModelRequestDto) *models.MachineModel {
	machineModel := new(models.MachineModel)
	machineModel.Name = updateDto.Name
	return machineModel
}
