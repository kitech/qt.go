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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QStyleOptionHeader struct {
	*QStyleOption
}
type QStyleOptionHeader_ITF interface {
	QStyleOption_ITF
	QStyleOptionHeader_PTR() *QStyleOptionHeader
}

func (ptr *QStyleOptionHeader) QStyleOptionHeader_PTR() *QStyleOptionHeader { return ptr }

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

// /usr/include/qt/QtWidgets/qstyleoption.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionHeader()

/*

 */
func (*QStyleOptionHeader) NewForInherit() *QStyleOptionHeader {
	return NewQStyleOptionHeader()
}
func NewQStyleOptionHeader() *QStyleOptionHeader {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionHeaderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionHeaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionHeader)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:219
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionHeader(int)

/*

 */
func (*QStyleOptionHeader) NewForInherit1(version int) *QStyleOptionHeader {
	return NewQStyleOptionHeader1(version)
}
func NewQStyleOptionHeader1(version int) *QStyleOptionHeader {
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

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionHeader__StyleOptionType = int

//
const QStyleOptionHeader__Type QStyleOptionHeader__StyleOptionType = 8

func (this *QStyleOptionHeader) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionHeader__Type: // 8
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionHeader_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionHeader
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionHeader__StyleOptionVersion = int

// 1
const QStyleOptionHeader__Version QStyleOptionHeader__StyleOptionVersion = 1

func (this *QStyleOptionHeader) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionHeader__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionHeader_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionHeader
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionHeader__SectionPosition = int

//
const QStyleOptionHeader__Beginning QStyleOptionHeader__SectionPosition = 0

//
const QStyleOptionHeader__Middle QStyleOptionHeader__SectionPosition = 1

//
const QStyleOptionHeader__End QStyleOptionHeader__SectionPosition = 2

//
const QStyleOptionHeader__OnlyOneSection QStyleOptionHeader__SectionPosition = 3

func (this *QStyleOptionHeader) SectionPositionItemName(val int) string {
	switch val {
	case QStyleOptionHeader__Beginning: // 0
		return "Beginning"
	case QStyleOptionHeader__Middle: // 1
		return "Middle"
	case QStyleOptionHeader__End: // 2
		return "End"
	case QStyleOptionHeader__OnlyOneSection: // 3
		return "OnlyOneSection"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionHeader_SectionPositionItemName(val int) string {
	var nilthis *QStyleOptionHeader
	return nilthis.SectionPositionItemName(val)
}

/*


 */
type QStyleOptionHeader__SelectedPosition = int

//
const QStyleOptionHeader__NotAdjacent QStyleOptionHeader__SelectedPosition = 0

//
const QStyleOptionHeader__NextIsSelected QStyleOptionHeader__SelectedPosition = 1

//
const QStyleOptionHeader__PreviousIsSelected QStyleOptionHeader__SelectedPosition = 2

//
const QStyleOptionHeader__NextAndPreviousAreSelected QStyleOptionHeader__SelectedPosition = 3

func (this *QStyleOptionHeader) SelectedPositionItemName(val int) string {
	switch val {
	case QStyleOptionHeader__NotAdjacent: // 0
		return "NotAdjacent"
	case QStyleOptionHeader__NextIsSelected: // 1
		return "NextIsSelected"
	case QStyleOptionHeader__PreviousIsSelected: // 2
		return "PreviousIsSelected"
	case QStyleOptionHeader__NextAndPreviousAreSelected: // 3
		return "NextAndPreviousAreSelected"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionHeader_SelectedPositionItemName(val int) string {
	var nilthis *QStyleOptionHeader
	return nilthis.SelectedPositionItemName(val)
}

/*


 */
type QStyleOptionHeader__SortIndicator = int

//
const QStyleOptionHeader__None QStyleOptionHeader__SortIndicator = 0

//
const QStyleOptionHeader__SortUp QStyleOptionHeader__SortIndicator = 1

//
const QStyleOptionHeader__SortDown QStyleOptionHeader__SortIndicator = 2

func (this *QStyleOptionHeader) SortIndicatorItemName(val int) string {
	switch val {
	case QStyleOptionHeader__None: // 0
		return "None"
	case QStyleOptionHeader__SortUp: // 1
		return "SortUp"
	case QStyleOptionHeader__SortDown: // 2
		return "SortDown"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionHeader_SortIndicatorItemName(val int) string {
	var nilthis *QStyleOptionHeader
	return nilthis.SortIndicatorItemName(val)
}

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
		log.Println(123)
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
