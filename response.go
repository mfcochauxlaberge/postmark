package postmark

import "time"

// Response ...
type Response struct {
	To          string
	SubmittedAt time.Time
	MessageID   string
	ErrorCode   int
	Message     string
}
