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
type QStyleOptionToolBar struct {
	*QStyleOption
}
type QStyleOptionToolBar_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBar_PTR() *QStyleOptionToolBar
}

func (ptr *QStyleOptionToolBar) QStyleOptionToolBar_PTR() *QStyleOptionToolBar { return ptr }

func (this *QStyleOptionToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionToolBarFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBar {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionToolBar{bcthis0}
}
func (*QStyleOptionToolBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBar {
	return NewQStyleOptionToolBarFromPointer(cthis)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionToolBar__StyleOptionType = int

//
const QStyleOptionToolBar__Type QStyleOptionToolBar__StyleOptionType = 14

func (this *QStyleOptionToolBar) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionToolBar__Type: // 14
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBar_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionToolBar
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionToolBar__StyleOptionVersion = int

// 1
const QStyleOptionToolBar__Version QStyleOptionToolBar__StyleOptionVersion = 1

func (this *QStyleOptionToolBar) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionToolBar__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBar_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionToolBar
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionToolBar__ToolBarPosition = int

//
const QStyleOptionToolBar__Beginning QStyleOptionToolBar__ToolBarPosition = 0

//
const QStyleOptionToolBar__Middle QStyleOptionToolBar__ToolBarPosition = 1

//
const QStyleOptionToolBar__End QStyleOptionToolBar__ToolBarPosition = 2

//
const QStyleOptionToolBar__OnlyOne QStyleOptionToolBar__ToolBarPosition = 3

func (this *QStyleOptionToolBar) ToolBarPositionItemName(val int) string {
	switch val {
	case QStyleOptionToolBar__Beginning: // 0
		return "Beginning"
	case QStyleOptionToolBar__Middle: // 1
		return "Middle"
	case QStyleOptionToolBar__End: // 2
		return "End"
	case QStyleOptionToolBar__OnlyOne: // 3
		return "OnlyOne"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBar_ToolBarPositionItemName(val int) string {
	var nilthis *QStyleOptionToolBar
	return nilthis.ToolBarPositionItemName(val)
}

/*


 */
type QStyleOptionToolBar__ToolBarFeature = int

//
const QStyleOptionToolBar__None QStyleOptionToolBar__ToolBarFeature = 0

//
const QStyleOptionToolBar__Movable QStyleOptionToolBar__ToolBarFeature = 1

func (this *QStyleOptionToolBar) ToolBarFeatureItemName(val int) string {
	switch val {
	case QStyleOptionToolBar__None: // 0
		return "None"
	case QStyleOptionToolBar__Movable: // 1
		return "Movable"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolBar_ToolBarFeatureItemName(val int) string {
	var nilthis *QStyleOptionToolBar
	return nilthis.ToolBarFeatureItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11019() {
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
