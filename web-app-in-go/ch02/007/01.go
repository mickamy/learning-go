package main

import (
	"context"
	"fmt"
)

type TraceId string

const ZeroTraceID = ""

type traceIdKey struct{}

func SetTraceID(ctx context.Context, tid TraceId) context.Context {
	return context.WithValue(ctx, traceIdKey{}, tid)
}

func GetTraceID(ctx context.Context) TraceId {
	if tid, ok := ctx.Value(traceIdKey{}).(TraceId); ok {
		return tid
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
