package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllCustomers(db *sql.DB) ([]models.Customer, error) {
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	customers := []models.Customer{}

	for rows.Next() {
		c := models.Customer{}
		err := rows.Scan(&c.Id, &c.Ogrn, &c.Name, &c.Logo)
		if err != nil {
			continue
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func CreateCustomer(db *sql.DB, customer *models.Customer) (*models.Customer, error) {
	row := db.QueryRow("INSERT INTO customers VALUES ($1, $2, $3) RETURNING id", customer.Ogrn, customer.Name, customer.Logo)
	err := row.Scan(&customer.Id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func DeleteCustomer(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCustomer(db *sql.DB, customer *models.Customer) (*models.Customer, error) {

	result, err := db.Exec("UPDATE posts SET ogrn = $1, name = $2, logo = $3 WHERE id = $4", customer.Ogrn, customer.Name, customer.Logo, customer.Id)
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

	return customer, nil
}
