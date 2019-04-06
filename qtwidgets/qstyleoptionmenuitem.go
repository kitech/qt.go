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
type QStyleOptionMenuItem struct {
	*QStyleOption
}
type QStyleOptionMenuItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionMenuItem_PTR() *QStyleOptionMenuItem
}

func (ptr *QStyleOptionMenuItem) QStyleOptionMenuItem_PTR() *QStyleOptionMenuItem { return ptr }

func (this *QStyleOptionMenuItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionMenuItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionMenuItemFromPointer(cthis unsafe.Pointer) *QStyleOptionMenuItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionMenuItem{bcthis0}
}
func (*QStyleOptionMenuItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionMenuItem {
	return NewQStyleOptionMenuItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionMenuItem()

/*

 */
func (*QStyleOptionMenuItem) NewForInherit() *QStyleOptionMenuItem {
	return NewQStyleOptionMenuItem()
}
func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionMenuItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:376
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionMenuItem(int)

/*

 */
func (*QStyleOptionMenuItem) NewForInherit1(version int) *QStyleOptionMenuItem {
	return NewQStyleOptionMenuItem1(version)
}
func NewQStyleOptionMenuItem1(version int) *QStyleOptionMenuItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionMenuItem)
	return gothis
}

func DeleteQStyleOptionMenuItem(this *QStyleOptionMenuItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionMenuItem__StyleOptionType = int

//
const QStyleOptionMenuItem__Type QStyleOptionMenuItem__StyleOptionType = 4

func (this *QStyleOptionMenuItem) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionMenuItem__Type: // 4
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionMenuItem_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionMenuItem
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionMenuItem__StyleOptionVersion = int

// 1
const QStyleOptionMenuItem__Version QStyleOptionMenuItem__StyleOptionVersion = 1

func (this *QStyleOptionMenuItem) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionMenuItem__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionMenuItem_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionMenuItem
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionMenuItem__MenuItemType = int

//
const QStyleOptionMenuItem__Normal QStyleOptionMenuItem__MenuItemType = 0

//
const QStyleOptionMenuItem__DefaultItem QStyleOptionMenuItem__MenuItemType = 1

//
const QStyleOptionMenuItem__Separator QStyleOptionMenuItem__MenuItemType = 2

//
const QStyleOptionMenuItem__SubMenu QStyleOptionMenuItem__MenuItemType = 3

//
const QStyleOptionMenuItem__Scroller QStyleOptionMenuItem__MenuItemType = 4

//
const QStyleOptionMenuItem__TearOff QStyleOptionMenuItem__MenuItemType = 5

//
const QStyleOptionMenuItem__Margin QStyleOptionMenuItem__MenuItemType = 6

//
const QStyleOptionMenuItem__EmptyArea QStyleOptionMenuItem__MenuItemType = 7

func (this *QStyleOptionMenuItem) MenuItemTypeItemName(val int) string {
	switch val {
	case QStyleOptionMenuItem__Normal: // 0
		return "Normal"
	case QStyleOptionMenuItem__DefaultItem: // 1
		return "DefaultItem"
	case QStyleOptionMenuItem__Separator: // 2
		return "Separator"
	case QStyleOptionMenuItem__SubMenu: // 3
		return "SubMenu"
	case QStyleOptionMenuItem__Scroller: // 4
		return "Scroller"
	case QStyleOptionMenuItem__TearOff: // 5
		return "TearOff"
	case QStyleOptionMenuItem__Margin: // 6
		return "Margin"
	case QStyleOptionMenuItem__EmptyArea: // 7
		return "EmptyArea"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionMenuItem_MenuItemTypeItemName(val int) string {
	var nilthis *QStyleOptionMenuItem
	return nilthis.MenuItemTypeItemName(val)
}

/*


 */
type QStyleOptionMenuItem__CheckType = int

//
const QStyleOptionMenuItem__NotCheckable QStyleOptionMenuItem__CheckType = 0

//
const QStyleOptionMenuItem__Exclusive QStyleOptionMenuItem__CheckType = 1

//
const QStyleOptionMenuItem__NonExclusive QStyleOptionMenuItem__CheckType = 2

func (this *QStyleOptionMenuItem) CheckTypeItemName(val int) string {
	switch val {
	case QStyleOptionMenuItem__NotCheckable: // 0
		return "NotCheckable"
	case QStyleOptionMenuItem__Exclusive: // 1
		return "Exclusive"
	case QStyleOptionMenuItem__NonExclusive: // 2
		return "NonExclusive"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionMenuItem_CheckTypeItemName(val int) string {
	var nilthis *QStyleOptionMenuItem
	return nilthis.CheckTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11023() {
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
