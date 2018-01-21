package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleOptionFocusRect struct {
	*QStyleOption
}

func (this *QStyleOptionFocusRect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func NewQStyleOptionFocusRectFromPointer(cthis unsafe.Pointer) *QStyleOptionFocusRect {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFocusRect{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:119
// index:0
// Public
// void QStyleOptionFocusRect()
func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:123
// index:1
// Protected
// void QStyleOptionFocusRect(int)
func NewQStyleOptionFocusRect_1(version int) *QStyleOptionFocusRect {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(cthis)
	return gothis
}

//  body block end
