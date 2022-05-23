// common contains the common types and interfaces used in the application
package common

// Message represents a standard log message with metadata
type Message struct {
	TimeStamp    int `json:"timestamp"` //(optional)
	UserID       string
	MessageGroup string `json:"messageGroup"`
	Message      string `json:"message"`
	Severity     string `json:"severity"`
}
