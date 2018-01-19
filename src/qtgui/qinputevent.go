//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QInputEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:72
// index:0
// virtual
// void ~QInputEvent()
func DeleteQInputEvent(*QInputEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:73
// index:0
// inline
// Qt::KeyboardModifiers modifiers()
func (this *QInputEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:75
// index:0
// inline
// ulong timestamp()
func (this *QInputEvent) Timestamp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9timestampEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:76
// index:0
// inline
// void setTimestamp(ulong)
func (this *QInputEvent) SetTimestamp(atimestamp uint) {
	// 0: (, ulong atimestamp), (&atimestamp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEvent12setTimestampEm", ffiqt.FFI_TYPE_VOID, this.cthis, &atimestamp)
	gopp.ErrPrint(err, rv)
}

//  body block end
