package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"encoding/base64"
)

func FromCustomerToDto(customer *models.Customer) *dto.CustomerDto {
	customerDto := new(dto.CustomerDto)
	customerDto.Id = customer.Id
	customerDto.Name = customer.Name
	customerDto.Ogrn = customer.Ogrn
	if customer.Logo != nil {
		*customerDto.Logo = base64.StdEncoding.EncodeToString(customer.Logo)
	}
	if customer.LogoExtension.Valid {
		customerDto.LogoExtension = &customer.Name
	}
	return customerDto
}

func FromCreateRequestDtoToCustomer(createRequestDto *dto.CreateCustomerRequestDto) *models.Customer {
	customer := new(models.Customer)
	customer.Ogrn = createRequestDto.Ogrn
	customer.Name = createRequestDto.Name
	if createRequestDto.Logo != nil {
		logo, err := base64.StdEncoding.DecodeString(*createRequestDto.Logo)
		if err != nil {
			customer.Logo = logo
		}
	}
	if createRequestDto.LogoExtension != nil {
		customer.LogoExtension.String = *createRequestDto.LogoExtension
		customer.LogoExtension.Valid = true
	}
	return customer
}

func FromUpdateRequestDtoToCustomer(updateRequestDto *dto.UpdateCustomerRequestDto) *models.Customer {
	customer := new(models.Customer)
	customer.Ogrn = updateRequestDto.Ogrn
	customer.Name = updateRequestDto.Name
	if updateRequestDto.Logo != nil {
		logo, err := base64.StdEncoding.DecodeString(*updateRequestDto.Logo)
		if err != nil {
			customer.Logo = logo
		}
	}
	if updateRequestDto.LogoExtension != nil {
		customer.LogoExtension.String = *updateRequestDto.LogoExtension
		customer.LogoExtension.Valid = true
	}
	return customer
}

func FromPatchRequestDtoToCustomer(patchRequestDto *dto.PatchCustomerRequestDto) *models.Customer {
	customer := new(models.Customer)
	customer.Ogrn = patchRequestDto.Ogrn
	customer.Name = patchRequestDto.Name

	return customer
}
