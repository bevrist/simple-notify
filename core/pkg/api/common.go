package api

// Message represents a standard log message with metadata
type Message struct {
	TimeStamp    int
	UserID       string
	MessageGroup string
	Message      string
	Severity     string
}
