package cmemory

/*
#include <stdint.h>

extern void _cmemory_uniform_callback_0();
*/
import "C"
import "unsafe"

var cbmap4any = NewCBMap()

func AddOnceRunner(f func()) (uint64, unsafe.Pointer) {
	cbno := cbmap4any.Add(f)
	return cbno, C._cmemory_uniform_callback_0
}

//export _cmemory_uniform_callback_0
func _cmemory_uniform_callback_0(cbno C.uint64_t) {
	if fx, ok := cbmap4any.Get(uint64(cbno)); ok {
		cbmap4any.Del(uint64(cbno))
		fx.(func())()
	}
}
