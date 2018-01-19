//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QSwipeGesture struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgesture.h:209
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSwipeGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:221
// index:0
// void QSwipeGesture(class QObject *)
func NewQSwipeGesture(parent unsafe.Pointer) *QSwipeGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSwipeGesture{cthis}
}

// /usr/include/qt/QtWidgets/qgesture.h:222
// index:0
// virtual
// void ~QSwipeGesture()
func DeleteQSwipeGesture(*QSwipeGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:224
// index:0
// QSwipeGesture::SwipeDirection horizontalDirection()
func (this *QSwipeGesture) HorizontalDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture19horizontalDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:225
// index:0
// QSwipeGesture::SwipeDirection verticalDirection()
func (this *QSwipeGesture) VerticalDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture17verticalDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:227
// index:0
// qreal swipeAngle()
func (this *QSwipeGesture) SwipeAngle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture10swipeAngleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:228
// index:0
// void setSwipeAngle(qreal)
func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	// 0: (, qreal value), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGesture13setSwipeAngleEd", ffiqt.FFI_TYPE_VOID, this.cthis, &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
