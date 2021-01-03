package postmark

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

const (
	PostmarkAPI   = "https://api.postmarkapp.com"
	EndpointEmail = PostmarkAPI + "/email"
	EndpointBatch = PostmarkAPI + "/email/batch"
)

// Server ...
type Server struct {
	Token string
}

// Send ...
func (s *Server) Send(ctx context.Context, email Email) (Response, error) {
	data, err := json.Marshal(email)
	if err != nil {
		return Response{}, err
	}

	// Make Postmark request
	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, EndpointEmail, body)
	if err != nil {
		return Response{}, err
	}

	// Send and get response
	response, err := s.send(req)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

// SendBatch ...
func (s *Server) SendBatch(ctx context.Context, emails []Email) (Response, error) {
	data, err := json.Marshal(emails)
	if err != nil {
		return Response{}, err
	}

	// Make Postmark request
	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, EndpointBatch, body)
	if err != nil {
		return Response{}, err
	}

	// Send and get response
	response, err := s.send(req)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (s *Server) send(req *http.Request) (Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Postmark-Server-Token", s.Token)

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