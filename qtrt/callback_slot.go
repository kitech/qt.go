package qtrt

/*
#include <stdint.h>

static void *derefp4dynslotcbarg(void *ptr) { return *((void**)ptr);}
*/
import "C"
import (
	"log"
	"reflect"
	"regexp"
	"unsafe"

	qt "github.com/kitech/qt.go/qtqt"
)

// see qobjectdefs.h:263
func QMETHOD(mthName string) string   { return "0" + mthName }
func QSLOT(slotName string) string    { return "1" + slotName }
func QSIGNAL(sigName string) string   { return "2" + sigName }
func UNQSIGNAL(sigName string) string { return sigName[1:] }
func UNQSLOT(slotName string) string  { return slotName[1:] }
func UNQMETHOD(mthName string) string { return mthName[1:] }

func DerefPtr2(ptr2 unsafe.Pointer) unsafe.Pointer { return C.derefp4dynslotcbarg(ptr2) }

// export为C的函数，并取该export函数的C地址：
// 在另一文件中采用extern，不能在当前文件中extern，否则会出现符号冲突

//export callbackAllQDynSlotObject
func callbackAllQDynSlotObject(obj unsafe.Pointer, callType C.int, callId C.int, argvals **C.uchar,
	name *C.char, argc C.int, argtys *C.int, cbptr unsafe.Pointer) {
	if debugDynSlot {
		log.Println("dynslot gocbslotfunc triggered:", obj, callType, callId, argvals)
		log.Println(C.GoString(name), argc, argtys, cbptr)
	}

	var argtysSlice = (*[32]C.int)(unsafe.Pointer(argtys))
	if debugDynSlot {
		log.Println(argtysSlice, len(argtysSlice))
	}

	callbackAllQDynSlotObjectGo(obj, int(callType), int(callId), argvals,
		C.GoString(name), int(argc),
		(*int)(unsafe.Pointer(argtys)), cbptr)
}

func callbackAllQDynSlotObjectGo(obj unsafe.Pointer, callType int, callId int, argvals **C.uchar,
	name string, argc int, argtys *int, cbptr unsafe.Pointer) {
	if debugDynSlot {
		log.Println(obj, callType, callId, argvals)
		log.Println(name, argc, argtys, cbptr)
	}

	dsobjx := qt.LendSignal(obj, name)
	NilPrint(dsobjx, "not found sigobj:", obj, name)
	if dsobjx != nil {
		dsobj := dsobjx.(*QDynSlotObject)
		_ = dsobj
		if signal := qt.LendSignal(dsobj.sigsrc, UNQSIGNAL(name)); signal != nil {
			if false {
				signal.(func())()
			}
			signal.(func(unsafe.Pointer, interface{}))(unsafe.Pointer(argvals), dsobj.sigobj)
			// callbackInvoke(signal, argvals, dsobj.sigobj)
		} else {
			log.Println("not found signal: ", obj, UNQSIGNAL(name))
		}
	}
}

type callbackSlotInvoker func(argvals /* **C.uchar*/ unsafe.Pointer, sigobj interface{})

