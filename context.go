package postmark

import (
	"context"
)

// WithContext ...
func WithContext(ctx context.Context, sender Sender) context.Context {
	ctx = context.WithValue(ctx, postmarkKey{}, sender)
	return ctx
}

// GetSender ...
func GetSender(ctx context.Context) Sender {
	if sender, ok := ctx.Value(postmarkKey{}).(Sender); ok {
		return sender
	}

	return NopSender{}
}

// Send ...
func Send(ctx context.Context, email Email) (Response, error) {
	return GetSender(ctx).Send(ctx, email)
}

// SendBatch ...
func SendBatch(ctx context.Context, emails ...Email) ([]Response, error) {
	return GetSender(ctx).SendBatch(ctx, emails...)
}

type postmarkKey struct{}
