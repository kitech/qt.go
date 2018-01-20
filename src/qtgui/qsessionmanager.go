//  header block begin
// /usr/include/qt/QtGui/qsessionmanager.h
// #include <qsessionmanager.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
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

// /usr/include/qt/QtGui/qsessionmanager.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSessionManager) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:65
// index:0
// QString sessionId()
func (this *QSessionManager) SessionId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager9sessionIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:66
// index:0
// QString sessionKey()
func (this *QSessionManager) SessionKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager10sessionKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:68
// index:0
// bool allowsInteraction()
func (this *QSessionManager) AllowsInteraction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17allowsInteractionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:69
// index:0
// bool allowsErrorInteraction()
func (this *QSessionManager) AllowsErrorInteraction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager22allowsErrorInteractionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:70
// index:0
// void release()
func (this *QSessionManager) Release() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager7releaseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:72
// index:0
// void cancel()
func (this *QSessionManager) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager6cancelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:80
// index:0
// void setRestartHint(enum QSessionManager::RestartHint)
func (this *QSessionManager) SetRestartHint(arg0 int) {
	// 0: (, QSessionManager::RestartHint arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager14setRestartHintENS_11RestartHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:81
// index:0
// QSessionManager::RestartHint restartHint()
func (this *QSessionManager) RestartHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager11restartHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:83
// index:0
// void setRestartCommand(const class QStringList &)
func (this *QSessionManager) SetRestartCommand(arg0 unsafe.Pointer) {
	// 0: (, const QStringList & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17setRestartCommandERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:84
// index:0
// QStringList restartCommand()
func (this *QSessionManager) RestartCommand() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager14restartCommandEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:85
// index:0
// void setDiscardCommand(const class QStringList &)
func (this *QSessionManager) SetDiscardCommand(arg0 unsafe.Pointer) {
	// 0: (, const QStringList & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager17setDiscardCommandERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:86
// index:0
// QStringList discardCommand()
func (this *QSessionManager) DiscardCommand() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager14discardCommandEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:88
// index:0
// void setManagerProperty(const class QString &, const class QString &)
func (this *QSessionManager) SetManagerProperty(name unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, name const QString &, value const QString &), (name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:89
// index:1
// void setManagerProperty(const class QString &, const class QStringList &)
func (this *QSessionManager) SetManagerProperty_1(name unsafe.Pointer, value unsafe.Pointer) {
	// 1: (, name const QString &, value const QStringList &), (name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:91
// index:0
// bool isPhase2()
func (this *QSessionManager) IsPhase2() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSessionManager8isPhase2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:92
// index:0
// void requestPhase2()
func (this *QSessionManager) RequestPhase2() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSessionManager13requestPhase2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