/*
真正回调函数的调用执行
f 函数
sigobj 是一个Qxxx的go实例，可以通过这个实例查询到signal的参数个数与参数类型信息
params:
  Ptr: caller should be Qxxx* qt class pointer
  string: caller should be QString&
  bool: caller should be C++'s bool or int8 or char that have 1B length
  unsafe.Pointer: caller should be a raw pointer
*/
func callbackSlotInvoke(f interface{}, argvalsp /* **C.uchar*/ unsafe.Pointer, sigobj interface{}) {
	sigobjv := reflect.ValueOf(sigobj)
	fv := reflect.ValueOf(f)

	if debugDynSlot {
		log.Println(sigobjv.Kind().String(), sigobjv.String())
	}
	Assert(fv.Kind() == reflect.Func, "cbobj must func", fv.Kind().String(), fv.String())

	// signal函数的参数类型一般还是比较简单的，像整数，布尔，QString
	in := []reflect.Value{} // 构造这个call in
	argvals := (*[30]unsafe.Pointer)(argvalsp)
	for i := 0; i < fv.Type().NumIn(); i++ {
		argty := fv.Type().In(i)
		if debugDynSlot {
			log.Println(i, argty.String(), argty.Kind())
			log.Println(i, argvals[i+1], reflect.TypeOf(argvals[i+1]).String())
			log.Println(i, &argvals[i+1])
		}

		switch argty.Kind() {
		case reflect.Int:
			in = append(in, reflect.ValueOf(int(*(*C.int)(argvals[i+1]))))
		case reflect.Bool: // it's come from qt's *bool, only 1B
			prmval := IfElse(int8(*(*C.int8_t)(argvals[i+1])) == 0, false, true).(bool)
			in = append(in, reflect.ValueOf(prmval))
		case reflect.Ptr:
			argd1ty := argty.Elem()
			switch argd1ty.Kind() {
			case reflect.Int:
			case reflect.Ptr:
			default:
				reg := regexp.MustCompile(`^(qt.*\.)?Q[A-Z](.+)`)
				if reg.MatchString(argd1ty.String()) {
					if debugDynSlot {
						log.Println("qt class pointer:", argd1ty.String())
					}
					prmval := reflect.New(argd1ty)
					setmthv := prmval.MethodByName("SetCthis")
					tmpmthv := prmval.MethodByName("QObject_PTR")
					if tmpmthv.IsValid() { // QObject or subclas
						tmpval := C.derefp4dynslotcbarg(argvals[i+1])
						setmthv.Call([]reflect.Value{reflect.ValueOf(tmpval)})
					} else {
						setmthv.Call([]reflect.Value{reflect.ValueOf(argvals[i+1])})
					}
					in = append(in, prmval)
				}
			}
		case reflect.UnsafePointer:
			tmpval := C.derefp4dynslotcbarg(argvals[i+1])
			in = append(in, reflect.ValueOf(tmpval))
		case reflect.String: // think caller supply a QString*
			var this_ unsafe.Pointer = argvals[i+1]
			ret1, err := InvokeQtFunc6("_ZNKR7QString6toUtf8Ev", FFI_TYPE_POINTER, this_)
			ErrPrint(err, this_)
			ret2, err := InvokeQtFunc6("_ZN10QByteArray4dataEv", FFI_TYPE_POINTER, unsafe.Pointer(uintptr(ret1)))
			ErrPrint(err, ret1)
			prmval := reflect.ValueOf(GoStringI(uint64(ret2)))
			in = append(in, prmval)
			InvokeQtFunc6("_ZN10QByteArrayD2Ev", FFI_TYPE_VOID, unsafe.Pointer(uintptr(ret1)))
		case reflect.Float64:
			in = append(in, reflect.ValueOf(float64(*(*C.double)(argvals[i+1]))))
		case reflect.Float32:
			in = append(in, reflect.ValueOf(float32(*(*C.float)(argvals[i+1]))))
		case reflect.Uint16:
			in = append(in, reflect.ValueOf(uint16(*(*C.uint16_t)(argvals[i+1]))))
		case reflect.Int16:
			in = append(in, reflect.ValueOf(int16(*(*C.int16_t)(argvals[i+1]))))
		case reflect.Uint32:
			in = append(in, reflect.ValueOf(uint32(*(*C.uint32_t)(argvals[i+1]))))
		case reflect.Int32:
			in = append(in, reflect.ValueOf(int32(*(*C.int32_t)(argvals[i+1]))))
		case reflect.Uint64:
			in = append(in, reflect.ValueOf(uint64(*(*C.uint64_t)(argvals[i+1]))))
		case reflect.Int64:
			in = append(in, reflect.ValueOf(int64(*(*C.int64_t)(argvals[i+1]))))
		default:
			log.Println("Unsupported:", argty.Kind().String(), argty.String())
		}
	}

	if debugDynSlot {
		log.Println(len(in), in)
	}
	if len(in) != fv.Type().NumIn() {
		log.Printf("Can not fill enough parameters, have:%d, want:%d\n", len(in), fv.Type().NumIn())
	}
	Assert(len(in) == fv.Type().NumIn(), "no engouth parameters")
	if true {
		out := fv.Call(in)
		if debugDynSlot {
			log.Println(len(out), out)
		}
	}
}
