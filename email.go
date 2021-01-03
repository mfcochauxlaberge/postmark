package postmark

import (
	"encoding/json"
)

// Email ...
type Email struct {
	From          string
	To            string
	MessageStream string
	Tag           string            `json:",omitempty"`
	Headers       map[string]string `json:",omitempty"`

	// Content
	Subject  string `json:",omitempty"`
	TextBody string
	HTMLBody string `json:"HtmlBody"`

	// Template
	TemplateID    uint
	TemplateAlias string
	TemplateModel map[string]interface{}
	InlineCSS     bool
}

// UsesTemplate ...
func (e Email) UsesTemplate() bool {
	return e.Subject == "" && (e.TemplateID != 0 || e.TemplateAlias != "")
}

// MarshalJSON ...
func (e Email) MarshalJSON() ([]byte, error) {
	email := map[string]interface{}{}

	// Common fields
	email["From"] = e.From
	email["To"] = e.To
	email["MessageStream"] = e.MessageStream
	email["Tag"] = e.Tag
	email["Headers"] = e.Headers

	if !e.UsesTemplate() {
		// Email content
		email["Subject"] = e.Subject
		email["TextBody"] = e.TextBody
		email["HtmlBody"] = e.HTMLBody
	} else {
		// Template
		email["TemplateID"] = e.TemplateID
		email["TemplateAlias"] = e.TemplateAlias
		email["TemplateModel"] = e.TemplateModel
		email["InlineCss"] = e.InlineCSS
	}

	data, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	return data, nil
}
