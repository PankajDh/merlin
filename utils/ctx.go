package utils

import (
	"context"
)

func GetContextValue(ctx context.Context, key string) any {
	if v := ctx.Value(key); v != nil {
		return v
	}

	return nil
}
