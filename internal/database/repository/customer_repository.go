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
		err := rows.Scan(&c.Id, &c.Ogrn, &c.Name, &c.Logo, &c.LogoExtension)
		if err != nil {
			continue
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func CreateCustomer(db *sql.DB, customer *models.Customer) (*models.Customer, error) {
	row := db.QueryRow("INSERT INTO customers (ogrn, name, logo, logo_extension) VALUES ($1, $2, $3, $4) RETURNING id", customer.Ogrn, customer.Name, customer.Logo, customer.LogoExtension)
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

func PatchCustomer(db *sql.DB, id uint64, customer *models.Customer) (*models.Customer, error) {

	result, err := db.Exec("UPDATE posts SET ogrn = $1, name = $2 WHERE id = $3", customer.Ogrn, customer.Name, id)
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

func UpdateCustomer(db *sql.DB, id uint64, customer *models.Customer) (*models.Customer, error) {

	result, err := db.Exec("UPDATE posts SET ogrn = $1, name = $2, logo = $3, logo_extension = $4 WHERE id = $5", customer.Ogrn, customer.Name, customer.Logo, customer.LogoExtension, id)
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

func ExistsCustomer(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM customers WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
