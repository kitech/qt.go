//  header block begin
// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QEventLoop struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qeventloop.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QEventLoop) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QEventLoop10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:56
// index:0
// void QEventLoop(class QObject *)
func NewQEventLoop(parent unsafe.Pointer) *QEventLoop {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoopC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QEventLoop{cthis}
}

// /usr/include/qt/QtCore/qeventloop.h:57
// index:0
// virtual
// void ~QEventLoop()
func DeleteQEventLoop(*QEventLoop) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoopD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:74
// index:0
// void exit(int)
func (this *QEventLoop) Exit(returnCode int) {
	// 0: (, int returnCode), (&returnCode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoop4exitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &returnCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:75
// index:0
// bool isRunning()
func (this *QEventLoop) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QEventLoop9isRunningEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:77
// index:0
// void wakeUp()
func (this *QEventLoop) WakeUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoop6wakeUpEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:79
// index:0
// virtual
// bool event(class QEvent *)
func (this *QEventLoop) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoop5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:82
// index:0
// void quit()
func (this *QEventLoop) Quit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QEventLoop4quitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
