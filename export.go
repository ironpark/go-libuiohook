package go_libuiohook

/*
#include "uiohook.h"
*/
import "C"

var (
	evCh = make(chan *C.uiohook_event, 1024)
)

//export goDispatchProc
func goDispatchProc(event *C.uiohook_event) {
	evCh <- event
}
