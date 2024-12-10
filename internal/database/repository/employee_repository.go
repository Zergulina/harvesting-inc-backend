package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllEmployees(db *sql.DB) ([]models.Employee, error) {
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	employees := []models.Employee{}

	for rows.Next() {
		e := models.Employee{}
		err := rows.Scan(&e.PeopleId, &e.PostId, &e.EmploymentDate, &e.FireDate, &e.Salary)
		if err != nil {
			continue
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func GetAllEmployeesByPeopleId(db *sql.DB, peopleId uint64) ([]models.Employee, error) {
	rows, err := db.Query("SELECT * FROM employees WHERE people_id = $1", peopleId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	employees := []models.Employee{}

	for rows.Next() {
		e := models.Employee{}
		err := rows.Scan(&e.PeopleId, &e.PostId, &e.EmploymentDate, &e.FireDate, &e.Salary)
		if err != nil {
			continue
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func CreateEmployee(db *sql.DB, employee *models.Employee) (*models.Employee, error) {
	_, err := db.Exec("INSERT INTO employees VALUES ($1, $2, $3, $4, $5)", &employee.PeopleId, &employee.PostId, &employee.EmploymentDate, &employee.FireDate, &employee.Salary)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func DeleteEmployee(db *sql.DB, peopleId uint64, postId uint64) error {
	_, err := db.Exec("DELETE FROM employees WHERE people_id = $1 AND post_id = $2", peopleId, postId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(db *sql.DB, employee *models.Employee) (*models.Employee, error) {

	result, err := db.Exec("UPDATE employees SET employment_date = $1, fire_date = $2, salary = $3 WHERE people_id = $4 AND post_id = $5", employee.EmploymentDate, employee.FireDate, employee.Salary, employee.PeopleId, employee.PostId)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("")
	}

	return employee, nil
}
