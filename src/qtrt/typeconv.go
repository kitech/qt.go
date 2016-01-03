package qtrt

import "C"
import "unsafe"
import "reflect"

func test_123() {
	// var a0 interface{}

}

func Byte2Charp(a0 interface{}) *C.uchar {
	var p0 = a0.([]byte)
	return (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(p0).UnsafeAddr())) // OK
}
