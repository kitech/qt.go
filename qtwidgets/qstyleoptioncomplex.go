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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QStyleOptionComplex struct {
	*QStyleOption
}

func (this *QStyleOptionComplex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionComplex) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionComplexFromPointer(cthis unsafe.Pointer) *QStyleOptionComplex {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionComplex{bcthis0}
}
func (*QStyleOptionComplex) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionComplex {
	return NewQStyleOptionComplexFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:509
// index:0
// Public
// void QStyleOptionComplex(int, int)
func NewQStyleOptionComplex(version int, type_ int) *QStyleOptionComplex {
	cthis := qtrt.Calloc(1, 256) // 72
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionComplexC2Eii", ffiqt.FFI_TYPE_VOID, cthis, version, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionComplexFromPointer(cthis)
	return gothis
}

type QStyleOptionComplex__StyleOptionType = int

const QStyleOptionComplex__Type QStyleOptionComplex__StyleOptionType = 983040

type QStyleOptionComplex__StyleOptionVersion = int

const QStyleOptionComplex__Version QStyleOptionComplex__StyleOptionVersion = 1

//  body block end
