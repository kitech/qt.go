package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QScreenOrientationChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQScreenOrientationChangeEventFromPointer(cthis unsafe.Pointer) *QScreenOrientationChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QScreenOrientationChangeEvent{bcthis0}
}
func (*QScreenOrientationChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QScreenOrientationChangeEvent {
	return NewQScreenOrientationChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:1038
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScreenOrientationChangeEvent(QScreen *, Qt::ScreenOrientation)
func NewQScreenOrientationChangeEvent(screen *QScreen /*777 QScreen **/, orientation int) *QScreenOrientationChangeEvent {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventC2EP7QScreenN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, convArg0, orientation)
	gopp.ErrPrint(err, rv)
	gothis := NewQScreenOrientationChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScreenOrientationChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1039
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScreenOrientationChangeEvent()
func DeleteQScreenOrientationChangeEvent(this *QScreenOrientationChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:1041
// index:0
// Public Visibility=Default Availability=Available
// [8] QScreen * screen()
func (this *QScreenOrientationChangeEvent) Screen() *QScreen /*777 QScreen **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent6screenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:1042
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation orientation()
func (this *QScreenOrientationChangeEvent) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
