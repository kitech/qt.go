//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QResizeEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:463
// index:0
// void QResizeEvent(const class QSize &, const class QSize &)
func NewQResizeEvent(size unsafe.Pointer, oldSize unsafe.Pointer) *QResizeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QResizeEventC2ERK5QSizeS2_", ffiqt.FFI_TYPE_VOID, cthis, size, oldSize)
	gopp.ErrPrint(err, rv)
	return &QResizeEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:464
// index:0
// virtual
// void ~QResizeEvent()
func DeleteQResizeEvent(*QResizeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QResizeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:466
// index:0
// inline
// const QSize & size()
func (this *QResizeEvent) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QResizeEvent4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:467
// index:0
// inline
// const QSize & oldSize()
func (this *QResizeEvent) OldSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QResizeEvent7oldSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
