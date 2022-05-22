package sqlite

import (
	"database/sql"
	"encoding/base64"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

var db *sql.DB
var SchemaVersion float32 = 1.0

// DbInit opens database and initializes if necessary
func DbInit() {
	// load env var for database location
	dbLocation := os.Getenv("DATABASE_DIR")
	if dbLocation == "" {
		dbLocation = "./data"
		// Create default database directory if it doesn't exist
		err := os.MkdirAll("./data/", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	var err error
	db, err = sql.Open("sqlite3", "file:"+dbLocation+"/database.db?cache=shared&mode=rwc&_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	// create main table
	// timestamp, user_id, message, message_group, message_severity, source
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS db (timestamp INTEGER, user_id TEXT, message TEXT, message_group TEXT, message_severity TEXT, source TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	// create metadata table
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS meta (key TEXT PRIMARY KEY, value TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	// add metadata into the metadata table
	statement, err = db.Prepare("INSERT INTO meta (key, value) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("SchemaVersion", SchemaVersion)
	statement.Close()

	// =========================
	// cache sql statements
	stInsert, err = db.Prepare("INSERT INTO db (timestamp, user_id, message, message_group, message_severity, source) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
}

var stInsert *sql.Stmt

// NewMessage stores a new api.message in the database
func NewMessage(msg api.Message, source string) error {
	msg = prepareMessage(msg)
	_, err := stInsert.Exec(msg.TimeStamp, msg.UserID, msg.Message, msg.MessageGroup, msg.Severity, source)
	if err != nil {
		return err
	}
	return nil
}

// prepareMessage encodes the message portion of the api.Message struct before storing it in the database
func prepareMessage(in api.Message) api.Message {
	in.Message = base64.StdEncoding.EncodeToString([]byte(in.Message))
	return in
}

// readMessage decodes the message portion of the api.Message struct before returning it
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
		msgList = append(msgList, readMessage(msg))
	}
	return msgList
}
