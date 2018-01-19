//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTapAndHoldGesture struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgesture.h:254
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTapAndHoldGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:260
// index:0
// void QTapAndHoldGesture(class QObject *)
func NewQTapAndHoldGesture(parent unsafe.Pointer) *QTapAndHoldGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTapAndHoldGesture{cthis}
}

// /usr/include/qt/QtWidgets/qgesture.h:261
// index:0
// virtual
// void ~QTapAndHoldGesture()
func DeleteQTapAndHoldGesture(*QTapAndHoldGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:263
// index:0
// QPointF position()
func (this *QTapAndHoldGesture) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:264
// index:0
// void setPosition(const class QPointF &)
func (this *QTapAndHoldGesture) SetPosition(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture11setPositionERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:266
// index:0
// static
// void setTimeout(int)
func (this *QTapAndHoldGesture) SetTimeout(msecs int) {
	// 0: (int msecs), (msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture10setTimeoutEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTapAndHoldGesture_SetTimeout(msecs int) {
	// 0: (int msecs), (msecs)
	var nilthis *QTapAndHoldGesture
	nilthis.SetTimeout(msecs)
}

// /usr/include/qt/QtWidgets/qgesture.h:267
// index:0
// static
// int timeout()
func (this *QTapAndHoldGesture) Timeout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture7timeoutEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTapAndHoldGesture_Timeout() {
	// 0: (), ()
	var nilthis *QTapAndHoldGesture
	nilthis.Timeout()
}

//  body block end
