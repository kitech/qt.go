package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
func (this *QSwipeGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQSwipeGestureFromPointer(cthis unsafe.Pointer) *QSwipeGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QSwipeGesture{bcthis0}
}
func (*QSwipeGesture) NewFromPointer(cthis unsafe.Pointer) *QSwipeGesture {
	return NewQSwipeGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSwipeGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSwipeGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSwipeGesture(QObject *)
func NewQSwipeGesture(parent *qtcore.QObject /*777 QObject **/) *QSwipeGesture {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSwipeGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSwipeGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:222
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSwipeGesture()
func DeleteQSwipeGesture(this *QSwipeGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSwipeGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:224
// index:0
// Public Visibility=Default Availability=Available
// [4] QSwipeGesture::SwipeDirection horizontalDirection()
func (this *QSwipeGesture) HorizontalDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSwipeGesture19horizontalDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] QSwipeGesture::SwipeDirection verticalDirection()
func (this *QSwipeGesture) VerticalDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSwipeGesture17verticalDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal swipeAngle()
func (this *QSwipeGesture) SwipeAngle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSwipeGesture10swipeAngleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgesture.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSwipeAngle(qreal)
func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSwipeGesture13setSwipeAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

type QSwipeGesture__SwipeDirection = int

const QSwipeGesture__NoDirection QSwipeGesture__SwipeDirection = 0
const QSwipeGesture__Left QSwipeGesture__SwipeDirection = 1
const QSwipeGesture__Right QSwipeGesture__SwipeDirection = 2
const QSwipeGesture__Up QSwipeGesture__SwipeDirection = 3
const QSwipeGesture__Down QSwipeGesture__SwipeDirection = 4

//  body block end
