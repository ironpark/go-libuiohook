go-libuiohook
====================================================================
go-libuiohook is a simple binding for [libuiohook](https://github.com/kwhat/libuiohook) library.

## Use

```golang
package main

import (
    "context"
    "fmt"
    uio "github.com/ironpark/go-libuiohook"
    "time"
)

func main() {
    // Move the mouse pointer to (100, 100)
    uio.PostEvent(&uio.Event{Type: uio.EvMouseMoved, Data: uio.EventData{X: 100, Y: 100}})
	
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
	// Hook the keyboard & mouse
    uio.SetDispatchProc(func(event *uio.Event) {
        fmt.Println("Event:", event)
    })
    uio.Start(ctx)
}
```

