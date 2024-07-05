package go_libuiohook

import (
	"context"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

/*
#cgo CFLAGS: -I./libuiohook/include -I./libuiohook/src -I./pbtools -I./proto
#cgo darwin CFLAGS:  -I./libuiohook/src/darwin -I./libuiohook/src -x objective-c -Wno-deprecated-declarations -DUSE_OBJC
#cgo darwin LDFLAGS: -framework Cocoa
#cgo windows CFLAGS: -I./libuiohook/src/windows
#cgo linux CFLAGS: -I./libuiohook/src/linux
#cgo linux CFLAGS: -I/usr/src
#cgo linux LDFLAGS: -L/usr/src -lX11 -lXtst
#cgo linux LDFLAGS: -lX11-xcb -lxcb -lxcb-xkb -lxkbcommon -lxkbcommon-x11
#include <stdio.h>
#include <stdlib.h>
#include "uiohook.h"
#include "pbtools.c"
#include "msg.h"
#include "msg.c"
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

void proto_encode(uiohook_event* const event, uint8_t *encoded, size_t *size){
	uint8_t workspace[128];
	struct uiohook_msg_t *evt_msg;
	evt_msg = uiohook_msg_new(&workspace[0], sizeof(workspace));
 	if (evt_msg == NULL) {
		return;
	}
	evt_msg->type = event->type;
	evt_msg->time = event->time;
	evt_msg->mask = event->mask;
	evt_msg->reserved = event->reserved;
	switch (event->type) {
    	case EVENT_KEY_TYPED:
    	case EVENT_KEY_PRESSED:
    	case EVENT_KEY_RELEASED:
			uiohook_msg_key_alloc(evt_msg);
			evt_msg->key_p->key_code = event->data.keyboard.keycode;
			evt_msg->key_p->raw_code = event->data.keyboard.rawcode;
			evt_msg->key_p->key_char = event->data.keyboard.keychar;
			break;
		case EVENT_MOUSE_CLICKED:
		case EVENT_MOUSE_PRESSED:
		case EVENT_MOUSE_RELEASED:
		case EVENT_MOUSE_MOVED:
		case EVENT_MOUSE_DRAGGED:
			uiohook_msg_mouse_alloc(evt_msg);
			evt_msg->mouse_p->button = event->data.mouse.button;
			evt_msg->mouse_p->clicks = event->data.mouse.clicks;
			evt_msg->mouse_p->x = event->data.mouse.x;
			evt_msg->mouse_p->y = event->data.mouse.y;
			break;
		case EVENT_MOUSE_WHEEL:
			uiohook_msg_wheel_alloc(evt_msg);
			evt_msg->wheel_p->clicks = event->data.wheel.clicks;
			evt_msg->wheel_p->x = event->data.wheel.x;
			evt_msg->wheel_p->y = event->data.wheel.y;
			evt_msg->wheel_p->type = event->data.wheel.type;
			evt_msg->wheel_p->amount = event->data.wheel.amount;
			evt_msg->wheel_p->rotation = event->data.wheel.rotation;
			evt_msg->wheel_p->direction = event->data.wheel.direction;
			break;
		default:
			break;
	}
	*size = uiohook_msg_encode(evt_msg, encoded, 64);
}

extern void goDispatchProc(uint8_t *encoded, size_t size);
static void __internel_dispatch(uiohook_event* const event) {
	size_t size;
	uint8_t encoded[64];
	proto_encode(event, &encoded[0], &size);
	if (size <= 0) {
		return;
	}
	goDispatchProc( &encoded[0], size);
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
				evw := &Event{
					Type:     EventType(ev.Type),
					Time:     time.Unix(0, int64(ev.Time)*int64(time.Microsecond)),
					Mask:     uint16(ev.Mask),
					Reserved: uint16(ev.Reserved),
					Data:     EventData{},
				}
				switch v := ev.GetData().(type) {
				case *Msg_Key:
					evw.Data.KeyCode = uint16(v.Key.KeyCode)
					evw.Data.RawCode = uint16(v.Key.RawCode)
					evw.Data.KeyChar = uint16(v.Key.KeyChar)
				case *Msg_Mouse:
					evw.Data.Button = uint16(v.Mouse.Button)
					evw.Data.Clicks = uint16(v.Mouse.Clicks)
					evw.Data.X = int16(v.Mouse.X)
					evw.Data.Y = int16(v.Mouse.Y)
				case *Msg_Wheel:
					evw.Data.Clicks = uint16(v.Wheel.Clicks)
					evw.Data.X = int16(v.Wheel.X)
					evw.Data.Y = int16(v.Wheel.Y)
					evw.Data.WheelType = uint8(v.Wheel.Type)
					evw.Data.Amount = uint16(v.Wheel.Amount)
					evw.Data.Rotation = int16(v.Wheel.Rotation)
					evw.Data.Direction = uint8(v.Wheel.Direction)
				default:
					fmt.Printf("Unknown event type %T\n", v)
				}
				dispatchProc(evw)
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
