package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/bevrist/simple-notify/core/pkg/api"
)

// set database test file ENV for test
func TestMain(m *testing.M) {
	testDbLocation := "/tmp/database_test/"
	//delete if test file exists
	if _, err := os.Stat(testDbLocation); err == nil {
		os.RemoveAll(testDbLocation)
	}
	// Create test database directory
	err := os.MkdirAll(testDbLocation, 0755)
	if err != nil {
		panic(err)
	}
	os.Setenv("DATABASE_DIR", testDbLocation)
	defer os.Unsetenv("DATABASE_DIR")
	dbInit() //reload init now that env is set
	os.Exit(m.Run())
}

func TestNewMessage(t *testing.T) {
	var testMsg = api.Message{
		TimeStamp:    123,
		UserID:       "1",
		Message:      "test message 1",
		MessageGroup: "test group",
		Severity:     "test severity",
	}
	err := NewMessage(testMsg, "TestNewMessage")
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
			TimeStamp:    123,
			UserID:       "1",
			Message:      "test message" + i,
			MessageGroup: "test group" + i,
			Severity:     "test severity" + i,
		}
		NewMessage(msg, "BenchmarkNewMessage")
	}
}

func TestGetAllMessages(t *testing.T) {
	var testMsg = api.Message{
		TimeStamp:    123,
		UserID:       "testgetallmessages",
		Message:      "test message 1",
		MessageGroup: "test group",
		Severity:     "test severity",
	}
	var testMsg2 = api.Message{
		TimeStamp:    124,
		UserID:       "testgetallmessages",
		Message:      "test message 2",
		MessageGroup: "test group 2",
		Severity:     "test severity",
	}
	err := NewMessage(testMsg, "testgetallmessages")
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}
	err = NewMessage(testMsg2, "testgetallmessages")
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}

	messages := GetAllMessages("testgetallmessages")
	if len(messages) != 2 {
		t.Errorf("GetAllMessages failed: expected 2 messages, got %d", len(messages))
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

// TODO: test meta table

// TODO: fuzz test database entries
