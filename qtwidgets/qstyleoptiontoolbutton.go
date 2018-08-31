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
type QStyleOptionToolButton struct {
	*QStyleOptionComplex
}
type QStyleOptionToolButton_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionToolButton_PTR() *QStyleOptionToolButton
}

func (ptr *QStyleOptionToolButton) QStyleOptionToolButton_PTR() *QStyleOptionToolButton { return ptr }

func (this *QStyleOptionToolButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionToolButton) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionToolButtonFromPointer(cthis unsafe.Pointer) *QStyleOptionToolButton {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionToolButton{bcthis0}
}
func (*QStyleOptionToolButton) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolButton {
	return NewQStyleOptionToolButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:579
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionToolButton()

/*

 */
func (*QStyleOptionToolButton) NewForInherit() *QStyleOptionToolButton {
	return NewQStyleOptionToolButton()
}
func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolButton)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:583
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionToolButton(int)

/*

 */
func (*QStyleOptionToolButton) NewForInherit_1(version int) *QStyleOptionToolButton {
	return NewQStyleOptionToolButton_1(version)
}
func NewQStyleOptionToolButton_1(version int) *QStyleOptionToolButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolButton)
	return gothis
}

func DeleteQStyleOptionToolButton(this *QStyleOptionToolButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionToolButton__StyleOptionType = int

//
const QStyleOptionToolButton__Type QStyleOptionToolButton__StyleOptionType = 983043

func (this *QStyleOptionToolButton) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionToolButton__Type: // 983043
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolButton_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionToolButton
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionToolButton__StyleOptionVersion = int

// 1
const QStyleOptionToolButton__Version QStyleOptionToolButton__StyleOptionVersion = 1

func (this *QStyleOptionToolButton) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionToolButton__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolButton_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionToolButton
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionToolButton__ToolButtonFeature = int

//
const QStyleOptionToolButton__None QStyleOptionToolButton__ToolButtonFeature = 0

//
const QStyleOptionToolButton__Arrow QStyleOptionToolButton__ToolButtonFeature = 1

//
const QStyleOptionToolButton__Menu QStyleOptionToolButton__ToolButtonFeature = 4

//
const QStyleOptionToolButton__MenuButtonPopup QStyleOptionToolButton__ToolButtonFeature = 4

//
const QStyleOptionToolButton__PopupDelay QStyleOptionToolButton__ToolButtonFeature = 8

//
const QStyleOptionToolButton__HasMenu QStyleOptionToolButton__ToolButtonFeature = 16

func (this *QStyleOptionToolButton) ToolButtonFeatureItemName(val int) string {
	switch val {
	case QStyleOptionToolButton__None: // 0
		return "None"
	case QStyleOptionToolButton__Arrow: // 1
		return "Arrow"
	case QStyleOptionToolButton__Menu: // 4
		return "Menu,MenuButtonPopup"
		// case QStyleOptionToolButton__MenuButtonPopup: // 4
		// return ""
	case QStyleOptionToolButton__PopupDelay: // 8
		return "PopupDelay"
	case QStyleOptionToolButton__HasMenu: // 16
		return "HasMenu"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionToolButton_ToolButtonFeatureItemName(val int) string {
	var nilthis *QStyleOptionToolButton
	return nilthis.ToolButtonFeatureItemName(val)
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
