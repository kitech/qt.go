package qtrt

/*
extern void callbackAllQDynSlotObject(void*, int, int, void**, char*, int, int*, void*);
*/
import "C"
import (
	"errors"
	"log"
	"strings"
	"unsafe"

	qt "github.com/kitech/qt.go/qtqt"
)

// usage: ffiqt.Connect()
// 也许放在qtrt包中更好，但在测试时放在了ffiqt包中，暂时放在这
type QDynSlotObject struct {
	cthis unsafe.Pointer

	//state
	sigsrc unsafe.Pointer
	sigobj CObjectITF // 保存类型信息
}

func (this *QDynSlotObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return this.cthis
}

func NewQDynSlotObject(signalName string, argc int) *QDynSlotObject {
	this := &QDynSlotObject{}
	// void *fnptr, char* name, int argc, int*argtys, void* cbptr
	name_ := C.CString(signalName)
	argc_ := int(argc)
	argtys_ := (*int)(nil)
	var cbptr_ unsafe.Pointer
	fnptr_ := unsafe.Pointer(C.callbackAllQDynSlotObject)

	rv, err := InvokeQtFunc6("QDynSlotObject_new", FFI_TYPE_POINTER,
		fnptr_, name_, argc_, argtys_, cbptr_)
	ErrPrint(err, rv)

	this.cthis = unsafe.Pointer(uintptr(rv))
	SetFinalizer(this, DeleteQDynSlotObject)
	return this
}

func DeleteQDynSlotObject(o *QDynSlotObject) {
	rv, err := InvokeQtFunc6("QDynSlotObject_delete", FFI_TYPE_VOID, o.cthis)
	ErrPrint(err, rv)
}

type CObjectITF interface {
	GetCthis() unsafe.Pointer
	SetCthis(unsafe.Pointer)
}

// 可以是signal的完整函数原型，也可以是名字
func Connect(cobj CObjectITF, signame string, f interface{}) {
	((*QDynSlotObject)(nil))._Connect(cobj, signame, f)
}

// cobj sender
func (*QDynSlotObject) _Connect(cobj CObjectITF, signame string, f interface{}) {
	cptr := cobj.GetCthis()
	NilPrint(cptr, "cptr nil")

	if cptr != nil {
		if !strings.Contains(signame, "(") {
			s, err := QObjectGetSignatureByName(cobj, signame)
			ErrPrint(err)
			signame = IfElseStr(err == nil, s, signame)
		}

		// 相当于初始化信号开关
		if !qt.ExistsSignal(cptr, signame) {
			// C.QMessageBox_ConnectButtonClicked(ptr.Pointer())
			ConnectSwitch(cptr, signame, true, cobj)
		}

		if signal := qt.LendSignal(cptr, signame); signal != nil {
			qt.ConnectSignal(cptr, signame, func(argvals /* **C.uchar*/ unsafe.Pointer, sigobj interface{}) {
				signal.(func(unsafe.Pointer, interface{}))(argvals, sigobj)
				callbackSlotInvoke(f, argvals, sigobj)
				if false {
					signal.(func())()
					f.(func())()
				}
			})
		} else {
			qt.ConnectSignal(cptr, signame, func(argvals /* **C.uchar*/ unsafe.Pointer, sigobj interface{}) {
				callbackSlotInvoke(f, argvals, sigobj)
			})
		}
	}
}

// 可以是signal的完整函数原型，也可以是名字
func ConnectSwitch(src unsafe.Pointer, signame string, on bool, cobj CObjectITF) {
	((*QDynSlotObject)(nil))._ConnectSwitch(src, signame, on, cobj)
}

func (*QDynSlotObject) _ConnectSwitch(src unsafe.Pointer, signame string, on bool, cobj CObjectITF) {
	signamep := QSIGNAL(signame)
	signame_ := C.CString(signamep)
	if debugDynSlot {
		log.Println(signame_, signamep, on)
	}

	// 相当于初始化信号开关
	var subDynSlot *QDynSlotObject
	if !qt.ExistsSignal(src, signamep) {
		subDynSlot = NewQDynSlotObject(signamep, int(456))
		subDynSlot.sigsrc = src
		subDynSlot.sigobj = cobj
	}
	if signal := qt.LendSignal(src, signamep); signal != nil {
		subDynSlot = signal.(*QDynSlotObject)
	} else {
		qt.ConnectSignal(subDynSlot.cthis, signamep, subDynSlot)
	}

	if on {
		rv, err := InvokeQtFunc6("QDynSlotObject_connect_switch", FFI_TYPE_VOID,
			src, signame_, subDynSlot.cthis, int(1))
		ErrPrint(err, rv)
	} else {
		rv, err := InvokeQtFunc6("QDynSlotObject_connect_switch", FFI_TYPE_VOID,
			src, signame_, subDynSlot.cthis, int(0))
		ErrPrint(err, rv)
	}
}

// TODO
func Disconnect(cobj CObjectITF, signame string) {
	qt.DisconnectAllSignals(cobj.GetCthis(), signame)
}

