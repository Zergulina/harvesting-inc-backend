package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromWorkTripToDto(workTrip *models.WorkTrip) *dto.WorkTripDto {
	workTripDto := new(dto.WorkTripDto)
	workTripDto.Id = workTrip.Id
	workTripDto.StartDateTime = workTrip.StartDateTime
	if workTrip.EndDateTime.Valid {
		workTripDto.EndDateTime = &workTrip.EndDateTime.Time
	}
	workTripDto.WorkId = workTrip.WorkId
	workTripDto.MachineInvNumber = workTrip.MachineInvNumber
	workTripDto.MachineModelId = workTrip.MachineModelId
	if workTrip.EquipmentInvNumber.Valid {
		*workTripDto.EquipmentInvNumber = uint64(workTrip.EquipmentInvNumber.Int64)
	}
	if workTrip.EquipmentModelId.Valid {
		*workTripDto.EquipmentModelId = uint64(workTrip.EquipmentModelId.Int64)
	}

	return workTripDto
}

func FromCreateRequestDtoToWorkTrip(createDto *dto.CreateWorkTripRequestDto, workId uint64) *models.WorkTrip {
	workTrip := new(models.WorkTrip)
	workTrip.StartDateTime = createDto.StartDateTime
	workTrip.MachineInvNumber = createDto.MachineInvNumber
	workTrip.MachineModelId = createDto.MachineModelId
	workTrip.WorkId = workId
	if createDto.EquipmentInvNumber != nil {
		workTrip.EquipmentInvNumber.Int64 = int64(*createDto.EquipmentInvNumber)
		workTrip.EquipmentInvNumber.Valid = true
	}
	if createDto.EquipmentModelId != nil {
		workTrip.EquipmentModelId.Int64 = int64(*createDto.EquipmentModelId)
		workTrip.EquipmentModelId.Valid = true
	}

	return workTrip
}

func FromUpdateRequestDtoToWorkTrip(updateDto *dto.UpdateWorkTripRequestDto) *models.WorkTrip {
	workTrip := new(models.WorkTrip)
	workTrip.StartDateTime = updateDto.StartDateTime
	if updateDto.EndDateTime != nil {
		workTrip.EndDateTime.Time = *updateDto.EndDateTime
		workTrip.EndDateTime.Valid = true
	}
	workTrip.CropAmount = updateDto.CropAmount

	return workTrip
}
