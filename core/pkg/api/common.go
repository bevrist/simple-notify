package api

type Message struct {
	Time           string
	UserID         int
	MessageGroup   string
	Message        string
	Severity       string
	AttachmentUrls []string
}
