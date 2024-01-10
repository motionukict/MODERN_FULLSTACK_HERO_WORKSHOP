package main

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Print(err)
	}
	return db
}

func GetNames(db *sql.DB) []string {
	var names []string
	//query the database and select from column `name` in table `names`
	rows, err := db.Query("Select NAMES from names")
	if err != nil {
		log.Print(err)
	}
	//loop as long as there are rows
	for rows.Next() {
		var name string
		//scan the row and assign the value of the column `name` to the variable `name`
		err = rows.Scan(&name)
		if err != nil {
			log.Print(err)
		}
		//append the value of `name` to the slice `names`
		names = append(names, name)
	}
	return names
}

func InsertName(db *sql.DB, name string) {
	//prepare the statement by inserting the value of `name` into the query
	stmt, err := db.Prepare("INSERT INTO names(NAMES) values(?)")

	if err != nil {
		log.Print(err)
	}
	// execute the statement
	_, err = stmt.Exec(name)
	if err != nil {
		log.Print(err)
	}

}