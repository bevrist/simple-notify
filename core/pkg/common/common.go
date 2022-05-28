// common contains the common types and interfaces used in the application
package common

// Message represents a standard log message with metadata
type Message struct {
	TimeStamp    int    `json:"timestamp"`    //(optional)
	MessageGroup string `json:"messageGroup"` //TODO probably rename this
	Message      string `json:"message"`
	Severity     string `json:"severity"`

	UserID         string
	DispatchedTime string `json:"dispatched"`
}
