package qtgui

// /usr/include/qt/QtGui/qsessionmanager.h
// #include <qsessionmanager.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSessionManager struct {
	*qtcore.QObject
}

func (this *QSessionManager) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSessionManager) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSessionManagerFromPointer(cthis unsafe.Pointer) *QSessionManager {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSessionManager{bcthis0}
}
func (*QSessionManager) NewFromPointer(cthis unsafe.Pointer) *QSessionManager {
	return NewQSessionManagerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsessionmanager.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSessionManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qsessionmanager.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionId()
func (this *QSessionManager) SessionId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager9sessionIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qsessionmanager.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionKey()
func (this *QSessionManager) SessionKey() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager10sessionKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qsessionmanager.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allowsInteraction()
func (this *QSessionManager) AllowsInteraction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17allowsInteractionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allowsErrorInteraction()
func (this *QSessionManager) AllowsErrorInteraction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager22allowsErrorInteractionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void release()
func (this *QSessionManager) Release() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager7releaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()
func (this *QSessionManager) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRestartHint(enum QSessionManager::RestartHint)
func (this *QSessionManager) SetRestartHint(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager14setRestartHintENS_11RestartHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QSessionManager::RestartHint restartHint()
func (this *QSessionManager) RestartHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager11restartHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRestartCommand(const QStringList &)
func (this *QSessionManager) SetRestartCommand(arg0 *qtcore.QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17setRestartCommandERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDiscardCommand(const QStringList &)
func (this *QSessionManager) SetDiscardCommand(arg0 *qtcore.QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17setDiscardCommandERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManagerProperty(const QString &, const QString &)
func (this *QSessionManager) SetManagerProperty(name string, value string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setManagerProperty(const QString &, const QStringList &)
func (this *QSessionManager) SetManagerProperty_1(name string, value *qtcore.QStringList) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPhase2()
func (this *QSessionManager) IsPhase2() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager8isPhase2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestPhase2()
func (this *QSessionManager) RequestPhase2() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager13requestPhase2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

func DeleteQSessionManager(this *QSessionManager) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManagerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QSessionManager__RestartHint = int

const QSessionManager__RestartIfRunning QSessionManager__RestartHint = 0
const QSessionManager__RestartAnyway QSessionManager__RestartHint = 1
const QSessionManager__RestartImmediately QSessionManager__RestartHint = 2
const QSessionManager__RestartNever QSessionManager__RestartHint = 3

//  body block end
