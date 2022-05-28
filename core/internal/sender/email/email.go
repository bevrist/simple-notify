package email

import (
	"context"
	"os"

	"github.com/nikoksr/notify/service/sendgrid"
)

var sendgridApiKey string

func init() {
	sendgridApiKey = os.Getenv("SENDGRID_API_KEY")
}

// TODO come up with proper implementation for this
func SendMessage() {
	// Create a telegram service. Ignoring error for demo simplicity.
	sgService := sendgrid.New(sendgridApiKey, "notifier@evri.st", "Notifier Bot")

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	sgService.AddReceivers("brettevrist10@gmail.com")
	sgService.Send(
		context.Background(),
		"Subject/Title",
		"The actual message - Hello, you awesome gophers! :)",
	)

	// // Tell our notifier to use the telegram service. You can repeat the above process
	// // for as many services as you like and just tell the notifier to use them.
	// // Inspired by http middlewares used in higher level libraries.
	// notify.UseServices(sendgridService)

	// // Send a test message.
	// _ = notify.Send(
	// 	context.Background(),
	// 	"Subject/Title",
	// 	"The actual message - Hello, you awesome gophers! :)",
	// )
}
