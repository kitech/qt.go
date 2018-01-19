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
type QWindowStateChangeEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:784
// index:0
// virtual
// void ~QWindowStateChangeEvent()
func DeleteQWindowStateChangeEvent(*QWindowStateChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QWindowStateChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:786
// index:0
// inline
// Qt::WindowStates oldState()
func (this *QWindowStateChangeEvent) OldState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QWindowStateChangeEvent8oldStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:787
// index:0
// bool isOverride()
func (this *QWindowStateChangeEvent) IsOverride() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QWindowStateChangeEvent10isOverrideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
