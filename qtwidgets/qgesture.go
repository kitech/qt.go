package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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

type QGesture struct {
	*qtcore.QObject
}

func (this *QGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGesture) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGestureFromPointer(cthis unsafe.Pointer) *QGesture {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGesture{bcthis0}
}
func (*QGesture) NewFromPointer(cthis unsafe.Pointer) *QGesture {
	return NewQGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGesture(QObject *)
func NewQGesture(parent *qtcore.QObject /*777 QObject **/) *QGesture {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGesture()
func DeleteQGesture(this *QGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::GestureType gestureType()
func (this *QGesture) GestureType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture11gestureTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::GestureState state()
func (this *QGesture) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF hotSpot()
func (this *QGesture) HotSpot() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture7hotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHotSpot(const QPointF &)
func (this *QGesture) SetHotSpot(value *qtcore.QPointF) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture10setHotSpotERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasHotSpot()
func (this *QGesture) HasHotSpot() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10hasHotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetHotSpot()
func (this *QGesture) UnsetHotSpot() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture12unsetHotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGestureCancelPolicy(enum QGesture::GestureCancelPolicy)
func (this *QGesture) SetGestureCancelPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture22setGestureCancelPolicyENS_19GestureCancelPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] QGesture::GestureCancelPolicy gestureCancelPolicy()
func (this *QGesture) GestureCancelPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture19gestureCancelPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QGesture__GestureCancelPolicy = int

const QGesture__CancelNone QGesture__GestureCancelPolicy = 0
const QGesture__CancelAllInContext QGesture__GestureCancelPolicy = 1

//  body block end
