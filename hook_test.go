package go_libuiohook

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetScreenInfo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	SetDispatchProc(func(event *Event) {
		fmt.Println("Event:", event)
		if event.Type.IsKeyEvent() {
			fmt.Println("Key:", string(event.Key()))
		}
	})
	go func() {
		err := Start(ctx)
		if err != nil {
			fmt.Println("Start() error:", err)
		}
	}()
	time.Sleep(10 * time.Second)
}
