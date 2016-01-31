package qtrt

/*
#include <stdint.h>
*/
import "C"
import "unsafe"
import "reflect"
import "log"

func test_123() {
	// var a0 interface{}

}

func Byte2Charp(a0 interface{}) *C.uchar {
	var p0 = a0.([]byte)
	return (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(p0).UnsafeAddr())) // OK
}

/*
cannot use arg1 (type unsafe.Pointer) as type **_Ctype_float
cannot use arg0 (type *_Ctype_char) as type *_Ctype_wchar_t
*/

func Rune2WCharp(a0 interface{}) *C.wchar_t {
	return (*C.wchar_t)((unsafe.Pointer)(reflect.ValueOf(a0.([]rune)).UnsafeAddr())) // OK
	// return (*C.wchar_t)(a0.([]rune))
}

func Float2Float(a0 interface{}) **C.float {
	return (**C.float)((unsafe.Pointer)(reflect.ValueOf(a0.([][]float32)).UnsafeAddr())) // OK
}

func Double2Double(a0 interface{}) **C.double {
	return (**C.double)((unsafe.Pointer)(reflect.ValueOf(a0.([][]float64)).UnsafeAddr())) // OK
}

func HandyConvert2c(fv interface{}, to reflect.Type) (retv interface{}, free bool) {
	from := reflect.TypeOf(fv)
	switch {
	case from.Kind() == reflect.String && to.Kind() == reflect.Ptr && to.Elem().Kind() == reflect.Uint8:
		retv = unsafe.Pointer(C.CString(fv.(string)))
		free = true
		return
	default:
	}

	switch to.Kind() {
	case reflect.Ptr:
		log.Panicln("can's convert:", from.Kind().String(), to.Kind().String(), to.Elem().Kind().String())
	default:
		log.Panicln("can's convert:", from.Kind().String(), to.Kind().String())
	}
	return
}

func HandyConvert2go(fv interface{}, to reflect.Type) (retv interface{}, free bool) {
	from := reflect.TypeOf(fv)

	switch to.Kind() {
	case reflect.Ptr:
		log.Panicln("can's convert:", from.Kind().String(), to.Kind().String(), to.Elem().Kind().String())
	default:
		log.Panicln("can's convert:", from.Kind().String(), to.Kind().String())
	}
	return
}
