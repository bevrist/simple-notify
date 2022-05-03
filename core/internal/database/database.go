package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/pkg/api"
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

	db, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		log.Fatal(err)
	}

	// create main table
	// timestamp, user_id, message, message_group, message_severity
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS db (timestamp INTEGER PRIMARY KEY, user_id TEXT, message TEXT, message_group TEXT, message_severity TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	// create metadata table
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS meta (key INTEGER PRIMARY KEY, version TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	// cache sql statements
	stInsert, err = db.Prepare("INSERT INTO db (timestamp, user_id, message, message_group, message_severity) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
}

var stInsert *sql.Stmt

// NewMessage stores a new message object in the database
func NewMessage(msg api.Message) {
	// TODO: test that adding multiple messages at same time actually functions as expected
	_, err := stInsert.Exec(time.Now().Unix(), msg.UserID, msg.Message, msg.MessageGroup, msg.Severity)
	// _, err := stInsert.Exec(time.Now().UnixNano(), msg.UserID, msg.Message, msg.MessageGroup, msg.Severity)
	if err != nil {
		log.Println("ERROR: database.NewMessage(): ", err)
	}
}

// GetMessages returns all messages for a specific user
func GetMessages(userId string) []api.Message {
	rows, _ := db.Query("SELECT timestamp, user_id, message, message_group, message_severity FROM db WHERE user_id=?", userId)
	var msgList []api.Message
	for rows.Next() {
		var msg api.Message
		rows.Scan(&msg.TimeStamp, &msg.UserID, &msg.Message, &msg.MessageGroup, &msg.Severity)
		msgList = append(msgList, msg)
	}
	return msgList
}
