package postmark

import (
	"context"
)

// Sender ...
type Sender interface {
	Send(context.Context, Email) (Response, error)
	SendBatch(context.Context, ...Email) ([]Response, error)
}

// NopSender ...
type NopSender struct{}

func (NopSender) Send(_ context.Context, _ Email) (Response, error) {
	return Response{}, nil
}
