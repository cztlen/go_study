package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	KEY = "trace_id"
)

func NewRequestID() string {
	return strings.Replace("qq", "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}
func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), getContextValue(ctx, KEY), message)
}
func getContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}
func ProceeEnter(ctx context.Context) {
	PrintLog(ctx, "阴天")
}

//获取结构体的所有tag
func main() {
	ProceeEnter(NewContextWithTraceID())
}
