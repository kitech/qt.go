package qtcore

// /usr/include/qt/QtCore/qsocketnotifier.h
// #include <qsocketnotifier.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSocketNotifier) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSocketNotifierFromPointer(cthis unsafe.Pointer) *QSocketNotifier {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSocketNotifier{bcthis0}
}
func (*QSocketNotifier) NewFromPointer(cthis unsafe.Pointer) *QSocketNotifier {
	return NewQSocketNotifierFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:50
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSocketNotifier) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsocketnotifier.h:56
// index:0
// Public
// void QSocketNotifier(qintptr, enum QSocketNotifier::Type, QObject *)
func NewQSocketNotifier(socket int64, arg1 int, parent *QObject /*777 QObject **/) *QSocketNotifier {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifierC2ExNS_4TypeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, socket, arg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQSocketNotifierFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsocketnotifier.h:57
// index:0
// Public virtual
// void ~QSocketNotifier()
func DeleteQSocketNotifier(*QSocketNotifier) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifierD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:59
// index:0
// Public
// qintptr socket()
func (this *QSocketNotifier) Socket() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier6socketEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qsocketnotifier.h:60
// index:0
// Public
// QSocketNotifier::Type type()
func (this *QSocketNotifier) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:62
// index:0
// Public
// bool isEnabled()
func (this *QSocketNotifier) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSocketNotifier9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsocketnotifier.h:65
// index:0
// Public
// void setEnabled(_Bool)
func (this *QSocketNotifier) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifier10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:71
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QSocketNotifier) Event(arg0 *QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSocketNotifier5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QSocketNotifier__Type = int

const QSocketNotifier__Read QSocketNotifier__Type = 0
const QSocketNotifier__Write QSocketNotifier__Type = 1
const QSocketNotifier__Exception QSocketNotifier__Type = 2

//  body block end
