package utils

import (
	"context"
	"reflect"
	"time"
)

func Contextualize(f func(context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	f(ctx)
}

func IsStructEmpty(i interface{}) bool {
	zero := reflect.New(reflect.TypeOf(i))
	return zero.Elem().Interface() == i
}
