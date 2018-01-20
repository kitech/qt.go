//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QStatusTipEvent struct {
	*qtcore.QEvent
}

func (this *QStatusTipEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:700
// index:0
// void QStatusTipEvent(const class QString &)
func NewQStatusTipEvent(tip unsafe.Pointer) *QStatusTipEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStatusTipEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, tip)
	gopp.ErrPrint(err, rv)
	gothis := NewQStatusTipEventFromPointer(cthis)
	return gothis
}
func NewQStatusTipEventFromPointer(cthis unsafe.Pointer) *QStatusTipEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QStatusTipEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:701
// index:0
// virtual
// void ~QStatusTipEvent()
func DeleteQStatusTipEvent(*QStatusTipEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStatusTipEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:703
// index:0
// inline
// QString tip()
func (this *QStatusTipEvent) Tip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QStatusTipEvent3tipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
