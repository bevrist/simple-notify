// common contains the common types and interfaces used in the application
package common

// Message represents a standard log message with metadata
type Message struct {
	TimeStamp int    `json:"timestamp"` //(optional)
	StreamID  string `json:"streamID"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`

	UserID         string
	DispatchedTime string `json:"dispatched"`
}

// Stream represents a stream of messages from a source to a destination
type Stream struct {
	StreamID    string `json:"streamID"`
	Destination []Destination
	Schedule    string `json:"schedule"`
}

// Source represents a source of messages
type Source struct {
	Type   string `json:"type"`
	Config string `json:"config"`
}

// Destination represents a destination that can receive messages from a source
type Destination struct {
	Type             string `json:"type"`
	Config           string `json:"config"`
	ScheduleOverride string `json:"schedule"`
}
