//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QWheelEvent struct {
	*QInputEvent
}

func (this *QWheelEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}
func NewQWheelEventFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QWheelEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:196
// index:0
// Public virtual
// void ~QWheelEvent()
func DeleteQWheelEvent(*QWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:199
// index:0
// Public inline
// QPoint pixelDelta()
func (this *QWheelEvent) PixelDelta() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10pixelDeltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:200
// index:0
// Public inline
// QPoint angleDelta()
func (this *QWheelEvent) AngleDelta() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10angleDeltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:202
// index:0
// Public inline
// int delta()
func (this *QWheelEvent) Delta() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5deltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:203
// index:0
// Public inline
// Qt::Orientation orientation()
func (this *QWheelEvent) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:206
// index:0
// Public inline
// QPoint pos()
func (this *QWheelEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:207
// index:0
// Public inline
// QPoint globalPos()
func (this *QWheelEvent) GlobalPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:208
// index:0
// Public inline
// int x()
func (this *QWheelEvent) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:209
// index:0
// Public inline
// int y()
func (this *QWheelEvent) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:210
// index:0
// Public inline
// int globalX()
func (this *QWheelEvent) GlobalX() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:211
// index:0
// Public inline
// int globalY()
func (this *QWheelEvent) GlobalY() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// Public inline
// const QPointF & posF()
func (this *QWheelEvent) PosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:214
// index:0
// Public inline
// const QPointF & globalPosF()
func (this *QWheelEvent) GlobalPosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10globalPosFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:216
// index:0
// Public inline
// Qt::MouseButtons buttons()
func (this *QWheelEvent) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// Public inline
// Qt::ScrollPhase phase()
func (this *QWheelEvent) Phase() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:219
// index:0
// Public inline
// bool inverted()
func (this *QWheelEvent) Inverted() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:221
// index:0
// Public inline
// Qt::MouseEventSource source()
func (this *QWheelEvent) Source() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
