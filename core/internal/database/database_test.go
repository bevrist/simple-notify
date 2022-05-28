package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/bevrist/simple-notify/core/pkg/common"
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
	DbInit()               //reload init now that env is set
	os.RemoveAll("./data") //clear unused auto-generated database folder
	os.Exit(m.Run())
}

func TestNewMessage(t *testing.T) {
	var testMsg = common.Message{
		TimeStamp: 123,
		UserID:    "1",
		Message:   "test-message-1",
		StreamID:  "test-group",
		Severity:  "test-severity",
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
		msg := common.Message{
			TimeStamp: 123,
			UserID:    "1",
			Message:   "test message" + i,
			StreamID:  "test group" + i,
			Severity:  "test severity" + i,
		}
		NewMessage(msg, "BenchmarkNewMessage")
	}
}

func TestGetAllMessages(t *testing.T) {
	var testMsg = common.Message{
		TimeStamp: 123,
		UserID:    "testgetallmessages",
		Message:   "test-message-1",
		StreamID:  "test-group",
		Severity:  "test-severity",
	}
	var testMsg2 = common.Message{
		TimeStamp: 124,
		UserID:    "testgetallmessages",
		Message:   "test-message-2",
		StreamID:  "test-group 2",
		Severity:  "test-severity",
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
		fmt.Printf("messages:\n")
		fmt.Printf("+%v\n", messages)
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
	if messages[0].StreamID != testMsg.StreamID {
		t.Errorf("GetAllMessages failed: expected messageGroup %s, got %s", testMsg.StreamID, messages[0].StreamID)
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

func TestGetNewMessages(t *testing.T) {
	var testOldMsg = common.Message{
		TimeStamp: 499,
		UserID:    "TestGetNewMessages",
		Message:   "test-message-0",
		StreamID:  "test-group",
		Severity:  "test-severity",
	}
	var testMsg = common.Message{
		TimeStamp: 500,
		UserID:    "TestGetNewMessages",
		Message:   "test-message-1",
		StreamID:  "test-group",
		Severity:  "test-severity",
	}
	var testMsg2 = common.Message{
		TimeStamp: 501,
		UserID:    "TestGetNewMessages",
		Message:   "test-message-2",
		StreamID:  "test-group 2",
		Severity:  "test-severity",
	}
	err := NewMessage(testOldMsg, "TestGetNewMessages")
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}
	err = NewMessage(testMsg, "TestGetNewMessages")
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}
	err = NewMessage(testMsg2, "TestGetNewMessages")
	if err != nil {
		t.Errorf("NewMessage failed: %v", err)
		t.FailNow()
	}

	messages := GetNewMessages("TestGetNewMessages", 500)
	if len(messages) != 2 {
		t.Errorf("GetAllMessages failed: expected 2 messages, got %d", len(messages))
		fmt.Printf("messages:\n")
		fmt.Printf("+%v\n", messages)
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
	if messages[0].StreamID != testMsg.StreamID {
		t.Errorf("GetAllMessages failed: expected messageGroup %s, got %s", testMsg.StreamID, messages[0].StreamID)
		t.FailNow()
	}
	if messages[0].Severity != testMsg.Severity {
		t.Errorf("GetAllMessages failed: expected severity %s, got %s", testMsg.Severity, messages[0].Severity)
		t.FailNow()
	}
}

// TODO: test meta

// TODO: fuzz test database entries
