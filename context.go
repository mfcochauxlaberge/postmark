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

type postmarkKey struct{}
