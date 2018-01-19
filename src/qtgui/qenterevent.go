//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QEnterEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:85
// index:0
// void QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
func NewQEnterEvent(localPos unsafe.Pointer, windowPos unsafe.Pointer, screenPos unsafe.Pointer) *QEnterEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QEnterEventC2ERK7QPointFS2_S2_", ffiqt.FFI_TYPE_VOID, cthis, localPos, windowPos, screenPos)
	gopp.ErrPrint(err, rv)
	return &QEnterEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:86
// index:0
// virtual
// void ~QEnterEvent()
func DeleteQEnterEvent(*QEnterEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QEnterEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:89
// index:0
// inline
// QPoint pos()
func (this *QEnterEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:90
// index:0
// inline
// QPoint globalPos()
func (this *QEnterEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:91
// index:0
// inline
// int x()
func (this *QEnterEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:92
// index:0
// inline
// int y()
func (this *QEnterEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:93
// index:0
// inline
// int globalX()
func (this *QEnterEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:94
// index:0
// inline
// int globalY()
func (this *QEnterEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:96
// index:0
// inline
// const QPointF & localPos()
func (this *QEnterEvent) LocalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent8localPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:97
// index:0
// inline
// const QPointF & windowPos()
func (this *QEnterEvent) WindowPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9windowPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:98
// index:0
// inline
// const QPointF & screenPos()
func (this *QEnterEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
