package postmark

import "time"

// Email ...
type Email struct {
	From          string
	To            string
	Subject       string `json:",omitempty"`
	TextBody      string
	HTMLBody      string `json:"HtmlBody"`
	MessageStream string
	Tag           string            `json:",omitempty"`
	Headers       map[string]string `json:",omitempty"`
}

// Response ...
type Response struct {
	To          string
	SubmittedAt time.Time
	MessageID   string
	ErrorCode   int
	Message     string
}
