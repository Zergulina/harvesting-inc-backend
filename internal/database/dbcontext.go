package database

import (
	"backend/internal/config"
	"backend/internal/database/repository"
	"backend/internal/helpers"
	"backend/internal/models"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connectionString, exists := os.LookupEnv("DB_CONNECTION_STRING")
	if !exists {
		panic(".env var DB_CONNECTION_STRING does not exist")
	}
	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	DB = dbConn

	log.Println(DB.Stats().InUse)

	err = DB.Ping()
	if err != nil {
		panic("Error: Unable to ping database")
	}

	err = initDb(dbConn)
	if err != nil {
		panic(err)
	}

	return nil
}

func initDb(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS people (
			id SERIAL NOT NULL PRIMARY KEY,
			lastname TEXT NOT NULL,
			firstname TEXT NOT NULL,
			middlename TEXT,
			birthdate DATE NOT NULL,
			login TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT UNIQUE
		);

		CREATE TABLE IF NOT EXISTS employees (
			people_id INTEGER NOT NULL REFERENCES people(id),
			post_id INTEGER NOT NULL REFERENCES posts(id),
			employment_date DATE NOT NULL,
			fire_date DATE,
			salary INTEGER NOT NULL,
			PRIMARY KEY(people_id, post_id)
		);
		
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL NOT NULL PRIMARY KEY,
			ogrn TEXT NOT NULL,
			name CHAR(13) NOT NULL,
			logo BYTEA
		);

		CREATE TABLE IF NOT EXISTS crop_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS crops (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			crop_type_id INTEGER NOT NULL REFERENCES crop_types(id),
			description TEXT
		);

		CREATE TABLE IF NOT EXISTS fields (
			id SERIAL NOT NULL PRIMARY KEY,
			coords TEXT NOT NULL,
			customer_id INTEGER NOT NULL REFERENCES customers(id),
			crop_iв INTEGER NOT NULL REFERENCES crops(id)
		);

		CREATE TABLE IF NOT EXISTS statuses (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			is_available BOOLEAN NOT NULL
		);

		CREATE TABLE IF NOT EXISTS machine_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS machine_models (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			machine_type_id INTEGER NOT NULL REFERENCES machine_types(id)
		);
		
		CREATE TABLE IF NOT EXISTS machines (
			inv_number INTEGER NOT NULL,
			machine_model_id INTEGER NOT NULL REFERENCES machine_models(id),
			status_id INTEGER NOT NULL REFERENCES statuses(id),
			buy_date DATE NOT NULL,
			draw_down_date DATE,
			PRIMARY KEY(inv_number, machine_model_id)
		);

		CREATE TABLE IF NOT EXISTS equipment_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS equipment_models (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id)
		);

		CREATE TABLE IF NOT EXISTS equipment (
			inv_number INTEGER NOT NULL,
			equipment_model_id INTEGER NOT NULL REFERENCES equipment_models(id),
			status_id INTEGER NOT NULL REFERENCES statuses(id),
			buy_date DATE NOT NULL,
			draw_down_date DATE,
			PRIMARY KEY(inv_number, equipment_model_id)
		);

		CREATE TABLE IF NOT EXISTS machine_equipment_types (
			machine_type_id INTEGER NOT NULL REFERENCES machine_types(id),
			equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id),
			PRIMARY KEY(machine_type_id, equipment_type_id)
		);

		CREATE TABLE IF NOT EXISTS works (
			id SERIAL NOT NULL PRIMARY KEY,
			start_date DATE NOT NULL,
			end_date DATE,
			field_id INTEGER NOT NULL REFERENCES fields(id)
		);

		CREATE TABLE IF NOT EXISTS work_trip (
			id SERIAL NOT NULL PRIMARY KEY,
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP,
			crop_amount INTEGER NOT NULL,
			work_id INTEGER NOT NULL REFERENCES works(id),
			machine_inv_number INTEGER NOT NULL,
			machine_model_id INTEGER NOT NULL,
			FOREIGN KEY (machine_inv_number, machine_model_id) REFERENCES machines(inv_number, machine_model_id),
			equipment_inv_number INTEGER NOT NULL,
			equipment_model_id INTEGER NOT NULL,
			FOREIGN KEY (equipment_inv_number, equipment_model_id) REFERENCES equipment(inv_number, equipment_model_id)
		);

		CREATE TABLE IF NOT EXISTS vacations (
			people_id INTEGER NOT NULL REFERENCES people(id),
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			PRIMARY KEY(people_id, start_date)
		);
		`)

	if err != nil {
		return err
	}

	var adminPostId uint64

	isExists, err := repository.ExistsPostByName(db, config.AdminRole)
	if err != nil {
		return err
	}
	if !isExists {
		post, err := repository.CreatePost(db, &models.Post{Name: config.AdminRole})
		if err != nil {
			return err
		}

		adminPostId = post.Id
	} else {
		adminPost, err := repository.GetPostByName(db, config.AdminRole)
		if err != nil {
			return err
		}

		adminPostId = adminPost.Id
	}

	isExists, err = repository.ExistsPostByName(db, config.HrRole)
	if err != nil {
		return err
	}
	if !isExists {
		_, err = repository.CreatePost(db, &models.Post{Name: config.HrRole})
		if err != nil {
			return err
		}
	}

	isExists, err = repository.ExistsPostByName(db, config.DriverRole)
	if err != nil {
		return err
	}
	if !isExists {
		_, err = repository.CreatePost(db, &models.Post{Name: config.DriverRole})
		if err != nil {
			return err
		}
	}

	var adminId uint64

	isExists, err = repository.ExistsPeopleByLogin(db, config.AdminLogin)
	if err != nil {
		return err
	}
	if !isExists {
		admin := new(models.People)
		admin = &models.People{
			LastName:     config.AdminLastname,
			FirstName:    config.AdminFirstname,
			MiddleName:   config.AdminMiddlename,
			BirthDate:    config.AdminBirthdate,
			Login:        config.AdminLogin,
			PasswordHash: helpers.EncodeSha256(config.AdminPassword, config.DbSecretKey),
		}
		admin, err = repository.CreatePeople(db, admin)
		if err != nil {
			return err
		}
		adminId = admin.Id
	} else {
		admin, err := repository.GetPeopleByLogin(db, config.AdminLogin)
		if err != nil {
			return err
		}
		adminId = admin.Id
	}

	isExists, err = repository.ExistsEmployee(db, adminId, adminPostId)
	if err != nil {
		return err
	}
	if !isExists {
		employee := new(models.Employee)
		employee = &models.Employee{
			PeopleId:       adminId,
			PostId:         adminPostId,
			EmploymentDate: config.AdminEmploymentDate,
			Salary:         config.AdminSalary,
		}
		_, err = repository.CreateEmployee(db, employee)
		if err != nil {
			return err
		}
	}

	return nil
}
