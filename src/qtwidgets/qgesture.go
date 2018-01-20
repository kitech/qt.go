//  header block begin
// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
type QGesture struct {
	*qtcore.QObject
}

func (this *QGesture) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQGestureFromPointer(cthis unsafe.Pointer) *QGesture {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGesture{bcthis0}
}

// /usr/include/qt/QtWidgets/qgesture.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGesture) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:73
// index:0
// Public
// void QGesture(class QObject *)
func NewQGesture(parent unsafe.Pointer) *QGesture {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGestureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:74
// index:0
// Public virtual
// void ~QGesture()
func DeleteQGesture(*QGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:76
// index:0
// Public
// Qt::GestureType gestureType()
func (this *QGesture) GestureType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture11gestureTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:78
// index:0
// Public
// Qt::GestureState state()
func (this *QGesture) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:80
// index:0
// Public
// QPointF hotSpot()
func (this *QGesture) HotSpot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture7hotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:81
// index:0
// Public
// void setHotSpot(const class QPointF &)
func (this *QGesture) SetHotSpot(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture10setHotSpotERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:82
// index:0
// Public
// bool hasHotSpot()
func (this *QGesture) HasHotSpot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10hasHotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgesture.h:83
// index:0
// Public
// void unsetHotSpot()
func (this *QGesture) UnsetHotSpot() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture12unsetHotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:90
// index:0
// Public
// void setGestureCancelPolicy(enum QGesture::GestureCancelPolicy)
func (this *QGesture) SetGestureCancelPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture22setGestureCancelPolicyENS_19GestureCancelPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:91
// index:0
// Public
// QGesture::GestureCancelPolicy gestureCancelPolicy()
func (this *QGesture) GestureCancelPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture19gestureCancelPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
