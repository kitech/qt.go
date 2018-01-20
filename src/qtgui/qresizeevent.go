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
	*qtcore.QEvent
}

func (this *QResizeEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQResizeEventFromPointer(cthis unsafe.Pointer) *QResizeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QResizeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:463
// index:0
// Public
// void QResizeEvent(const class QSize &, const class QSize &)
func NewQResizeEvent(size *qtcore.QSize, oldSize *qtcore.QSize) *QResizeEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = size.GetCthis()
	var convArg1 = oldSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QResizeEventC2ERK5QSizeS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQResizeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:464
// index:0
// Public virtual
// void ~QResizeEvent()
func DeleteQResizeEvent(*QResizeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QResizeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:466
// index:0
// Public inline
// const QSize & size()
func (this *QResizeEvent) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QResizeEvent4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:467
// index:0
// Public inline
// const QSize & oldSize()
func (this *QResizeEvent) OldSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QResizeEvent7oldSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
