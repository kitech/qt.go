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
type QStyleOptionTab struct {
	*QStyleOption
}
type QStyleOptionTab_ITF interface {
	QStyleOption_ITF
	QStyleOptionTab_PTR() *QStyleOptionTab
}

func (ptr *QStyleOptionTab) QStyleOptionTab_PTR() *QStyleOptionTab { return ptr }

func (this *QStyleOptionTab) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTab) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabFromPointer(cthis unsafe.Pointer) *QStyleOptionTab {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTab{bcthis0}
}
func (*QStyleOptionTab) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTab {
	return NewQStyleOptionTabFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTab()

/*

 */
func NewQStyleOptionTab() *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:289
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTab(int)

/*

 */
func NewQStyleOptionTab_1(version int) *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

func DeleteQStyleOptionTab(this *QStyleOptionTab) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionTab__StyleOptionType = int

//
const QStyleOptionTab__Type QStyleOptionTab__StyleOptionType = 3

func (this *QStyleOptionTab) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionTab__Type: // 3
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionTab__StyleOptionVersion = int

// 1
const QStyleOptionTab__Version QStyleOptionTab__StyleOptionVersion = 3

func (this *QStyleOptionTab) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionTab__Version: // 3
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionTab__TabPosition = int

//
const QStyleOptionTab__Beginning QStyleOptionTab__TabPosition = 0

//
const QStyleOptionTab__Middle QStyleOptionTab__TabPosition = 1

//
const QStyleOptionTab__End QStyleOptionTab__TabPosition = 2

//
const QStyleOptionTab__OnlyOneTab QStyleOptionTab__TabPosition = 3

func (this *QStyleOptionTab) TabPositionItemName(val int) string {
	switch val {
	case QStyleOptionTab__Beginning: // 0
		return "Beginning"
	case QStyleOptionTab__Middle: // 1
		return "Middle"
	case QStyleOptionTab__End: // 2
		return "End"
	case QStyleOptionTab__OnlyOneTab: // 3
		return "OnlyOneTab"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_TabPositionItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.TabPositionItemName(val)
}

/*


 */
type QStyleOptionTab__SelectedPosition = int

//
const QStyleOptionTab__NotAdjacent QStyleOptionTab__SelectedPosition = 0

//
const QStyleOptionTab__NextIsSelected QStyleOptionTab__SelectedPosition = 1

//
const QStyleOptionTab__PreviousIsSelected QStyleOptionTab__SelectedPosition = 2

func (this *QStyleOptionTab) SelectedPositionItemName(val int) string {
	switch val {
	case QStyleOptionTab__NotAdjacent: // 0
		return "NotAdjacent"
	case QStyleOptionTab__NextIsSelected: // 1
		return "NextIsSelected"
	case QStyleOptionTab__PreviousIsSelected: // 2
		return "PreviousIsSelected"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_SelectedPositionItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.SelectedPositionItemName(val)
}

/*


 */
type QStyleOptionTab__CornerWidget = int

//
const QStyleOptionTab__NoCornerWidgets QStyleOptionTab__CornerWidget = 0

//
const QStyleOptionTab__LeftCornerWidget QStyleOptionTab__CornerWidget = 1

//
const QStyleOptionTab__RightCornerWidget QStyleOptionTab__CornerWidget = 2

func (this *QStyleOptionTab) CornerWidgetItemName(val int) string {
	switch val {
	case QStyleOptionTab__NoCornerWidgets: // 0
		return "NoCornerWidgets"
	case QStyleOptionTab__LeftCornerWidget: // 1
		return "LeftCornerWidget"
	case QStyleOptionTab__RightCornerWidget: // 2
		return "RightCornerWidget"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_CornerWidgetItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.CornerWidgetItemName(val)
}

/*


 */
type QStyleOptionTab__TabFeature = int

//
const QStyleOptionTab__None QStyleOptionTab__TabFeature = 0

//
const QStyleOptionTab__HasFrame QStyleOptionTab__TabFeature = 1

func (this *QStyleOptionTab) TabFeatureItemName(val int) string {
	switch val {
	case QStyleOptionTab__None: // 0
		return "None"
	case QStyleOptionTab__HasFrame: // 1
		return "HasFrame"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTab_TabFeatureItemName(val int) string {
	var nilthis *QStyleOptionTab
	return nilthis.TabFeatureItemName(val)
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
