package utils

import (
	"context"
	"time"
)

const (
	contextGRPCServerTimeout = 30 * time.Second
	contextGRPCDialTimeout   = 5 * time.Second
)

// ContextWithTimeout ...
func ContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

// ContextGRPCServerWithTimeout ...
func ContextGRPCServerWithTimeout() (context.Context, context.CancelFunc) {
	return ContextWithTimeout(contextGRPCServerTimeout)
}

// ContextGRPCDialWithTimeout ...
func ContextGRPCDialWithTimeout() (context.Context, context.CancelFunc) {
	return ContextWithTimeout(contextGRPCDialTimeout)
}
