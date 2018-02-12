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

type QStyleOptionHeader struct {
	*QStyleOption
}

func (this *QStyleOptionHeader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionHeader) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionHeaderFromPointer(cthis unsafe.Pointer) *QStyleOptionHeader {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionHeader{bcthis0}
}
func (*QStyleOptionHeader) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionHeader {
	return NewQStyleOptionHeaderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionHeader()
func NewQStyleOptionHeader() *QStyleOptionHeader {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionHeaderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionHeaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionHeader)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:230
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionHeader(int)
func NewQStyleOptionHeader_1(version int) *QStyleOptionHeader {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionHeaderC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionHeaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionHeader)
	return gothis
}

func DeleteQStyleOptionHeader(this *QStyleOptionHeader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionHeaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionHeader__StyleOptionType = int

const QStyleOptionHeader__Type QStyleOptionHeader__StyleOptionType = 8

type QStyleOptionHeader__StyleOptionVersion = int

const QStyleOptionHeader__Version QStyleOptionHeader__StyleOptionVersion = 1

type QStyleOptionHeader__SectionPosition = int

const QStyleOptionHeader__Beginning QStyleOptionHeader__SectionPosition = 0
const QStyleOptionHeader__Middle QStyleOptionHeader__SectionPosition = 1
const QStyleOptionHeader__End QStyleOptionHeader__SectionPosition = 2
const QStyleOptionHeader__OnlyOneSection QStyleOptionHeader__SectionPosition = 3

type QStyleOptionHeader__SelectedPosition = int

const QStyleOptionHeader__NotAdjacent QStyleOptionHeader__SelectedPosition = 0
const QStyleOptionHeader__NextIsSelected QStyleOptionHeader__SelectedPosition = 1
const QStyleOptionHeader__PreviousIsSelected QStyleOptionHeader__SelectedPosition = 2
const QStyleOptionHeader__NextAndPreviousAreSelected QStyleOptionHeader__SelectedPosition = 3

type QStyleOptionHeader__SortIndicator = int

const QStyleOptionHeader__None QStyleOptionHeader__SortIndicator = 0
const QStyleOptionHeader__SortUp QStyleOptionHeader__SortIndicator = 1
const QStyleOptionHeader__SortDown QStyleOptionHeader__SortIndicator = 2

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
