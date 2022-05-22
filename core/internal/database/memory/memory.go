package memory

import (
	"github.com/bevrist/simple-notify/core/pkg/api"
)

var db []api.Message
var SchemaVersion float32 = 1.0

// DbInit does nothing for memory database
func DbInit() {
}

// NewMessage stores a new api.message in the database
func NewMessage(msg api.Message, source string) error {
	db = append(db, msg)
	return nil
}

// GetAllMessages returns a slice of all api.messages for a specific userId
func GetAllMessages(userId string) []api.Message {
	// only return db entries with api.userId == userId
	var msgList []api.Message
	for _, msg := range db {
		if msg.UserID == userId {
			msgList = append(msgList, msg)
		}
	}
	return msgList
}

// GetNewMessages returns a slice of all api.messages for a specific userId newer than specified timestamp
func GetNewMessages(userId string, timestamp int) []api.Message {
	// loop through db in reverse until we find the first message older than timestamp
	for i := len(db) - 1; i >= 0; i-- {
		if db[i].TimeStamp < timestamp {
			return db[i+1:]
		}
	}
	return []api.Message{}
}
