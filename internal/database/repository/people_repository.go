package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetPeopleByLogin(db *sql.DB, login string) (*models.People, error) {
	people := new(models.People)
	row := db.QueryRow("SELECT * FROM people WHERE login = $1", login)
	err := row.Scan(&people.Id, &people.LastName, &people.FirstName, &people.MiddleName, &people.BirthDate, &people.Login, &people.PasswordHash)
	if err != nil {
		return nil, err
	}
	return people, nil
}

func ExistsPeopleByLogin(db *sql.DB, login string) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM people WHERE login = $1))", login)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func GetAllPeople(db *sql.DB) ([]models.People, error) {
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	people := []models.People{}

	for rows.Next() {
		p := models.People{}
		err := rows.Scan(&p.Id, &p.LastName, &p.FirstName, &p.MiddleName, &p.BirthDate)
		if err != nil {
			continue
		}
		people = append(people, p)
	}

	return people, nil
}

func CreatePeople(db *sql.DB, people *models.People) (*models.People, error) {
	row := db.QueryRow("INSERT INTO people VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", people.LastName, people.FirstName, people.MiddleName, people.BirthDate, people.Login, people.PasswordHash)
	err := row.Scan(&people.Id)
	if err != nil {
		return nil, err
	}
	return people, nil
}

func DeletePeople(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM people WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePeople(db *sql.DB, people *models.People) (*models.People, error) {

	result, err := db.Exec("UPDATE people SET lastname = $1, firstname = $2, middlename = $3, birthdate = $4 WHERE id = $5", people.LastName, people.FirstName, people.MiddleName, people.BirthDate, people.Login)
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

	return people, nil
}

func ExistsPeople(db *sql.DB, id uint64) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM people WHERE id = $1))", id)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
