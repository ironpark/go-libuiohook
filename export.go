package go_libuiohook

/*
#include "uiohook.h"
*/
import "C"
import (
	pb "google.golang.org/protobuf/proto"
	"unsafe"
)

var (
	evCh = make(chan *Msg, 1024)
)

//export goDispatchProc
func goDispatchProc(data *C.uint8_t, size C.size_t) {
	goSlice := C.GoBytes(unsafe.Pointer(data), C.int(size))
	msg := Msg{}
	if err := pb.Unmarshal(goSlice, &msg); err == nil {
		evCh <- &msg
	}
}
