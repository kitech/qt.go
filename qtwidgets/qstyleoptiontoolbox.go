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
type QStyleOptionToolBox struct {
	*QStyleOption
}
type QStyleOptionToolBox_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBox_PTR() *QStyleOptionToolBox
}

func (ptr *QStyleOptionToolBox) QStyleOptionToolBox_PTR() *QStyleOptionToolBox { return ptr }

func (this *QStyleOptionToolBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionToolBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionToolBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBox {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionToolBox{bcthis0}
}
func (*QStyleOptionToolBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBox {
	return NewQStyleOptionToolBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:461
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBox()

/*

 */
func (*QStyleOptionToolBox) NewForInherit() *QStyleOptionToolBox {
	return NewQStyleOptionToolBox()
}
func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:465
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBox(int)

/*

 */
func (*QStyleOptionToolBox) NewForInherit1(version int) *QStyleOptionToolBox {
	return NewQStyleOptionToolBox1(version)
}
func NewQStyleOptionToolBox1(version int) *QStyleOptionToolBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBox)
	return gothis
}

func DeleteQStyleOptionToolBox(this *QStyleOptionToolBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionToolBox__StyleOptionType = int

//
const QStyleOptionToolBox__Type QStyleOptionToolBox__StyleOptionType = 7

func (this *QStyleOptionToolBox) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionToolBox__Type: // 7
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBox_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionToolBox
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionToolBox__StyleOptionVersion = int

// 1
const QStyleOptionToolBox__Version QStyleOptionToolBox__StyleOptionVersion = 2

func (this *QStyleOptionToolBox) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionToolBox__Version: // 2
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBox_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionToolBox
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionToolBox__TabPosition = int

//
const QStyleOptionToolBox__Beginning QStyleOptionToolBox__TabPosition = 0

//
const QStyleOptionToolBox__Middle QStyleOptionToolBox__TabPosition = 1

//
const QStyleOptionToolBox__End QStyleOptionToolBox__TabPosition = 2

//
const QStyleOptionToolBox__OnlyOneTab QStyleOptionToolBox__TabPosition = 3

func (this *QStyleOptionToolBox) TabPositionItemName(val int) string {
	switch val {
	case QStyleOptionToolBox__Beginning: // 0
		return "Beginning"
	case QStyleOptionToolBox__Middle: // 1
		return "Middle"
	case QStyleOptionToolBox__End: // 2
		return "End"
	case QStyleOptionToolBox__OnlyOneTab: // 3
		return "OnlyOneTab"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBox_TabPositionItemName(val int) string {
	var nilthis *QStyleOptionToolBox
	return nilthis.TabPositionItemName(val)
}

/*


 */
type QStyleOptionToolBox__SelectedPosition = int

//
const QStyleOptionToolBox__NotAdjacent QStyleOptionToolBox__SelectedPosition = 0

//
const QStyleOptionToolBox__NextIsSelected QStyleOptionToolBox__SelectedPosition = 1

//
const QStyleOptionToolBox__PreviousIsSelected QStyleOptionToolBox__SelectedPosition = 2

func (this *QStyleOptionToolBox) SelectedPositionItemName(val int) string {
	switch val {
	case QStyleOptionToolBox__NotAdjacent: // 0
		return "NotAdjacent"
	case QStyleOptionToolBox__NextIsSelected: // 1
		return "NextIsSelected"
	case QStyleOptionToolBox__PreviousIsSelected: // 2
		return "PreviousIsSelected"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBox_SelectedPositionItemName(val int) string {
	var nilthis *QStyleOptionToolBox
	return nilthis.SelectedPositionItemName(val)
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
