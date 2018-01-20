//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QStyleOptionSlider struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSlider) GetCthis() unsafe.Pointer {
	return this.QStyleOptionComplex.GetCthis()
}
func NewQStyleOptionSliderFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSlider{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:533
// index:0
// Public
// void QStyleOptionSlider()
func NewQStyleOptionSlider() *QStyleOptionSlider {
	cthis := qtrt.Calloc(1, 256) // 128
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:537
// index:1
// Protected
// void QStyleOptionSlider(int)
func NewQStyleOptionSlider_1(version int) *QStyleOptionSlider {
	cthis := qtrt.Calloc(1, 256) // 128
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(cthis)
	return gothis
}

//  body block end
