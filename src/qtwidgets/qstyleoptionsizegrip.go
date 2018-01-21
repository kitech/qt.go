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
type QStyleOptionSizeGrip struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSizeGrip) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func NewQStyleOptionSizeGripFromPointer(cthis unsafe.Pointer) *QStyleOptionSizeGrip {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSizeGrip{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:653
// index:0
// Public
// void QStyleOptionSizeGrip()
func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:656
// index:1
// Protected
// void QStyleOptionSizeGrip(int)
func NewQStyleOptionSizeGrip_1(version int) *QStyleOptionSizeGrip {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(cthis)
	return gothis
}

//  body block end
