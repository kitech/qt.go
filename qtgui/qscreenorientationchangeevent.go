package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QScreenOrientationChangeEvent struct {
	*qtcore.QEvent
}
type QScreenOrientationChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QScreenOrientationChangeEvent_PTR() *QScreenOrientationChangeEvent
}

func (ptr *QScreenOrientationChangeEvent) QScreenOrientationChangeEvent_PTR() *QScreenOrientationChangeEvent {
	return ptr
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
func NewQScreenOrientationChangeEvent(screen QScreen_ITF /*777 QScreen **/, orientation int) *QScreenOrientationChangeEvent {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventC2EP7QScreenN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, convArg0, orientation)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScreenOrientationChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScreenOrientationChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1039
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScreenOrientationChangeEvent()
func DeleteQScreenOrientationChangeEvent(this *QScreenOrientationChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QScreenOrientationChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:1041
// index:0
// Public Visibility=Default Availability=Available
// [8] QScreen * screen()
func (this *QScreenOrientationChangeEvent) Screen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent6screenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:1042
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation orientation()
func (this *QScreenOrientationChangeEvent) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QScreenOrientationChangeEvent11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
