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
type QStyleOptionTabWidgetFrame struct {
	*QStyleOption
}

func (this *QStyleOptionTabWidgetFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func NewQStyleOptionTabWidgetFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabWidgetFrame{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:171
// index:0
// Public
// void QStyleOptionTabWidgetFrame()
func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabWidgetFrameFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:176
// index:1
// Protected
// void QStyleOptionTabWidgetFrame(int)
func NewQStyleOptionTabWidgetFrame_1(version int) *QStyleOptionTabWidgetFrame {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabWidgetFrameFromPointer(cthis)
	return gothis
}

//  body block end
