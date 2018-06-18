package qtrt

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

extern uint64_t callbackAllInherits(void* cbobj, char* iname, int*handled, int argc, uint64_t p0, uint64_t p1, uint64_t p2, uint64_t p3, uint64_t p4, uint64_t p5, uint64_t p6, uint64_t p7, uint64_t p8, uint64_t p9);
*/
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"unsafe"

	qt "github.com/kitech/qt.go/qtqt"
)

/*
 */
// depcreated
func SetInheritCallback2c(name string, fnptr unsafe.Pointer) {
	symname := fmt.Sprintf("set_callback%s", name)
	InvokeQtFunc6(symname, FFI_TYPE_VOID, fnptr)
}

func SetAllInheritCallback(cbobj CObjectITF, name string, f interface{}) {
	if cbobj.GetCthis() == nil {
		log.Println("obj is nil", name)
	}

	if cbobj.GetCthis() != nil {

		if signal := qt.LendSignal(cbobj.GetCthis(), name); signal != nil {
			log.Println("already exists:", name, reflect.TypeOf(cbobj))
			qt.ConnectSignal(cbobj.GetCthis(), name, func(args ...uint64) {
				signal.(func(...uint64))(args...)
				callbackInheritInvokeGo(f, args...)
			})
		} else {
			qt.ConnectSignal(cbobj.GetCthis(), name, f)
		}
	}
}

// depcreated
// 在go中的projected继承代理方法调用
func CallbackAllInherits(cbobj unsafe.Pointer, name string, args ...interface{}) interface{} {
	if debugDynSlot {
		log.Println(name, cbobj, len(args), args)
	}

	if signal := qt.GetSignal(cbobj, name); signal != nil {
		return signal.(func(...interface{}) interface{})(args...)
	} else {
		// NilPrint(nil, "not found sigobj:", cbobj, name, args)
	}

	return nil
}

// depcreated
func callbackInheritInvoke(f interface{}, args ...interface{}) interface{} {
	fv := reflect.ValueOf(f)
	// signal函数的参数类型一般还是比较简单的，像整数，布尔，QString
	in := []reflect.Value{} // 构造这个call in
	for i := 0; i < fv.Type().NumIn(); i++ {
		argty := fv.Type().In(i)
		if debugDynSlot {
			log.Println(i, argty.String(), argty.Kind())
			log.Println(i, args[i], reflect.TypeOf(args[i]).String())
		}

		switch argty.Kind() {
		case reflect.Int:
			in = append(in, reflect.ValueOf(args[i]))
		case reflect.Bool:
			in = append(in, reflect.ValueOf(args[i]))
		case reflect.Ptr:
			argd1ty := argty.Elem()
			switch argd1ty.Kind() {
			case reflect.Int:
			case reflect.Ptr:
			default:
				reg := regexp.MustCompile(`^(qt.*\.)?Q[A-Z](.+)`)
				if reg.MatchString(argd1ty.String()) {
					if debugDynSlot {
						log.Println("qt class pointer:")
					}
					prmval := reflect.New(argd1ty)
					prmval.MethodByName("SetCthis").Call([]reflect.Value{reflect.ValueOf(args[i])})
					in = append(in, prmval)
				}
			}
		default:
			log.Println("Unsupported:", argty.Kind().String(), argty.String())
		}
	}

	if debugDynSlot {
		log.Println(len(in), in)
	}
	if len(in) != fv.Type().NumIn() {
		log.Println("can not fill enough parameters,", len(in), fv.Type().NumIn())
	}
	Assert(len(in) == fv.Type().NumIn(), "no engouth parameters")
	if true {
		out := fv.Call(in)
		if debugDynSlot {
			log.Println(len(out), out)
		}
		if len(out) > 0 {
			return out[0].Interface()
		}
	}
	return nil
}

// depcreated
func callbackInheritInvokeDefault(f interface{}, args ...interface{}) interface{} {
	return nil
}

func setAllInheritCallback2c(name string, fnptr unsafe.Pointer) {
	symname := fmt.Sprintf("set_callback%s", name)
	InvokeQtFunc6(symname, FFI_TYPE_VOID, fnptr)
}

func init_callack_inherit() { setAllInheritCallback2c("AllInherits", C.callbackAllInherits) }

// 在C绑定中的projected继承代理方法直接调用
//export callbackAllInherits
func callbackAllInherits(cbobj unsafe.Pointer, iname *C.char, handled *C.int, argc C.int, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 C.uint64_t) C.uint64_t {
	name_ := C.GoString(iname)
	var handled_ int
	rv := C.uint64_t(callbackAllInheritsGo(cbobj, name_, &handled_, int(argc),
		uint64(p0), uint64(p1), uint64(p2), uint64(p3), uint64(p4),
		uint64(p5), uint64(p6), uint64(p7), uint64(p8), uint64(p9)))
	*handled = C.int(handled_)
	return C.uint64_t(rv)
}

