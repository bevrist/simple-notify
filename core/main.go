package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/internal/database"
)

func main() {
	database.AddPeep("John", "Doe")
	database.AddPeep("Jeff", "Bezos")

	for _, message := range database.GetMessages() {
		fmt.Println(message)
	}
}
