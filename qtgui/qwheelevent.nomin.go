//  header block begin

// +build !minimal

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
// extern C begin: 0
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

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] QPoint pixelDelta() const

/*
 */
func (this *QWheelEvent) PixelDelta() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.Qtcc3(1776295009, "_ZNK11QWheelEvent10pixelDeltaEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QPointFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:214
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] QPoint angleDelta() const

/*
 */
func (this *QWheelEvent) AngleDelta() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.Qtcc3(2989481828, "_ZNK11QWheelEvent10angleDeltaEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QPointFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

func DeleteQWheelEvent(this *QWheelEvent) {
	rv, err := qtrt.Qtcc3(2730191963, "_ZN11QWheelEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10076() {
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
