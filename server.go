package postmark

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	PostmarkAPI               = "https://api.postmarkapp.com"
	EndpointEmail             = PostmarkAPI + "/email"
	EndpointEmailWithTemplate = PostmarkAPI + "/email/withTemplate"
	EndpointBatch             = PostmarkAPI + "/email/batch"
	EndpointBatchWithTemplate = PostmarkAPI + "/email/batchWithTemplates"
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

	endpoint := EndpointEmail
	if email.UsesTemplate() {
		endpoint = EndpointEmailWithTemplate
	}

	// Make Postmark request
	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return Response{}, err
	}

	// Send and get response
	resData, err := s.send(req)
	if err != nil {
		return Response{}, err
	}

	response := Response{}

	err = json.Unmarshal(resData, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

// SendBatch ...
func (s *Server) SendBatch(ctx context.Context, emails ...Email) ([]Response, error) {
	if len(emails) == 0 {
		return nil, nil
	}

	batches := [][]Email{}

	if len(emails) <= 500 {
		batches = [][]Email{emails}
	} else {
		n := (len(emails) / 500) + 1

		batches = make([][]Email, 0, n)
	}

	messages := map[string][]Email{
		"Messages": emails,
	}

	data, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}

	endpoint := EndpointBatch
	if emails[0].UsesTemplate() {
		endpoint = EndpointBatchWithTemplate
	}

	// Make Postmark request
	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}

	// Send and get response
	resData, err := s.send(req)
	if err != nil {
		return nil, err
	}

	responses := []Response{}

	err = json.Unmarshal(resData, &responses)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (s *Server) send(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Postmark-Server-Token", s.Token)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Close body
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return resData, nil
}
