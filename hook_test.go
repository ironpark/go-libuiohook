package go_libuiohook

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetScreenInfo(t *testing.T) {
	got, err := GetScreenInfo()
	if err != nil {
		t.Errorf("GetScreenInfo() error = %v", err)
		return
	}
	t.Logf("GetScreenInfo() = %v", got)
	for i, data := range got {
		t.Logf("Screen %d: %+v", i, data)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	SetDispatchProc(func(event *Event) {
		fmt.Println("Event:", event)
		if event.Type.IsKeyEvent() {
			fmt.Println("Key:", string(event.Key()))
		}
	})
	fmt.Println(Start(ctx))
}
