package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("start")
	//Create a folder/directory at a full qualified path
	err := os.MkdirAll("./data/", 0755)
	if err != nil {
		log.Fatal(err)
	}

	database, _ := sql.Open("sqlite3", "./data/bogo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Rob", "Gronkowski")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people VALUES")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
