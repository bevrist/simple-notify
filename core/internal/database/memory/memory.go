// memory is a in-memory based implementation of the database interface
package memory

import (
	"github.com/bevrist/simple-notify/core/pkg/common"
)

var db []common.Message
var SchemaVersion float32 = 1.0

// DbInit does nothing for memory database
func DbInit() {
}

// NewMessage stores a new common.message in the database
func NewMessage(msg common.Message, source string) error {
	db = append(db, msg)
	return nil
}

// GetAllMessages returns a slice of all common.messages for a specific userId
func GetAllMessages(userId string) []common.Message {
	// only return db entries with common.userId == userId
	var msgList []common.Message
	for _, msg := range db {
		if msg.UserID == userId {
			msgList = append(msgList, msg)
		}
	}
	return msgList
}

// GetNewMessages returns a slice of all common.messages for a specific userId newer than specified timestamp
func GetNewMessages(userId string, timestamp int) []common.Message {
	// loop through db in reverse until we find the first message older than timestamp
	for i := len(db) - 1; i >= 0; i-- {
		if db[i].TimeStamp < timestamp {
			return db[i+1:]
		}
	}
	return []common.Message{}
}
