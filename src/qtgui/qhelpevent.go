//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QHelpEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:680
// index:0
// void QHelpEvent(enum QEvent::Type, const class QPoint &, const class QPoint &)
func NewQHelpEvent(type_ int, pos unsafe.Pointer, globalPos unsafe.Pointer) *QHelpEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventC2EN6QEvent4TypeERK6QPointS4_", ffiqt.FFI_TYPE_VOID, cthis, &type_, pos, globalPos)
	gopp.ErrPrint(err, rv)
	return &QHelpEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:681
// index:0
// virtual
// void ~QHelpEvent()
func DeleteQHelpEvent(*QHelpEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:683
// index:0
// inline
// int x()
func (this *QHelpEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:684
// index:0
// inline
// int y()
func (this *QHelpEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:685
// index:0
// inline
// int globalX()
func (this *QHelpEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:686
// index:0
// inline
// int globalY()
func (this *QHelpEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:688
// index:0
// inline
// const QPoint & pos()
func (this *QHelpEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:689
// index:0
// inline
// const QPoint & globalPos()
func (this *QHelpEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
