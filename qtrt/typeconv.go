package qtrt

/*
#include <stdlib.h>
#include <stdint.h>

static void carr_set_item(char** pp, int idx, char*p)
{ pp[idx] = p; }
static char* carr_get_item(char** pp, int idx)
{ return pp[idx]; }
*/
import "C"
import "unsafe"
import "reflect"
import "log"
import "math"

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

func HandyConvert2go(fvi interface{}, to reflect.Type) (reti interface{}) {
	from := reflect.TypeOf(fvi)
	fv := reflect.ValueOf(fvi)

	switch {
	case from.Kind() == reflect.UnsafePointer && to.Kind() == reflect.Struct:
		retv := reflect.New(to)
		retv.Elem().FieldByName("Qclsinst").Set(fv)
		reti = retv.Interface()
	}

	return
}

// for argument of primitive type
func PrimConv(fvi interface{}, to reflect.Type) (reti interface{}) {
	from := reflect.TypeOf(fvi)
	fv := reflect.ValueOf(fvi)

	reti = fvi
	if from.ConvertibleTo(to) {
		reti = fv.Convert(to).Interface()
	}

	return
}

func StringSliceToCCharPP(ss []string) unsafe.Pointer {
	var tp *C.char
	p := C.calloc(C.size_t(len(ss)+1), C.size_t(unsafe.Sizeof(tp)))
	var pp **C.char = (**C.char)(p)

	for i, _ := range ss {
		s := C.CString(ss[i])
		C.carr_set_item(pp, C.int(i), s)
		C.carr_set_item(pp, C.int(i+1), nil)
	}

	return p
}

func CCharPPToStringSlice(charpp unsafe.Pointer) []string {
	ss := []string{}
	var pp **C.char = (**C.char)(charpp)
	for i := 0; i < math.MaxInt32; i++ {
		p := C.carr_get_item(pp, C.int(i))
		if p == nil {
			break
		}
		ss = append(ss, C.GoString(p))
	}
	return ss
}

func Cretval2go(ty string, rv uint64) interface{} {
	switch ty {
	case "int":
		return int((C.int)(rv))
	case "float64":
		return float64(*(*C.double)(unsafe.Pointer(&rv)))
	default:
		log.Println("Unknown type:", ty)
	}
	return rv
}

func Cpretval2go(ty string, rv uint64) interface{} {
	switch ty {
	case "int":
		return int(*(*C.int)(unsafe.Pointer(uintptr(rv))))
	case "float64":
		return float64(*(*C.double)(unsafe.Pointer(uintptr(rv))))
	default:
		log.Println("Unknown type:", ty)
	}
	return rv
}

func GoBool2C(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CBool2Go(b int) bool {
	if b == 0 {
		return false
	}
	return true
}
