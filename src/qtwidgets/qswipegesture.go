package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

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
	*QGesture
}

func (this *QSwipeGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func NewQSwipeGestureFromPointer(cthis unsafe.Pointer) *QSwipeGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QSwipeGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:209
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSwipeGesture) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:221
// index:0
// Public
// void QSwipeGesture(class QObject *)
func NewQSwipeGesture(parent *qtcore.QObject /*444 QObject **/) *QSwipeGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSwipeGestureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:222
// index:0
// Public virtual
// void ~QSwipeGesture()
func DeleteQSwipeGesture(*QSwipeGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:224
// index:0
// Public
// QSwipeGesture::SwipeDirection horizontalDirection()
func (this *QSwipeGesture) HorizontalDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture19horizontalDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:225
// index:0
// Public
// QSwipeGesture::SwipeDirection verticalDirection()
func (this *QSwipeGesture) VerticalDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture17verticalDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:227
// index:0
// Public
// qreal swipeAngle()
func (this *QSwipeGesture) SwipeAngle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSwipeGesture10swipeAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgesture.h:228
// index:0
// Public
// void setSwipeAngle(qreal)
func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSwipeGesture13setSwipeAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
