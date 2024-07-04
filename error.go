package go_libuiohook

/*
#include "uiohook.h"
*/
import "C"
import "errors"

var (
	// System level errors.

	ErrOutOfMemory = errors.New("failed to allocate memory")

	// X11 specific errors.

	ErrXOpenDisplay         = errors.New("failed to open X11 display")
	ErrXRecordNotFound      = errors.New("unable to locate XRecord extension")
	ErrXRecordAllocRange    = errors.New("unable to allocate XRecord range")
	ErrXRecordCreateContext = errors.New("unable to allocate XRecord context")
	ErrXRecordEnableContext = errors.New("failed to enable XRecord context")

	// Windows specific errors.

	ErrSetWindowsHookEx = errors.New("failed to register low level windows hook")

	// Darwin specific errors.

	ErrAxapiDisabled       = errors.New("failed to enable access for assistive devices")
	ErrCreateEventPort     = errors.New("failed to create apple event port")
	ErrCreateRunLoopSource = errors.New("failed to create apple run loop source")
	ErrGetRunloop          = errors.New("failed to acquire apple run loop")
	ErrCreateObserver      = errors.New("failed to create apple run loop observer")
	ErrUnknown             = errors.New("an unknown hook error occurred")
)

func getError(err int) error {
	switch err {
	case C.UIOHOOK_ERROR_OUT_OF_MEMORY:
		return ErrOutOfMemory
	case C.UIOHOOK_ERROR_X_OPEN_DISPLAY:
		return ErrXOpenDisplay
	case C.UIOHOOK_ERROR_X_RECORD_NOT_FOUND:
		return ErrXRecordNotFound
	case C.UIOHOOK_ERROR_X_RECORD_ALLOC_RANGE:
		return ErrXRecordAllocRange
	case C.UIOHOOK_ERROR_X_RECORD_CREATE_CONTEXT:
		return ErrXRecordCreateContext
	case C.UIOHOOK_ERROR_X_RECORD_ENABLE_CONTEXT:
		return ErrXRecordEnableContext
	case C.UIOHOOK_ERROR_SET_WINDOWS_HOOK_EX:
		return ErrSetWindowsHookEx
	case C.UIOHOOK_ERROR_AXAPI_DISABLED:
		return ErrAxapiDisabled
	case C.UIOHOOK_ERROR_CREATE_EVENT_PORT:
		return ErrCreateEventPort
	case C.UIOHOOK_ERROR_CREATE_RUN_LOOP_SOURCE:
		return ErrCreateRunLoopSource
	case C.UIOHOOK_ERROR_GET_RUNLOOP:
		return ErrGetRunloop
	case C.UIOHOOK_ERROR_CREATE_OBSERVER:
		return ErrCreateObserver
	default:
		return ErrUnknown
	}
}
