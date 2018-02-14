package qtqml

// /usr/include/qt/QtQml/qqmldebug.h
// #include <qqmldebug.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlDebuggingEnabler struct {
	*qtrt.CObject
}
type QQmlDebuggingEnabler_ITF interface {
	QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler
}

func (ptr *QQmlDebuggingEnabler) QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler { return ptr }

func (this *QQmlDebuggingEnabler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlDebuggingEnabler) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlDebuggingEnablerFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return &QQmlDebuggingEnabler{&qtrt.CObject{cthis}}
}
func (*QQmlDebuggingEnabler) NewFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return NewQQmlDebuggingEnablerFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmldebug.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlDebuggingEnabler(_Bool)
func NewQQmlDebuggingEnabler(printWarning bool) *QQmlDebuggingEnabler {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnablerC2Eb", qtrt.FFI_TYPE_POINTER, printWarning)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlDebuggingEnablerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlDebuggingEnabler)
	return gothis
}

// /usr/include/qt/QtQml/qqmldebug.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList debuggerServices()
func (this *QQmlDebuggingEnabler) DebuggerServices() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler16debuggerServicesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QQmlDebuggingEnabler_DebuggerServices() *qtcore.QStringList /*123*/ {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.DebuggerServices()
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList inspectorServices()
func (this *QQmlDebuggingEnabler) InspectorServices() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler17inspectorServicesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QQmlDebuggingEnabler_InspectorServices() *qtcore.QStringList /*123*/ {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.InspectorServices()
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:62
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList profilerServices()
func (this *QQmlDebuggingEnabler) ProfilerServices() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler16profilerServicesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QQmlDebuggingEnabler_ProfilerServices() *qtcore.QStringList /*123*/ {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.ProfilerServices()
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList nativeDebuggerServices()
func (this *QQmlDebuggingEnabler) NativeDebuggerServices() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler22nativeDebuggerServicesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QQmlDebuggingEnabler_NativeDebuggerServices() *qtcore.QStringList /*123*/ {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.NativeDebuggerServices()
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setServices(const QStringList &)
func (this *QQmlDebuggingEnabler) SetServices(services qtcore.QStringList_ITF) {
	var convArg0 = services.QStringList_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler11setServicesERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QQmlDebuggingEnabler_SetServices(services qtcore.QStringList_ITF) {
	var nilthis *QQmlDebuggingEnabler
	nilthis.SetServices(services)
}

// /usr/include/qt/QtQml/qqmldebug.h:67
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool startTcpDebugServer(int, enum QQmlDebuggingEnabler::StartMode, const QString &)
func (this *QQmlDebuggingEnabler) StartTcpDebugServer(port int, mode int, hostName string) bool {
	var tmpArg2 = qtcore.NewQString_5(hostName)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler19startTcpDebugServerEiNS_9StartModeERK7QString", qtrt.FFI_TYPE_POINTER, port, mode, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlDebuggingEnabler_StartTcpDebugServer(port int, mode int, hostName string) bool {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.StartTcpDebugServer(port, mode, hostName)
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool connectToLocalDebugger(const QString &, enum QQmlDebuggingEnabler::StartMode)
func (this *QQmlDebuggingEnabler) ConnectToLocalDebugger(socketFileName string, mode int) bool {
	var tmpArg0 = qtcore.NewQString_5(socketFileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler22connectToLocalDebuggerERK7QStringNS_9StartModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlDebuggingEnabler_ConnectToLocalDebugger(socketFileName string, mode int) bool {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.ConnectToLocalDebugger(socketFileName, mode)
	return rv
}

func DeleteQQmlDebuggingEnabler(this *QQmlDebuggingEnabler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQmlDebuggingEnablerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QQmlDebuggingEnabler__StartMode = int

const QQmlDebuggingEnabler__DoNotWaitForClient QQmlDebuggingEnabler__StartMode = 0
const QQmlDebuggingEnabler__WaitForClient QQmlDebuggingEnabler__StartMode = 1

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
