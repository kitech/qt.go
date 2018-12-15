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
// extern C begin: 5
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
type QStyleOptionFocusRect struct {
	*QStyleOption
}
type QStyleOptionFocusRect_ITF interface {
	QStyleOption_ITF
	QStyleOptionFocusRect_PTR() *QStyleOptionFocusRect
}

func (ptr *QStyleOptionFocusRect) QStyleOptionFocusRect_PTR() *QStyleOptionFocusRect { return ptr }

func (this *QStyleOptionFocusRect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionFocusRect) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionFocusRectFromPointer(cthis unsafe.Pointer) *QStyleOptionFocusRect {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFocusRect{bcthis0}
}
func (*QStyleOptionFocusRect) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionFocusRect {
	return NewQStyleOptionFocusRectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionFocusRect()

/*

 */
func (*QStyleOptionFocusRect) NewForInherit() *QStyleOptionFocusRect {
	return NewQStyleOptionFocusRect()
}
func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFocusRect)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:112
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionFocusRect(int)

/*

 */
func (*QStyleOptionFocusRect) NewForInherit1(version int) *QStyleOptionFocusRect {
	return NewQStyleOptionFocusRect1(version)
}
func NewQStyleOptionFocusRect1(version int) *QStyleOptionFocusRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFocusRect)
	return gothis
}

func DeleteQStyleOptionFocusRect(this *QStyleOptionFocusRect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionFocusRect__StyleOptionType = int

//
const QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = 1

func (this *QStyleOptionFocusRect) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionFocusRect__Type: // 1
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionFocusRect_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionFocusRect
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionFocusRect__StyleOptionVersion = int

// 1
const QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = 1

func (this *QStyleOptionFocusRect) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionFocusRect__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionFocusRect_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionFocusRect
	return nilthis.StyleOptionVersionItemName(val)
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
