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
	a := map[string]bool{}

	SetDispatchProc(func(event *Event) {
		fmt.Println("Event:", event)
		if event.Type.IsKeyEvent() {
			fmt.Println("Key:", string(event.Key()))
			a[string(event.Key())] = true
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