// or direct use qtcore.QObject_Connect
// only for support signal to signal
func ConnectSignal(sender CObjectITF, signal string, receiver CObjectITF, member string) {
	signal, member = autoCompleteSignature(sender, signal), autoCompleteSignature(receiver, member)

	connectqq_impl(sender, QSIGNAL(signal), receiver, QSIGNAL(member))
}

// quickly connect an object's signal to another object's slot
func ConnectSlot(sender CObjectITF, signal string, receiver CObjectITF, member string) {
	signal, member = autoCompleteSignature(sender, signal), autoCompleteSignature(receiver, member)

	connectqq_impl(sender, QSIGNAL(signal), receiver, QSLOT(member))
}

func autoCompleteSignature(sender CObjectITF, signal string) string {
	if !strings.Contains(signal, "(") {
		s, err := QObjectGetSignatureByName(sender, signal)
		ErrPrint(err)
		signal = IfElseStr(err == nil, s, signal)
	}
	return signal
}

func connectqq_impl(sender CObjectITF, signal string, receiver CObjectITF, member string) {

	var convArg0 = sender.GetCthis()
	var convArg1 = CString(signal)
	defer FreeMem(convArg1)
	var convArg2 = receiver.GetCthis()
	var convArg3 = CString(member)
	defer FreeMem(convArg3)
	var arg4 int = 2 // Qt::QueuedConnection	2
	rv, err := InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	ErrPrint(err, rv)

	// manually fix memory leak of return
	InvokeQtFunc6("_ZN11QMetaObject10ConnectionD2Ev", FFI_TYPE_VOID, unsafe.Pointer(uintptr(rv)))
	// return int(rv)
}

// 这个可以简化connect时的书写。仅在只有一个同名的情况下适用。
func QObjectGetSignatureByName(qobj CObjectITF, name string) (string, error) {
	var convArg1 = CString(name)
	defer FreeMem(convArg1)
	rv, err := InvokeQtFunc6("QObject_get_meta_signature_by_name", FFI_TYPE_POINTER, qobj.GetCthis(), convArg1)
	ErrPrint(err, rv)
	switch rv {
	case 0:
		return "", errors.New("not found")
	case 1:
		return "", errors.New("not unique")
	default:
		defer FreeMemI(rv)
		return GoStringI(rv), nil
	}
}

// 可以是signal的完整函数原型，也可以是名字
// 仅供内部使用
// TODO use only one QDynSlotObject
var destroyedDynSlot *QDynSlotObject
var destroyedSingalName = "destroyed(QObject*)"
var destroyedSingalNamep = QSIGNAL(destroyedSingalName)

func init_destroyedDynSlot() { destroyedDynSlot = NewQDynSlotObject("destroyed(QObject*)", 0) }
func ConnectDestroyed(senderCobj CObjectITF, className string) {
	f := getOnQObjectDestroyed(senderCobj, className)
	((*QDynSlotObject)(nil))._Connect(senderCobj, destroyedSingalName, f)
}

// disconnect obj's connected signal and cleanup
func getOnQObjectDestroyed(senderCobj CObjectITF, className string) func(unsafe.Pointer) {
	return func(senderCobj1 unsafe.Pointer) {
		beforeDestroyedQObject(senderCobj.GetCthis(), nil, className)
		if senderCobj1 != senderCobj.GetCthis() {
			log.Println("wtf")
		}
		var subDynSlot *QDynSlotObject
		var subDynSlotCthis unsafe.Pointer
		if signal := qt.LendSignal(senderCobj.GetCthis(), destroyedSingalNamep); signal != nil {
			subDynSlot = signal.(*QDynSlotObject)
			subDynSlotCthis = subDynSlot.GetCthis()
			if false { // because it's automatically when new, not need
				DeleteQDynSlotObject(subDynSlot)
			}
			subDynSlot.sigobj.SetCthis(nil)
		} else {
			log.Println("wtf")
		}
		qt.DisconnectAllSignals(senderCobj.GetCthis(), destroyedSingalNamep)
		afterDestroyedQObject(senderCobj1, subDynSlotCthis, className)
	}
}

var beforeDestroyedQObject = func(senderCobj, recverCobj unsafe.Pointer, className string) {
	if debugDynSlot {
		log.Println(senderCobj, recverCobj, className)
	}
}
var afterDestroyedQObject = func(senderCobj, recverCobj unsafe.Pointer, className string) {
	if debugDynSlot {
		log.Println(senderCobj, recverCobj, className)
	}
}

//
var debugDynSlot bool = false

func SetDebugDynSlot(on bool) {
	debugDynSlot = on
	InvokeQtFunc6("QDynSlotObject_set_debug", FFI_TYPE_VOID, IfElseInt(on, 1, 0))
}

// test
func (this *QDynSlotObject) Connect_test(to *QDynSlotObject) {
	rv, err := InvokeQtFunc6("QDynSlotObject_connect_test", FFI_TYPE_VOID, this.cthis, to.cthis)
	ErrPrint(err, rv)
}

func (this *QDynSlotObject) Trigger_test() {
	rv, err := InvokeQtFunc6("QDynSlotObject_trigger_test", FFI_TYPE_VOID, this.cthis)
	ErrPrint(err, rv)
}
