// database interfaces with the data store for the application
package database

import (
	// FIXME: add check to ensure this never goes into production `^[^\/\/]*"(?:.+\/sqlite)"`
	db "github.com/bevrist/simple-notify/core/internal/database/memory"
	"github.com/bevrist/simple-notify/core/pkg/common"
)

// DbInit initializes the database with current env vars
func DbInit() {
	db.DbInit()
}

// NewMessage stores a new common.message in the database noted with the specified receiver source
func NewMessage(msg common.Message, source string) error {
	return db.NewMessage(msg, source)
}

// GetAllMessages returns a slice of all common.messages for a specific userId
func GetAllMessages(userId string) []common.Message {
	return db.GetAllMessages(userId)
}

// GetNewMessages returns a slice of all common.messages for a specific userId newer than specified timestamp
func GetNewMessages(userId string, timestamp int) []common.Message {
	return db.GetNewMessages(userId, timestamp)
}
