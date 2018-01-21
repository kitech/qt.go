package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QStyleOptionToolButton struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionToolButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func NewQStyleOptionToolButtonFromPointer(cthis unsafe.Pointer) *QStyleOptionToolButton {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionToolButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:579
// index:0
// Public
// void QStyleOptionToolButton()
func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:583
// index:1
// Protected
// void QStyleOptionToolButton(int)
func NewQStyleOptionToolButton_1(version int) *QStyleOptionToolButton {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(cthis)
	return gothis
}

//  body block end
