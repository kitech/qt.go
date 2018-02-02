package ffiqt

/*
#include <stdint.h>

extern uint64_t callbackAllInherits(void* cbobj, char* iname, void*p0, void*p1, void*p2, void*p3, void*p4, void*p5, void*p6, void*p7, void*p8, void*p9);
*/
import "C"
import (
	"fmt"
	"gopp"
	"log"
	"reflect"
	"regexp"
	"unsafe"

	"github.com/therecipe/qt"
)

/*
 */
func SetInheritCallback2c(name string, fnptr unsafe.Pointer) {
	symname := fmt.Sprintf("set_callback%s", name)
	InvokeQtFunc6(symname, FFI_TYPE_VOID, fnptr)
}

func SetAllInheritCallback(cbobj CObjectIF, name string, f interface{}) {
	if cbobj.GetCthis() != nil {

		if signal := qt.LendSignal(cbobj.GetCthis(), name); signal != nil {
			log.Println("already exists:", name, reflect.TypeOf(cbobj))
			qt.ConnectSignal(cbobj.GetCthis(), name, func(args ...interface{}) {
				signal.(func(...interface{}))(args...)
				callbackInheritInvoke(f, args...)
			})
		} else {
			qt.ConnectSignal(cbobj.GetCthis(), name, f)
		}
	}
}

// 在go中的projected继承代理方法调用
func CallbackAllInherits(cbobj unsafe.Pointer, name string, args ...interface{}) interface{} {
	if debugDynSlot {
		log.Println(name, cbobj, len(args), args)
	}

	if signal := qt.GetSignal(cbobj, name); signal != nil {
		return signal.(func(...interface{}) interface{})(args...)
	} else {
		gopp.NilPrint(nil, "not found sigobj:", cbobj, name, args)
	}

	return nil
}

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
	gopp.Assert(len(in) == fv.Type().NumIn(), "no engouth parameters")
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

func callbackInheritInvokeDefault(f interface{}, args ...interface{}) interface{} {
	return nil
}

// 在C绑定中的projected继承代理方法直接调用
//export callbackAllInherits
func callbackAllInherits(cbobj unsafe.Pointer, iname *C.char, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 unsafe.Pointer) C.uint64_t {
	name_ := C.GoString(iname)
	return C.uint64_t(callbackAllInheritsGo(cbobj, name_, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func setAllInheritCallback2c(name string, fnptr unsafe.Pointer) {
	symname := fmt.Sprintf("set_callback%s", name)
	InvokeQtFunc6(symname, FFI_TYPE_VOID, fnptr)
}

func init() { setAllInheritCallback2c("AllInherits", C.callbackAllInherits) }

func callbackAllInheritsGo(cbobj unsafe.Pointer, name string, pargs ...unsafe.Pointer) uint64 {
	return 0
}