// return true to show inherit debug info
/*
Example:
debugAllInheritsFilterFunc = func(name string) bool {
    return name != "event"
}
*/
var debugAllInheritsFilterFunc func(name string) bool

func SetDebugAllinheritsFilterFunc(f func(name string) bool) {
	debugAllInheritsFilterFunc = f
}

func callbackAllInheritsGo(cbobj unsafe.Pointer, name string, handled *int, argc int, pargs ...uint64) uint64 {
	if !strings.Contains(strings.ToLower(name), "event") {
		if debugFFICall {
			log.Println(cbobj, name, handled, argc, pargs)
		}
	}
	if debugAllInheritsFilterFunc != nil && debugAllInheritsFilterFunc(name) {
		if debugFFICall {
			log.Println(cbobj, name, handled, argc, pargs)
		}
	}

	if signal := qt.GetSignal(cbobj, name); signal != nil {
		*handled = 1
		rv := callbackInheritInvokeGo(signal, pargs[:argc]...)
		if rv == nil {
			return 0
		}
		retv := reflect.ValueOf(rv)
		switch retv.Kind() { // TODO
		case reflect.Uint64, reflect.Int, reflect.Uint, reflect.Uint16, reflect.Int16,
			reflect.Int32, reflect.Uint32, reflect.Int8, reflect.Uint8:
			return retv.Convert(reflect.TypeOf(uint64(0))).Interface().(uint64)
		case reflect.Bool:
			return uint64(IfElseInt(rv.(bool), 1, 0))
		case reflect.Float64: // NOTE: need caller copy out
			return Float64AsInt(rv.(float64))
		case reflect.Float32:
			return Float32AsInt(rv.(float32))
		case reflect.UnsafePointer:
			// impled
			var co CObjectITF
			if retv.Type().Implements(reflect.TypeOf(co)) {
				crvpx := retv.MethodByName("GetCthis").Call([]reflect.Value{})
				return uint64(uintptr(crvpx[0].Interface().(unsafe.Pointer)))
			} else {
				return uint64(uintptr(rv.(unsafe.Pointer)))
			}
		default:
			return rv.(uint64)
		}
	} else {
		// 这个函数还是比较可靠的，没有找到那一定是没有了
		// NilPrint(nil, "not found sigobj:", cbobj, name, pargs)
	}

	return 0
}

func callbackInheritInvokeGo(f interface{}, args ...uint64) interface{} {
	fv := reflect.ValueOf(f)
	// signal函数的参数类型一般还是比较简单的，像整数，布尔，QString
	in := []reflect.Value{} // 构造这个call in
	for i := 0; i < fv.Type().NumIn(); i++ {
		argty := fv.Type().In(i)
		if debugDynSlot {
			log.Println(i, argty.String(), argty.Kind())
			log.Println(i, args[i], reflect.TypeOf(args[i]).String())
		}

		switch argty.Kind() {
		case reflect.Int:
			in = append(in, reflect.ValueOf(int(args[i])))
		case reflect.Bool:
			in = append(in, reflect.ValueOf(args[i] == 1))
		case reflect.Float64:
			in = append(in, reflect.ValueOf((float64)(*(*C.double)(unsafe.Pointer(uintptr(args[i]))))))
		case reflect.Float32:
			in = append(in, reflect.ValueOf((float32)(*(*C.float)(unsafe.Pointer(uintptr(args[i]))))))
		case reflect.Ptr:
			argd1ty := argty.Elem()
			switch argd1ty.Kind() {
			case reflect.Int:
			case reflect.Ptr:
			default:
				reg := regexp.MustCompile(`^(qt.*\.)?Q[A-Z](.+)`)
				if reg.MatchString(argd1ty.String()) {
					if debugDynSlot {
						log.Println("qt class pointer:")
					}
					prmval := reflect.New(argd1ty)
					prmval.MethodByName("SetCthis").Call([]reflect.Value{
						reflect.ValueOf(unsafe.Pointer(uintptr(args[i])))})
					in = append(in, prmval)
				}
			}
		default:
			log.Println("Unsupported:", argty.Kind().String(), argty.String())
		}
	}

	if debugDynSlot {
		log.Println(len(in), in)
	}
	if len(in) != fv.Type().NumIn() {
		log.Println("can not fill enough parameters,", len(in), fv.Type().NumIn())
	}
	Assert(len(in) == fv.Type().NumIn(), "no engouth parameters")
	if true {
		out := fv.Call(in)
		if debugDynSlot {
			log.Println(len(out), out)
		}
		if len(out) > 0 {
			return out[0].Interface()
		}
	}
	return nil
}

// 把浮点数存储在uint64中。注意这不是强制类型转换，是原样存储，不丢失精度
func Float64AsInt(n float64) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 8)
	return
}
func Float32AsInt(n float32) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 4)
	return
}
func IntAsFloat64(v uint64) (n float64) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 8)
	return
}
func IntAsFloat32(v uint64) (n float32) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 4)
	return
}
