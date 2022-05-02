package database

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	_ "github.com/bevrist/simple-notify/core/pkg/api"
)

var db *sql.DB

//init opens database and initializes if necessary
func init() {
	log.Println("database.init()")
	//Create a folder/directory at a full qualified path
	err := os.MkdirAll("./data/", 0755)
	if err != nil {
		log.Fatal(err)
	}

	db, _ = sql.Open("sqlite3", "./data/database.db")

	// create main table
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS db (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement.Close()

	// create metadata table
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS meta (id INTEGER PRIMARY KEY, version TEXT)")
	statement.Exec()
	statement.Close()

	// cache sql statements
	stInsert, _ = db.Prepare("INSERT INTO db (firstname, lastname) VALUES (?, ?)")
}

var stInsert *sql.Stmt

func AddPeep(fn, ln string) {
	stInsert.Exec(fn, ln)
}

// func GetMessages() []api.Message {
func GetMessages() []string {
	// var row struct {
	//   age  int
	//   name string
	// }
	// err = db.QueryRow("SELECT|people|age,name|age=?", 3).Scan(&row.age, &row.name)
	rows, _ := db.Query("SELECT id, firstname, lastname FROM db")
	var id int
	var firstname string
	var lastname string
	var list []string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		list = append(list, strconv.Itoa(id)+": "+firstname+" "+lastname)
	}
	return list
}
