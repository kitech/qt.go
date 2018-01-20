//  header block begin
// /usr/include/qt/QtCore/qsocketnotifier.h
// #include <qsocketnotifier.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QSocketNotifier struct {
	*QObject
}

func (this *QSocketNotifier) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qsocketnotifier.h:50
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSocketNotifier) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:56
// index:0
// void QSocketNotifier(qintptr, enum QSocketNotifier::Type, class QObject *)
func NewQSocketNotifier(socket int64, arg1 int, parent unsafe.Pointer) *QSocketNotifier {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifierC2ExNS_4TypeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &socket, &arg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSocketNotifierFromPointer(cthis)
	return gothis
}
func NewQSocketNotifierFromPointer(cthis unsafe.Pointer) *QSocketNotifier {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSocketNotifier{bcthis0}
}

// /usr/include/qt/QtCore/qsocketnotifier.h:57
// index:0
// virtual
// void ~QSocketNotifier()
func DeleteQSocketNotifier(*QSocketNotifier) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifierD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:59
// index:0
// qintptr socket()
func (this *QSocketNotifier) Socket() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier6socketEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:60
// index:0
// QSocketNotifier::Type type()
func (this *QSocketNotifier) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:62
// index:0
// bool isEnabled()
func (this *QSocketNotifier) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:65
// index:0
// void setEnabled(_Bool)
func (this *QSocketNotifier) SetEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifier10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:71
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSocketNotifier) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifier5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
