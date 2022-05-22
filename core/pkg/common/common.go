// common contains the common types and interfaces used in the application
package common

// Message represents a standard log message with metadata
type Message struct {
	TimeStamp    int
	UserID       string
	MessageGroup string
	Message      string
	Severity     string
}
