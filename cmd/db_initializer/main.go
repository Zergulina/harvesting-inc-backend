package main

import "database/sql"

func main() {
	connStr := "user=postgres password=1234 dbname=agro sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
