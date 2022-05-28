// dispatcher is responsible for sending messages
package dispatcher

import (
	"fmt"
	"time"
)

func StartDispatcher() {
	go dispatchLoop()
}

// check database and dispatch messages on regular interval
func dispatchLoop() {
	for range time.Tick(time.Second * 10) {
		fmt.Println("Foo")
		//TODO
		// send messages that are not yet dispatched from database
		// get messageGroup (stream?) that message is part of for sender addresses and rules
		// send messages
		// mark messages as dispatched
	}
}
