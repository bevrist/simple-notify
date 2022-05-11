package database

import (
	"database/sql"
	"encoding/base64"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

var db *sql.DB

func init() {
	log.Println("database.init()")
	// initialize when loading package.
	// dbInit() is separate so it can be called again with different env vars when unit testing
	dbInit()
}

// dbIinit opens database and initializes if necessary
func dbInit() {
	// load env var for database file location
	dbLocation := os.Getenv("DATABASE_FILE")
	if dbLocation == "" {
		dbLocation = "./data/database.db"
		// Create default database directory if it doesn't exist
		err := os.MkdirAll("./data/", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	var err error
	db, err = sql.Open("sqlite3", "file:"+dbLocation+"?cache=shared&mode=rwc&_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	// create main table
	// timestamp, user_id, message, message_group, message_severity
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS db (timestamp INTEGER PRIMARY KEY AUTOINCREMENT, user_id TEXT, message TEXT, message_group TEXT, message_severity TEXT)")
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
	// // check or populate version
	// stInsert, err = db.Prepare("INSERT INTO db (timestamp, user_id, message, message_group, message_severity) VALUES (?, ?, ?, ?, ?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//TODO: add metadata into the metadata table

	// cache sql statements
	stInsert, err = db.Prepare("INSERT INTO db (timestamp, user_id, message, message_group, message_severity) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
}

var stInsert *sql.Stmt

// NewMessage stores a new api.message in the database
func NewMessage(msg api.Message) error {
	// TODO: move timestamp out of database and into receivers
	// _, err := stInsert.Exec(msg.TimeStamp, msg.UserID, msg.Message, msg.MessageGroup, msg.Severity)
	msg = prepareMessage(msg)
	_, err := stInsert.Exec(time.Now().UnixNano(), msg.UserID, msg.Message, msg.MessageGroup, msg.Severity)
	if err != nil {
		return err
	}
	return nil
}

func prepareMessage(in api.Message) api.Message {
	in.Message = base64.StdEncoding.EncodeToString([]byte(in.Message))
	return in
}

func readMessage(in api.Message) api.Message {
	msgDec, err := base64.StdEncoding.DecodeString(in.Message)
	if err != nil {
		log.Printf("Some error occurred during base64 decode. Error %s", err.Error())
	}
	in.Message = string(msgDec)
	return in
}

// GetAllMessages returns a slice of all api.messages for a specific userId
func GetAllMessages(userId string) []api.Message {
	rows, err := db.Query("SELECT timestamp, user_id, message, message_group, message_severity FROM db WHERE user_id=?", userId)
	if err != nil {
		log.Println("ERROR: database.GetAllMessages(): ", userId, err)
	}
	var msgList []api.Message
	for rows.Next() {
		var msg api.Message
		rows.Scan(&msg.TimeStamp, &msg.UserID, &msg.Message, &msg.MessageGroup, &msg.Severity)
		msgList = append(msgList, readMessage(msg))
	}
	return msgList
}

// GetNewMessages returns a slice of all api.messages for a specific userId newer than specified timestamp
func GetNewMessages(userId string, timestamp int) []api.Message {
	rows, err := db.Query("SELECT timestamp, user_id, message, message_group, message_severity FROM db WHERE user_id=? AND timestamp>=?", userId, timestamp)
	if err != nil {
		log.Println("ERROR: database.GetNewMessages(): ", userId, err)
	}
	var msgList []api.Message
	for rows.Next() {
		var msg api.Message
		rows.Scan(&msg.TimeStamp, &msg.UserID, &msg.Message, &msg.MessageGroup, &msg.Severity)
		msgList = append(msgList, msg)
	}
	return msgList
}
