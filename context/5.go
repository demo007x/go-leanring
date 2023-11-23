package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 附加值
	valueCtx := context.WithValue(ctx, key, "【监控】")
	go watch5(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func watch5(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine 监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
