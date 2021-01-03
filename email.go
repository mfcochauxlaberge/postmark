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
	TextBody string `json:",omitempty"`
	HTMLBody string `json:"HtmlBody"`

	// Template
	TemplateID    uint                   `json:",omitempty"`
	TemplateAlias string                 `json:",omitempty"`
	TemplateModel map[string]interface{} `json:",omitempty"`
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

	if len(e.Headers) > 0 {
		email["Headers"] = e.Headers
	}

	if !e.UsesTemplate() {
		// Email content
		email["Subject"] = e.Subject
		email["TextBody"] = e.TextBody
		email["HtmlBody"] = e.HTMLBody
	} else {
		// Template
		if e.TemplateID != 0 {
			email["TemplateId"] = e.TemplateID
		} else if e.TemplateAlias != "" {
			email["TemplateAlias"] = e.TemplateAlias
		}

		if len(e.TemplateModel) > 0 {
			email["TemplateModel"] = e.TemplateModel
		}

		email["InlineCss"] = e.InlineCSS
	}

	data, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	return data, nil
}
