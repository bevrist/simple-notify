package database

import (
	db "github.com/bevrist/simple-notify/core/internal/database/memory"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

// DbInit initializes the database with current env vars
func DbInit() {
	db.DbInit()
}

// NewMessage stores a new api.message in the database noted with the specified receiver source
func NewMessage(msg api.Message, source string) error {
	return db.NewMessage(msg, source)
}

// GetAllMessages returns a slice of all api.messages for a specific userId
func GetAllMessages(userId string) []api.Message {
	return db.GetAllMessages(userId)
}

// GetNewMessages returns a slice of all api.messages for a specific userId newer than specified timestamp
func GetNewMessages(userId string, timestamp int) []api.Message {
	return db.GetNewMessages(userId, timestamp)
}
