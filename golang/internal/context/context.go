package context

import (
	"context"
)

var ctx context.Context

func init() {
	ctx = context.TODO()
}

func GetContext() *context.Context {
	return &ctx
}
