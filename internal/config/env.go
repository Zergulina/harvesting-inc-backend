package config

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var DbSecretKey string
var JwtSecretKey string

var AdminFirstname string
var AdminLastname string
var AdminMiddlename sql.NullString
var AdminLogin string
var AdminPassword string
var AdminBirthdate time.Time
var AdminEmploymentDate time.Time
var AdminSalary uint64

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dbSecretKey, exists := os.LookupEnv("DB_SECRET_KEY")
	if !exists {
		panic(".env var DB_SECRET_KEY does not exist")
	}
	DbSecretKey = dbSecretKey
	jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		panic(".env var JWT_SECRET_KEY does not exist")
	}
	JwtSecretKey = jwtSecretKey

	adminFirstname, exists := os.LookupEnv("ADMIN_FIRSTNAME")
	if !exists {
		panic(".env var ADMIN_FIRSTNAME does not exist")
	}
	AdminFirstname = adminFirstname
	adminLastname, exists := os.LookupEnv("ADMIN_LASTNAME")
	if !exists {
		panic(".env var ADMIN_LASTNAME does not exist")
	}
	AdminLastname = adminLastname
	var adminMiddlename sql.NullString
	adminMiddlenameStr, exists := os.LookupEnv("ADMIN_MIDDLENAME")
	if !exists {
		adminMiddlename = sql.NullString{Valid: false}
	} else {
		adminMiddlename = sql.NullString{String: adminMiddlenameStr, Valid: true}
	}
	AdminMiddlename = adminMiddlename
	adminLogin, exists := os.LookupEnv("ADMIN_LOGIN")
	if !exists {
		panic(".env var ADMIN_LOGIN does not exist")
	}
	AdminLogin = adminLogin
	adminPassword, exists := os.LookupEnv("ADMIN_LOGIN")
	if !exists {
		panic(".env var ADMIN_LOGIN does not exist")
	}
	AdminPassword = adminPassword
	adminBirthdateStr, exists := os.LookupEnv("ADMIN_BIRTHDATE")
	if !exists {
		panic(".env var ADMIN_BIRTHDATE does not exist")
	}
	adminBirthdate, err := time.Parse("2006-01-02", adminBirthdateStr)
	if err != nil {
		panic(err)
	}
	AdminBirthdate = adminBirthdate
	adminEmploymentDateStr, exists := os.LookupEnv("ADMIN_EMPLOYMENT_DATE")
	if !exists {
		panic(".env var ADMIN_EMPLOYMENT_DATE does not exist")
	}
	adminEmploymentDate, err := time.Parse("2006-01-02", adminEmploymentDateStr)
	if err != nil {
		panic(err)
	}
	AdminEmploymentDate = adminEmploymentDate
	adminSalaryStr, exists := os.LookupEnv("ADMIN_SALARY")
	if !exists {
		panic(".env var ADMIN_SALARY does not exist")
	}
	adminSalary, err := strconv.ParseUint(adminSalaryStr, 10, 64)
	if err != nil {
		panic(err)
	}
	AdminSalary = adminSalary
}
