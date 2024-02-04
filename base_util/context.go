package base_util

import (
	"context"
	"time"
)

const (
	ContextUserKey = "_user"
	ContextTimeKey = "_now"
)

func SetUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, ContextUserKey, user)
}

func GetUser(ctx context.Context) string {
	if user, ok := ctx.Value(ContextUserKey).(string); ok {
		return user
	}
	return ""
}

func SetNow(ctx context.Context, now time.Time) context.Context {
	return context.WithValue(ctx, ContextTimeKey, now)
}

func GetNow(ctx context.Context) time.Time {
	if now, ok := ctx.Value(ContextTimeKey).(time.Time); ok {
		return now
	}
	return time.Now()
}
