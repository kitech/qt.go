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
// Public virtual
// const QMetaObject * metaObject()
func (this *QTapAndHoldGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:260
// index:0
// Public
// void QTapAndHoldGesture(class QObject *)
func NewQTapAndHoldGesture(parent *qtcore.QObject /*777 QObject **/) *QTapAndHoldGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTapAndHoldGestureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:261
// index:0
// Public virtual
// void ~QTapAndHoldGesture()
func DeleteQTapAndHoldGesture(*QTapAndHoldGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:263
// index:0
// Public
// QPointF position()
func (this *QTapAndHoldGesture) Position() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture8positionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:264
// index:0
// Public
// void setPosition(const class QPointF &)
func (this *QTapAndHoldGesture) SetPosition(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:266
// index:0
// Public static
// void setTimeout(int)
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
// Public static
// int timeout()
func (this *QTapAndHoldGesture) Timeout() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTapAndHoldGesture7timeoutEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QTapAndHoldGesture_Timeout() int {
	var nilthis *QTapAndHoldGesture
	rv := nilthis.Timeout()
	return rv
}

//  body block end
