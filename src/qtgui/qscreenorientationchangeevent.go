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
type QScreenOrientationChangeEvent struct {
	*qtcore.QEvent
}

func (this *QScreenOrientationChangeEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:1038
// index:0
// void QScreenOrientationChangeEvent(class QScreen *, Qt::ScreenOrientation)
func NewQScreenOrientationChangeEvent(screen unsafe.Pointer, orientation int) *QScreenOrientationChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventC2EP7QScreenN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, cthis, screen, &orientation)
	gopp.ErrPrint(err, rv)
	gothis := NewQScreenOrientationChangeEventFromPointer(cthis)
	return gothis
}
func NewQScreenOrientationChangeEventFromPointer(cthis unsafe.Pointer) *QScreenOrientationChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QScreenOrientationChangeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:1039
// index:0
// virtual
// void ~QScreenOrientationChangeEvent()
func DeleteQScreenOrientationChangeEvent(*QScreenOrientationChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1041
// index:0
// QScreen * screen()
func (this *QScreenOrientationChangeEvent) Screen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent6screenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:1042
// index:0
// Qt::ScreenOrientation orientation()
func (this *QScreenOrientationChangeEvent) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
