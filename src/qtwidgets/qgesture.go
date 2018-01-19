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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgesture.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGesture) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:73
// index:0
// void QGesture(class QObject *)
func NewQGesture(parent unsafe.Pointer) *QGesture {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGesture{cthis}
}

// /usr/include/qt/QtWidgets/qgesture.h:74
// index:0
// virtual
// void ~QGesture()
func DeleteQGesture(*QGesture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGestureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:76
// index:0
// Qt::GestureType gestureType()
func (this *QGesture) GestureType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture11gestureTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:78
// index:0
// Qt::GestureState state()
func (this *QGesture) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture5stateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:80
// index:0
// QPointF hotSpot()
func (this *QGesture) HotSpot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture7hotSpotEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:81
// index:0
// void setHotSpot(const class QPointF &)
func (this *QGesture) SetHotSpot(value unsafe.Pointer) {
	// 0: (, const QPointF & value), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture10setHotSpotERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:82
// index:0
// bool hasHotSpot()
func (this *QGesture) HasHotSpot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture10hasHotSpotEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:83
// index:0
// void unsetHotSpot()
func (this *QGesture) UnsetHotSpot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture12unsetHotSpotEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:90
// index:0
// void setGestureCancelPolicy(enum QGesture::GestureCancelPolicy)
func (this *QGesture) SetGestureCancelPolicy(policy int) {
	// 0: (, QGesture::GestureCancelPolicy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QGesture22setGestureCancelPolicyENS_19GestureCancelPolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:91
// index:0
// QGesture::GestureCancelPolicy gestureCancelPolicy()
func (this *QGesture) GestureCancelPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QGesture19gestureCancelPolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
