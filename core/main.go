package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bevrist/simple-notify/core/internal/database"
	"github.com/bevrist/simple-notify/core/pkg/api"
)

var randd *rand.Rand

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	randd = rand.New(s1)

	err := database.NewMessage(randMsg(), "main")
	if err != nil {
		log.Fatal(err)
	}

	setMsgTime := time.Now()
	for i := 0; i < 200; i++ {
		database.NewMessage(randMsg(), "main")
		if i%100000 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println(time.Since(setMsgTime))

	getMsgTime := time.Now()
	// var count int
	// for range database.GetAllMessages("1") {
	// 	// for range database.GetNewMessages("3", int(time.Now().UnixNano()-5000000000000)) {
	// 	count++
	// }
	fmt.Println("getting")
	for i := 0; i < 1; i++ {
		// database.GetAMessage(1)
		// msg := database.GetAMessage(i / 25)
		// for _, m := range msg {
		// 	fmt.Printf("+%v", m)
		// }
		// fmt.Println("")
	}
	fmt.Println("done ", time.Since(getMsgTime))
}

func randMsg() api.Message {
	msg := api.Message{
		UserID:       fmt.Sprint(randd.Intn(500)),
		Message:      fmt.Sprint(randd.Intn(99999)),
		MessageGroup: fmt.Sprint(randd.Intn(5)),
		Severity:     fmt.Sprint(randd.Intn(5)),
	}
	return msg
}
