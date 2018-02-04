package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
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
	*QGesture
}

func (this *QTapAndHoldGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func (this *QTapAndHoldGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQTapAndHoldGestureFromPointer(cthis unsafe.Pointer) *QTapAndHoldGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QTapAndHoldGesture{bcthis0}
}
func (*QTapAndHoldGesture) NewFromPointer(cthis unsafe.Pointer) *QTapAndHoldGesture {
	return NewQTapAndHoldGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:254
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTapAndHoldGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTapAndHoldGesture(QObject *)
func NewQTapAndHoldGesture(parent *qtcore.QObject /*777 QObject **/) *QTapAndHoldGesture {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTapAndHoldGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:261
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTapAndHoldGesture()
func DeleteQTapAndHoldGesture(this *QTapAndHoldGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:263
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position()
func (this *QTapAndHoldGesture) Position() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)
func (this *QTapAndHoldGesture) SetPosition(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:266
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTimeout(int)
func (this *QTapAndHoldGesture) SetTimeout(msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture10setTimeoutEi", ffiqt.FFI_TYPE_POINTER, msecs)
	gopp.ErrPrint(err, rv)
}
func QTapAndHoldGesture_SetTimeout(msecs int) {
	var nilthis *QTapAndHoldGesture
	nilthis.SetTimeout(msecs)
}

// /usr/include/qt/QtWidgets/qgesture.h:267
// index:0
// Public static Visibility=Default Availability=Available
// [4] int timeout()
func (this *QTapAndHoldGesture) Timeout() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture7timeoutEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QTapAndHoldGesture_Timeout() int {
	var nilthis *QTapAndHoldGesture
	rv := nilthis.Timeout()
	return rv
}

//  body block end
