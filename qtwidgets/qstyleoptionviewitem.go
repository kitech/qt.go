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
type QStyleOptionViewItem struct {
	*QStyleOption
}
type QStyleOptionViewItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionViewItem_PTR() *QStyleOptionViewItem
}

func (ptr *QStyleOptionViewItem) QStyleOptionViewItem_PTR() *QStyleOptionViewItem { return ptr }

func (this *QStyleOptionViewItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionViewItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionViewItemFromPointer(cthis unsafe.Pointer) *QStyleOptionViewItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionViewItem{bcthis0}
}
func (*QStyleOptionViewItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionViewItem {
	return NewQStyleOptionViewItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:442
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem()

/*

 */
func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:446
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem(int)

/*

 */
func NewQStyleOptionViewItem_1(version int) *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

func DeleteQStyleOptionViewItem(this *QStyleOptionViewItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionViewItem__StyleOptionType = int

//
const QStyleOptionViewItem__Type QStyleOptionViewItem__StyleOptionType = 10

func (this *QStyleOptionViewItem) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionViewItem__Type: // 10
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionViewItem_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionViewItem
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionViewItem__StyleOptionVersion = int

// 1
const QStyleOptionViewItem__Version QStyleOptionViewItem__StyleOptionVersion = 4

func (this *QStyleOptionViewItem) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionViewItem__Version: // 4
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionViewItem_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionViewItem
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionViewItem__Position = int

//
const QStyleOptionViewItem__Left QStyleOptionViewItem__Position = 0

//
const QStyleOptionViewItem__Right QStyleOptionViewItem__Position = 1

//
const QStyleOptionViewItem__Top QStyleOptionViewItem__Position = 2

//
const QStyleOptionViewItem__Bottom QStyleOptionViewItem__Position = 3

func (this *QStyleOptionViewItem) PositionItemName(val int) string {
	switch val {
	case QStyleOptionViewItem__Left: // 0
		return "Left"
	case QStyleOptionViewItem__Right: // 1
		return "Right"
	case QStyleOptionViewItem__Top: // 2
		return "Top"
	case QStyleOptionViewItem__Bottom: // 3
		return "Bottom"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionViewItem_PositionItemName(val int) string {
	var nilthis *QStyleOptionViewItem
	return nilthis.PositionItemName(val)
}

/*


 */
type QStyleOptionViewItem__ViewItemFeature = int

//
const QStyleOptionViewItem__None QStyleOptionViewItem__ViewItemFeature = 0

//
const QStyleOptionViewItem__WrapText QStyleOptionViewItem__ViewItemFeature = 1

//
const QStyleOptionViewItem__Alternate QStyleOptionViewItem__ViewItemFeature = 2

//
const QStyleOptionViewItem__HasCheckIndicator QStyleOptionViewItem__ViewItemFeature = 4

//
const QStyleOptionViewItem__HasDisplay QStyleOptionViewItem__ViewItemFeature = 8

//
const QStyleOptionViewItem__HasDecoration QStyleOptionViewItem__ViewItemFeature = 16

func (this *QStyleOptionViewItem) ViewItemFeatureItemName(val int) string {
	switch val {
	case QStyleOptionViewItem__None: // 0
		return "None"
	case QStyleOptionViewItem__WrapText: // 1
		return "WrapText"
	case QStyleOptionViewItem__Alternate: // 2
		return "Alternate"
	case QStyleOptionViewItem__HasCheckIndicator: // 4
		return "HasCheckIndicator"
	case QStyleOptionViewItem__HasDisplay: // 8
		return "HasDisplay"
	case QStyleOptionViewItem__HasDecoration: // 16
		return "HasDecoration"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionViewItem_ViewItemFeatureItemName(val int) string {
	var nilthis *QStyleOptionViewItem
	return nilthis.ViewItemFeatureItemName(val)
}

/*


 */
type QStyleOptionViewItem__ViewItemPosition = int

//
const QStyleOptionViewItem__Invalid QStyleOptionViewItem__ViewItemPosition = 0

//
const QStyleOptionViewItem__Beginning QStyleOptionViewItem__ViewItemPosition = 1

//
const QStyleOptionViewItem__Middle QStyleOptionViewItem__ViewItemPosition = 2

//
const QStyleOptionViewItem__End QStyleOptionViewItem__ViewItemPosition = 3

//
const QStyleOptionViewItem__OnlyOne QStyleOptionViewItem__ViewItemPosition = 4

func (this *QStyleOptionViewItem) ViewItemPositionItemName(val int) string {
	switch val {
	case QStyleOptionViewItem__Invalid: // 0
		return "Invalid"
	case QStyleOptionViewItem__Beginning: // 1
		return "Beginning"
	case QStyleOptionViewItem__Middle: // 2
		return "Middle"
	case QStyleOptionViewItem__End: // 3
		return "End"
	case QStyleOptionViewItem__OnlyOne: // 4
		return "OnlyOne"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionViewItem_ViewItemPositionItemName(val int) string {
	var nilthis *QStyleOptionViewItem
	return nilthis.ViewItemPositionItemName(val)
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
