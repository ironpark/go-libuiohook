package go_libuiohook

/*
#include "uiohook.h"
*/
import "C"

var (
	evCh = make(chan *Event, 1024)
)

//export goDispatchProc
func goDispatchProc(event *C.uiohook_event) {
	evCh <- cvtEvent(event)
}
