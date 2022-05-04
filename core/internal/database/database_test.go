package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

// set database test file ENV for test
func TestMain(m *testing.M) {
	testDbFile := "/tmp/test.db"
	//delete if file exists
	if _, err := os.Stat(testDbFile); err == nil {
		os.Remove(testDbFile)
	}
	os.Setenv("DATABASE_FILE", testDbFile)
	defer os.Unsetenv("DATABASE_FILE")
	dbInit() //reload init now that env is set
	os.Exit(m.Run())
}

var testMsg = api.Message{
	UserID:       "1",
	Message:      "test message 1",
	MessageGroup: "test group",
	Severity:     "test severity",
}

var testMsg2 = api.Message{
	UserID:       "1",
	Message:      "test message 2",
	MessageGroup: "test group",
	Severity:     "test severity",
}

var testMsg3 = api.Message{
	UserID:       "2",
	Message:      "test message",
	MessageGroup: "test group",
	Severity:     "test severity",
}

func TestNewMessage(t *testing.T) {
	err := NewMessage(testMsg)
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}
}

func BenchmarkNewMessage(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		i := fmt.Sprint(n)
		msg := api.Message{
			UserID:       "1",
			Message:      "test message" + i,
			MessageGroup: "test group" + i,
			Severity:     "test severity" + i,
		}
		NewMessage(msg)
	}
}

// TODO: BenchmarkGetAllMessage

func TestGetAllMessage(t *testing.T) {
	messages := GetAllMessages("1")
	if len(messages) != 1 {
		t.Errorf("GetAllMessages failed: expected 1 messages, got %d", len(messages))
		t.FailNow()
	}
	if messages[0].UserID != testMsg.UserID {
		t.Errorf("GetAllMessages failed: expected userId %s, got %s", testMsg.UserID, messages[0].UserID)
		t.FailNow()
	}
	if messages[0].Message != testMsg.Message {
		t.Errorf("GetAllMessages failed: expected message %s, got %s", testMsg.Message, messages[0].Message)
		t.FailNow()
	}
	if messages[0].MessageGroup != testMsg.MessageGroup {
		t.Errorf("GetAllMessages failed: expected messageGroup %s, got %s", testMsg.MessageGroup, messages[0].MessageGroup)
		t.FailNow()
	}
	if messages[0].Severity != testMsg.Severity {
		t.Errorf("GetAllMessages failed: expected severity %s, got %s", testMsg.Severity, messages[0].Severity)
		t.FailNow()
	}
}

func TestGetAllMessageInvalidUser(t *testing.T) {
	messages := GetAllMessages("999")
	if len(messages) != 0 {
		t.Errorf("GetAllMessages failed: expected 0 messages, got %d", len(messages))
		t.FailNow()
	}
}

// TODO: test GetNewMessages()
//get current timestamp
//create new message for user
//get all messages for user > 1
//get new messages after timestamp = 1
