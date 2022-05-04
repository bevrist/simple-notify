package database

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

var db *sql.DB

// init opens database and initializes if necessary
func init() {
	log.Println("database.init()")

	dbLocation := os.Getenv("DATABASE_FILE")
	if dbLocation == "" {
		dbLocation = "./data/database.db"
		//Create default database directory if it doesn't exist
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
	//TODO: add meta into metadata table

	// create random source
	randSource := rand.NewSource(time.Now().UnixNano())
	rand1 = rand.New(randSource)

	// cache sql statements
	stInsert, err = db.Prepare("INSERT INTO db (timestamp, user_id, message, message_group, message_severity) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
}

var rand1 *rand.Rand
var stInsert *sql.Stmt

// NewMessage stores a new message in the database
func NewMessage(msg api.Message) error {
	// add tiny bit of randomness to timestamp to ensure uniqueness for key
	_, err := stInsert.Exec(time.Now().UnixNano()+rand1.Int63n(999), msg.UserID, msg.Message, msg.MessageGroup, msg.Severity)
	if err != nil {
		return err
	}
	return nil
}

// GetAllMessages returns all messages for a specific user
func GetAllMessages(userId string) []api.Message {
	rows, err := db.Query("SELECT timestamp, user_id, message, message_group, message_severity FROM db WHERE user_id=?", userId)
	if err != nil {
		log.Println("ERROR: database.GetAllMessages(): ", userId, err)
	}
	var msgList []api.Message
	for rows.Next() {
		var msg api.Message
		rows.Scan(&msg.TimeStamp, &msg.UserID, &msg.Message, &msg.MessageGroup, &msg.Severity)
		msgList = append(msgList, msg)
	}
	return msgList
}
