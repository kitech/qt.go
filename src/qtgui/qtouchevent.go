//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QTouchEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:947
// index:0
// virtual
// void ~QTouchEvent()
func DeleteQTouchEvent(*QTouchEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:949
// index:0
// inline
// QWindow * window()
func (this *QTouchEvent) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6windowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:950
// index:0
// inline
// QObject * target()
func (this *QTouchEvent) Target() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6targetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:954
// index:0
// inline
// Qt::TouchPointStates touchPointStates()
func (this *QTouchEvent) TouchPointStates() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent16touchPointStatesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:955
// index:0
// inline
// const QList<QTouchEvent::TouchPoint> & touchPoints()
func (this *QTouchEvent) TouchPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent11touchPointsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:956
// index:0
// inline
// QTouchDevice * device()
func (this *QTouchEvent) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTouchEvent6deviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:959
// index:0
// inline
// void setWindow(class QWindow *)
func (this *QTouchEvent) SetWindow(awindow unsafe.Pointer) {
	// 0: (, QWindow * awindow), (awindow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setWindowEP7QWindow", ffiqt.FFI_TYPE_VOID, this.cthis, awindow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:960
// index:0
// inline
// void setTarget(class QObject *)
func (this *QTouchEvent) SetTarget(atarget unsafe.Pointer) {
	// 0: (, QObject * atarget), (atarget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setTargetEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, atarget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:963
// index:0
// inline
// void setDevice(class QTouchDevice *)
func (this *QTouchEvent) SetDevice(adevice unsafe.Pointer) {
	// 0: (, QTouchDevice * adevice), (adevice)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTouchEvent9setDeviceEP12QTouchDevice", ffiqt.FFI_TYPE_VOID, this.cthis, adevice)
	gopp.ErrPrint(err, rv)
}

//  body block end
