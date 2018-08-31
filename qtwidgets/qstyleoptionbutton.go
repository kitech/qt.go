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
type QStyleOptionButton struct {
	*QStyleOption
}
type QStyleOptionButton_ITF interface {
	QStyleOption_ITF
	QStyleOptionButton_PTR() *QStyleOptionButton
}

func (ptr *QStyleOptionButton) QStyleOptionButton_PTR() *QStyleOptionButton { return ptr }

func (this *QStyleOptionButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionButton) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionButtonFromPointer(cthis unsafe.Pointer) *QStyleOptionButton {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionButton{bcthis0}
}
func (*QStyleOptionButton) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionButton {
	return NewQStyleOptionButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionButton()

/*

 */
func (*QStyleOptionButton) NewForInherit() *QStyleOptionButton {
	return NewQStyleOptionButton()
}
func NewQStyleOptionButton() *QStyleOptionButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionButtonC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionButton)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:252
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionButton(int)

/*

 */
func (*QStyleOptionButton) NewForInherit_1(version int) *QStyleOptionButton {
	return NewQStyleOptionButton_1(version)
}
func NewQStyleOptionButton_1(version int) *QStyleOptionButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionButtonC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionButton)
	return gothis
}

func DeleteQStyleOptionButton(this *QStyleOptionButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionButton__StyleOptionType = int

//
const QStyleOptionButton__Type QStyleOptionButton__StyleOptionType = 2

func (this *QStyleOptionButton) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionButton__Type: // 2
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionButton_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionButton
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionButton__StyleOptionVersion = int

// 1
const QStyleOptionButton__Version QStyleOptionButton__StyleOptionVersion = 1

func (this *QStyleOptionButton) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionButton__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionButton_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionButton
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionButton__ButtonFeature = int

//
const QStyleOptionButton__None QStyleOptionButton__ButtonFeature = 0

//
const QStyleOptionButton__Flat QStyleOptionButton__ButtonFeature = 1

//
const QStyleOptionButton__HasMenu QStyleOptionButton__ButtonFeature = 2

//
const QStyleOptionButton__DefaultButton QStyleOptionButton__ButtonFeature = 4

//
const QStyleOptionButton__AutoDefaultButton QStyleOptionButton__ButtonFeature = 8

//
const QStyleOptionButton__CommandLinkButton QStyleOptionButton__ButtonFeature = 16

func (this *QStyleOptionButton) ButtonFeatureItemName(val int) string {
	switch val {
	case QStyleOptionButton__None: // 0
		return "None"
	case QStyleOptionButton__Flat: // 1
		return "Flat"
	case QStyleOptionButton__HasMenu: // 2
		return "HasMenu"
	case QStyleOptionButton__DefaultButton: // 4
		return "DefaultButton"
	case QStyleOptionButton__AutoDefaultButton: // 8
		return "AutoDefaultButton"
	case QStyleOptionButton__CommandLinkButton: // 16
		return "CommandLinkButton"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionButton_ButtonFeatureItemName(val int) string {
	var nilthis *QStyleOptionButton
	return nilthis.ButtonFeatureItemName(val)
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
