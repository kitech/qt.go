//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

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
type QStyleOptionHeader struct {
	*QStyleOption
}

func (this *QStyleOptionHeader) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}
func NewQStyleOptionHeaderFromPointer(cthis unsafe.Pointer) *QStyleOptionHeader {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionHeader{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:226
// index:0
// Public
// void QStyleOptionHeader()
func NewQStyleOptionHeader() *QStyleOptionHeader {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionHeaderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionHeaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:230
// index:1
// Protected
// void QStyleOptionHeader(int)
func NewQStyleOptionHeader_1(version int) *QStyleOptionHeader {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionHeaderC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionHeaderFromPointer(cthis)
	return gothis
}

//  body block end
