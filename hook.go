package go_libuiohook

import (
	"context"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

/*
#cgo darwin CFLAGS: -I./libuiohook/include -I./libuiohook/src/darwin -I./libuiohook/src -x objective-c -Wno-deprecated-declarations
#cgo darwin LDFLAGS: -framework Cocoa

#cgo windows CFLAGS: -I./libuiohook/include -I./libuiohook/src/windows
#cgo linux CFLAGS: -I./libuiohook/include -I./libuiohook/src/linux
#cgo linux CFLAGS: -I/usr/src
#cgo linux LDFLAGS: -L/usr/src -lX11 -lXtst
#cgo linux LDFLAGS: -lX11-xcb -lxcb -lxcb-xkb -lxkbcommon -lxkbcommon-x11
#include "uiohook.h"

#ifndef UIOHOOK_SHARED
#ifndef __UIOHOOK_INTERNAL_H__
#define __UIOHOOK_INTERNAL_H__
#include "logger.c"
#include "input_helper.c"
#include "input_hook.c"
#include "post_event.c"
#include "system_properties.c"
#endif
#endif

extern void goDispatchProc(uiohook_event* const event);
static void __internel_dispatch(uiohook_event* const event) {
	//clone the event
	uiohook_event* clone = (uiohook_event*)malloc(sizeof(uiohook_event));
	if (clone == NULL) {
		return;
	}
	memcpy(clone, event, sizeof(uiohook_event));
	goDispatchProc(clone);
}

void registerDispatchProc() {
	hook_set_dispatch_proc(&__internel_dispatch);
}
*/
import "C"

type ScreenData struct {
	Number int
	Width  int
	Height int
	X      int
	Y      int
}

var (
	lock         = sync.RWMutex{}
	dispatchProc = func(event *Event) {}
)

func init() {
	C.registerDispatchProc()
}

// PostEvent Send a virtual event back to the system.
func PostEvent(event *Event) {
	C.hook_post_event(event.cvtC())
}

func SetDispatchProc(f func(event *Event)) {
	lock.Lock()
	defer lock.Unlock()
	dispatchProc = f
}

// Start Insert the event hook.
func Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case ev := <-evCh:
				lock.RLock()
				dispatchProc(cvtEvent(ev))
				C.free(unsafe.Pointer(ev))
				lock.RUnlock()
			case <-ctx.Done():
				for C.hook_stop() != C.UIOHOOK_SUCCESS {
					fmt.Println("Failed to stop the hook. Retrying...")
					time.Sleep(1 * time.Second)
				}
				return
			}
		}
	}(ctx)
	res := int(C.hook_run())
	cancel()
	// Drain the event channel
	for {
		select {
		case <-evCh:
		default:
			if res != C.UIOHOOK_SUCCESS {
				return getError(res)
			}
			return nil
		}
	}
}

// Stop Withdraw the event hook.
func Stop() int {
	return int(C.hook_stop())
}

func GetScreenInfo() ([]ScreenData, error) {
	var count C.uchar
	cScreenData := C.hook_create_screen_info(&count)
	defer C.free(unsafe.Pointer(cScreenData))

	if cScreenData == nil {
		return nil, fmt.Errorf("failed to retrieve screen data")
	}

	// Convert the result to a Go slice
	screenDataSlice := make([]ScreenData, count)
	cArray := (*[1 << 28]C.screen_data)(unsafe.Pointer(cScreenData))[:count:count]
	for i := 0; i < int(count); i++ {
		screenDataSlice[i] = ScreenData{
			Number: int(cArray[i].number),
			Width:  int(cArray[i].width),
			Height: int(cArray[i].height),
			X:      int(cArray[i].x),
			Y:      int(cArray[i].y),
		}
	}

	return screenDataSlice, nil
}

// GetAutoRepeatRate Retrieves the keyboard auto-repeat rate.
func GetAutoRepeatRate() int {
	return int(C.hook_get_auto_repeat_rate())
}

// GetAutoRepeatDelay Retrieves the keyboard auto-repeat delay.
func GetAutoRepeatDelay() int {
	return int(C.hook_get_auto_repeat_delay())
}

// GetPointerAccelerationMultiplier Retrieves the mouse acceleration multiplier.
func GetPointerAccelerationMultiplier() int {
	return int(C.hook_get_pointer_acceleration_multiplier())
}

// GetPointerAccelerationThreshold Retrieves the mouse acceleration threshold.
func GetPointerAccelerationThreshold() int {
	return int(C.hook_get_pointer_acceleration_threshold())
}

// GetPointerSensitivity Retrieves the mouse sensitivity.
func GetPointerSensitivity() int {
	return int(C.hook_get_pointer_sensitivity())
}

// GetMultiClickTime Retrieves the double/triple click interval.
func GetMultiClickTime() int {
	return int(C.hook_get_multi_click_time())
}
