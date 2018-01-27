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
import "gopp"
import "qt.go/cffiqt"
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
type QStyleOptionSpinBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSpinBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSpinBox{bcthis0}
}
func (*QStyleOptionSpinBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	return NewQStyleOptionSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:552
// index:0
// Public
// void QStyleOptionSpinBox()
func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:556
// index:1
// Protected
// void QStyleOptionSpinBox(int)
func NewQStyleOptionSpinBox_1(version int) *QStyleOptionSpinBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(cthis)
	return gothis
}

type QStyleOptionSpinBox__StyleOptionType = int

const QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = 983042

type QStyleOptionSpinBox__StyleOptionVersion = int

const QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = 1

//  body block end
