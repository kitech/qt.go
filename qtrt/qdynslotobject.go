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
	sigobj interface{} // 保存类型信息
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

type CObjectIF interface {
	GetCthis() unsafe.Pointer
	SetCthis(unsafe.Pointer)
}

// 可以是signal的完整函数原型，也可以是名字
func Connect(cobj CObjectIF, signame string, f interface{}) {
	((*QDynSlotObject)(nil))._Connect(cobj, signame, f)
}

func (*QDynSlotObject) _Connect(cobj CObjectIF, signame string, f interface{}) {
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
func ConnectSwitch(src unsafe.Pointer, signame string, on bool, cobj CObjectIF) {
	((*QDynSlotObject)(nil))._ConnectSwitch(src, signame, on, cobj)
}

func (*QDynSlotObject) _ConnectSwitch(src unsafe.Pointer, signame string, on bool, cobj CObjectIF) {
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
func Disconnect(cobj CObjectIF, signame string) {
	qt.DisconnectAllSignals(cobj.GetCthis(), signame)
}

// or direct use qtcore.QObject_Connect
func ConnectSignal(sender CObjectIF, signal string, receiver CObjectIF, member string) {
	if !strings.Contains(signal, "(") {
		s, err := QObjectGetSignatureByName(sender, signal)
		ErrPrint(err)
		signal = IfElseStr(err == nil, s, signal)
	}
	if !strings.Contains(member, "(") {
		s, err := QObjectGetSignatureByName(receiver, member)
		ErrPrint(err)
		member = IfElseStr(err == nil, s, member)
	}

	var convArg0 = sender.GetCthis()
	var convArg1 = CString(QSIGNAL(signal))
	defer FreeMem(convArg1)
	var convArg2 = receiver.GetCthis()
	var convArg3 = CString(QSIGNAL(member))
	defer FreeMem(convArg3)
	var arg4 int = 2 // Qt::QueuedConnection	2
	rv, err := InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	ErrPrint(err, rv)
	// return int(rv)
}

// 这个可以简化connect时的书写。仅在只有一个同名的情况下适用。
func QObjectGetSignatureByName(qobj CObjectIF, name string) (string, error) {
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
