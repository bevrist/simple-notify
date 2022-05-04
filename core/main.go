package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/internal/database"
	"github.com/bevrist/simple-notify/core/pkg/api"
)

func main() {
	msg := api.Message{
		UserID:       "1",
		Message:      "test message",
		MessageGroup: "test group",
		Severity:     "test severity",
	}
	err := database.NewMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
	msg.Severity = "2222"
	err = database.NewMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
	msg.Severity = "33333"
	err = database.NewMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range database.GetAllMessages("1") {
		fmt.Printf("%+v \n", message)
	}
}
