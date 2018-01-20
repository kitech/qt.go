//  header block begin
// /usr/include/qt/QtGui/qsessionmanager.h
// #include <qsessionmanager.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
		ffiqt.KeepMe()
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
	return this.QObject.GetCthis()
}
func NewQSessionManagerFromPointer(cthis unsafe.Pointer) *QSessionManager {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSessionManager{bcthis0}
}

// /usr/include/qt/QtGui/qsessionmanager.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSessionManager) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:65
// index:0
// Public
// QString sessionId()
func (this *QSessionManager) SessionId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager9sessionIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:66
// index:0
// Public
// QString sessionKey()
func (this *QSessionManager) SessionKey() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager10sessionKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:68
// index:0
// Public
// bool allowsInteraction()
func (this *QSessionManager) AllowsInteraction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17allowsInteractionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:69
// index:0
// Public
// bool allowsErrorInteraction()
func (this *QSessionManager) AllowsErrorInteraction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager22allowsErrorInteractionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:70
// index:0
// Public
// void release()
func (this *QSessionManager) Release() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager7releaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:72
// index:0
// Public
// void cancel()
func (this *QSessionManager) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:80
// index:0
// Public
// void setRestartHint(enum QSessionManager::RestartHint)
func (this *QSessionManager) SetRestartHint(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager14setRestartHintENS_11RestartHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:81
// index:0
// Public
// QSessionManager::RestartHint restartHint()
func (this *QSessionManager) RestartHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager11restartHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:83
// index:0
// Public
// void setRestartCommand(const class QStringList &)
func (this *QSessionManager) SetRestartCommand(arg0 *qtcore.QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17setRestartCommandERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:84
// index:0
// Public
// QStringList restartCommand()
func (this *QSessionManager) RestartCommand() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager14restartCommandEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:85
// index:0
// Public
// void setDiscardCommand(const class QStringList &)
func (this *QSessionManager) SetDiscardCommand(arg0 *qtcore.QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17setDiscardCommandERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:86
// index:0
// Public
// QStringList discardCommand()
func (this *QSessionManager) DiscardCommand() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager14discardCommandEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:88
// index:0
// Public
// void setManagerProperty(const class QString &, const class QString &)
func (this *QSessionManager) SetManagerProperty(name *qtcore.QString, value *qtcore.QString) {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:89
// index:1
// Public
// void setManagerProperty(const class QString &, const class QStringList &)
func (this *QSessionManager) SetManagerProperty_1(name *qtcore.QString, value *qtcore.QStringList) {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:91
// index:0
// Public
// bool isPhase2()
func (this *QSessionManager) IsPhase2() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager8isPhase2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsessionmanager.h:92
// index:0
// Public
// void requestPhase2()
func (this *QSessionManager) RequestPhase2() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager13requestPhase2Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
