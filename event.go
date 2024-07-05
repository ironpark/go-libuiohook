package go_libuiohook

/*
#include "uiohook.h"

// keyboard event data
uint16_t get_keycode(uiohook_event* const event) {
	return event->data.keyboard.keycode;
}
uint16_t get_rawcode(uiohook_event* const event) {
	return event->data.keyboard.rawcode;
}
uint16_t get_keychar(uiohook_event* const event) {
	return event->data.keyboard.keychar;
}
// mouse event data
uint16_t get_mouse_button(uiohook_event* const event) {
	return event->data.mouse.button;
}
uint16_t get_mouse_clicks(uiohook_event* const event) {
	return event->data.mouse.clicks;
}
int16_t get_mouse_x(uiohook_event* const event) {
	return event->data.mouse.x;
}
int16_t get_mouse_y(uiohook_event* const event) {
	return event->data.mouse.y;
}
// mouse wheel event data
uint16_t get_wheel_clicks(uiohook_event* const event) {
	return event->data.wheel.clicks;
}
int16_t get_wheel_x(uiohook_event* const event) {
	return event->data.wheel.x;
}
int16_t get_wheel_y(uiohook_event* const event) {
	return event->data.wheel.y;
}
uint8_t get_wheel_type(uiohook_event* const event) {
	return event->data.wheel.type;
}
uint16_t get_wheel_amount(uiohook_event* const event) {
	return event->data.wheel.amount;
}
int16_t get_wheel_rotation(uiohook_event* const event) {
	return event->data.wheel.rotation;
}
uint8_t get_wheel_direction(uiohook_event* const event) {
	return event->data.wheel.direction;
}
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type EventType uint8

const (
	EvHookEnabled   = EventType(C.EVENT_HOOK_ENABLED)
	EvHookDisabled  = EventType(C.EVENT_HOOK_DISABLED)
	EvKeyTyped      = EventType(C.EVENT_KEY_TYPED)
	EvKeyPressed    = EventType(C.EVENT_KEY_PRESSED)
	EvKeyReleased   = EventType(C.EVENT_KEY_RELEASED)
	EvMouseClicked  = EventType(C.EVENT_MOUSE_CLICKED)
	EvMousePressed  = EventType(C.EVENT_MOUSE_PRESSED)
	EvMouseReleased = EventType(C.EVENT_MOUSE_RELEASED)
	EvMouseMoved    = EventType(C.EVENT_MOUSE_MOVED)
	EvMouseDragged  = EventType(C.EVENT_MOUSE_DRAGGED)
	EvMouseWheel    = EventType(C.EVENT_MOUSE_WHEEL)
)

func (et EventType) IsKeyEvent() bool {
	switch et {
	case EvKeyTyped, EvKeyPressed, EvKeyReleased:
		return true
	}
	return false
}

func (et EventType) IsMouseEvent() bool {
	switch et {
	case EvMouseClicked, EvMousePressed, EvMouseReleased, EvMouseMoved, EvMouseDragged, EvMouseWheel:
		return true
	}
	return false
}

func (et EventType) String() string {
	switch et {
	case EvHookEnabled:
		return "HookEnabled"
	case EvHookDisabled:
		return "HookDisabled"
	case EvKeyTyped:
		return "KeyTyped"
	case EvKeyPressed:
		return "KeyPressed"
	case EvKeyReleased:
		return "KeyReleased"
	case EvMouseClicked:
		return "MouseClicked"
	case EvMousePressed:
		return "MousePressed"
	case EvMouseReleased:
		return "MouseReleased"
	case EvMouseMoved:
		return "MouseMoved"
	case EvMouseDragged:
		return "MouseDragged"
	case EvMouseWheel:
		return "MouseWheel"
	}
	return "Unknown"
}

type EventData struct {
	// KeyEvt
	KeyCode uint16
	RawCode uint16
	KeyChar uint16
	// MouseEvt
	Button uint16
	Clicks uint16
	X      int16
	Y      int16
	// WheelEvt
	WheelType uint8
	Amount    uint16
	Rotation  int16
	Direction uint8
}

type Event struct {
	Type     EventType
	Time     time.Time
	Mask     uint16
	Reserved uint16
	Data     EventData
}

type MouseWheelEventData struct {
	Type      uint8
	Amount    uint16
	Rotation  int16
	Direction uint8
}

func (e *Event) String() string {
	if e.Type.IsKeyEvent() {
		return fmt.Sprintf("[%-12s] %s, KeyCode: %d, RawCode: %d, KeyChar: %d", e.Type, e.Time, e.Data.KeyCode, e.Data.RawCode, e.Data.KeyChar)
	}
	if e.Type == EvMouseWheel {
		return fmt.Sprintf("[%-12s] %s, Clicks: %d, X: %d, Y: %d, WheelType: %d, Amount: %d, Rotation: %d, Direction: %d", e.Type, e.Time, e.Data.Clicks, e.Data.X, e.Data.Y, e.Data.WheelType, e.Data.Amount, e.Data.Rotation, e.Data.Direction)
	}
	if e.Type.IsMouseEvent() {
		return fmt.Sprintf("[%-12s] %s, Button: %d, Clicks: %d, X: %d, Y: %d", e.Type, e.Time, e.Data.Button, e.Data.Clicks, e.Data.X, e.Data.Y)
	}
	return fmt.Sprintf("[%-12s] %s, Mask: %d, Reserved: %d", e.Type, e.Time, e.Mask, e.Reserved)
}

func (e *Event) Key() rune {
	return rune(e.Data.KeyChar)
}

func (e *Event) cvtC() (ev *C.uiohook_event) {
	ev = &C.uiohook_event{
		_type:    C.event_type(e.Type),
		mask:     C.uint16_t(e.Mask),
		reserved: C.uint16_t(e.Reserved),
	}
	ev.time = C.uint64_t(e.Time.UnixNano() / int64(time.Microsecond))
	switch e.Type {
	case EvKeyTyped, EvKeyPressed, EvKeyReleased:
		keyboard := C.keyboard_event_data{
			keycode: C.uint16_t(e.Data.KeyCode),
			rawcode: C.uint16_t(e.Data.RawCode),
			keychar: C.uint16_t(e.Data.KeyChar),
		}
		*(*C.keyboard_event_data)(unsafe.Pointer(&ev.data)) = keyboard
	case EvMouseClicked, EvMousePressed, EvMouseReleased, EvMouseMoved, EvMouseDragged:
		mouse := C.mouse_event_data{
			button: C.uint16_t(e.Data.Button),
			clicks: C.uint16_t(e.Data.Clicks),
			x:      C.int16_t(e.Data.X),
			y:      C.int16_t(e.Data.Y),
		}
		*(*C.mouse_event_data)(unsafe.Pointer(&ev.data)) = mouse
	case EvMouseWheel:
		wheel := C.mouse_wheel_event_data{
			clicks:    C.uint16_t(e.Data.Clicks),
			x:         C.int16_t(e.Data.X),
			y:         C.int16_t(e.Data.Y),
			_type:     C.uint8_t(e.Data.WheelType),
			amount:    C.uint16_t(e.Data.Amount),
			rotation:  C.int16_t(e.Data.Rotation),
			direction: C.uint8_t(e.Data.Direction),
		}
		*(*C.mouse_wheel_event_data)(unsafe.Pointer(&ev.data)) = wheel
	}
	return
}

func cvtEvent(event *C.uiohook_event) *Event {
	e := &Event{
		Type:     EventType(event._type),
		Time:     time.Unix(0, int64(event.time)*int64(time.Microsecond)),
		Mask:     uint16(event.mask),
		Reserved: uint16(event.reserved),
	}
	switch e.Type {
	case EvKeyTyped, EvKeyPressed, EvKeyReleased:
		e.Data.KeyCode = uint16(C.get_keycode(event))
		e.Data.RawCode = uint16(C.get_rawcode(event))
		e.Data.KeyChar = uint16(C.get_keychar(event))
	case EvMouseClicked, EvMousePressed, EvMouseReleased, EvMouseMoved, EvMouseDragged:
		e.Data.Button = uint16(C.get_mouse_button(event))
		e.Data.Clicks = uint16(C.get_mouse_clicks(event))
		e.Data.X = int16(C.get_mouse_x(event))
		e.Data.Y = int16(C.get_mouse_y(event))
	case EvMouseWheel:
		e.Data.Clicks = uint16(C.get_wheel_clicks(event))
		e.Data.X = int16(C.get_wheel_x(event))
		e.Data.Y = int16(C.get_wheel_y(event))
		e.Data.WheelType = uint8(C.get_wheel_type(event))
		e.Data.Amount = uint16(C.get_wheel_amount(event))
		e.Data.Rotation = int16(C.get_wheel_rotation(event))
		e.Data.Direction = uint8(C.get_wheel_direction(event))
	}
	return e
}
