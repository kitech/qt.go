package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionRubberBand struct {
	*QStyleOption
}
type QStyleOptionRubberBand_ITF interface {
	QStyleOption_ITF
	QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand
}

func (ptr *QStyleOptionRubberBand) QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand { return ptr }

func (this *QStyleOptionRubberBand) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionRubberBand) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionRubberBandFromPointer(cthis unsafe.Pointer) *QStyleOptionRubberBand {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionRubberBand{bcthis0}
}
func (*QStyleOptionRubberBand) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBandFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:491
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand()
func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:495
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand(int)
func NewQStyleOptionRubberBand_1(version int) *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

func DeleteQStyleOptionRubberBand(this *QStyleOptionRubberBand) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionRubberBand__StyleOptionType = int

const QStyleOptionRubberBand__Type QStyleOptionRubberBand__StyleOptionType = 13

type QStyleOptionRubberBand__StyleOptionVersion = int

const QStyleOptionRubberBand__Version QStyleOptionRubberBand__StyleOptionVersion = 1

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
