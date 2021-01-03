package postmark

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Client ...
type Client struct {
	APIToken string
}

// Send ...
func (c *Client) Send(ctx context.Context, email Email) (Response, error) {
	url := "https://api.postmarkapp.com/email"

	data, err := json.Marshal(email)
	if err != nil {
		return Response{}, err
	}

	// Make Postmark request
	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Postmark-Server-Token", c.APIToken)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	// Decode response
	dec := json.NewDecoder(res.Body)

	var response Response

	err = dec.Decode(&response)
	if err != nil {
		return Response{}, err
	}

	// Close body
	err = res.Body.Close()
	if err != nil {
		return Response{}, err
	}

	return response, nil
}
