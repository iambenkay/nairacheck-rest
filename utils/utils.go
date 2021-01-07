package utils

import (
	"context"
	"time"
)

func Contextualize(f func(context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	f(ctx)
}
