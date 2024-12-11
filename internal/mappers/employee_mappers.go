package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
)

func FromEmployeeToDto(employee *models.Employee) *dto.EmployeeDto {
	employeeDto := new(dto.EmployeeDto)
	employeeDto.PeopleId = employee.PeopleId
	employeeDto.PostId = employee.PostId
	employeeDto.EmploymentDate = employee.EmploymentDate
	if employee.FireDate.Valid {
		employeeDto.FireDate = &employee.FireDate.Time
	}
	employeeDto.Salary = employee.Salary
	return employeeDto
}

func FromCreateRequestDtoToEmployee(employeeDto *dto.CreateEmployeeRequestDto, peopleId uint64) *models.Employee {
	employee := new(models.Employee)
	employee.PeopleId = peopleId
	employee.PostId = employeeDto.PostId
	employee.EmploymentDate = employeeDto.EmploymentDate
	employee.Salary = employeeDto.Salary
	return employee
}

func FromUpdateRequestDtoToEmployee(employeeDto *dto.UpdateEmployeeRequestDto) *models.Employee {
	employee := new(models.Employee)
	employee.EmploymentDate = employeeDto.EmploymentDate
	if employeeDto.FireDate != nil {
		employee.FireDate.Time = *employeeDto.FireDate
		employee.FireDate.Valid = true
	}
	employee.Salary = employeeDto.Salary
	return employee
}
