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
type QStyleOptionGroupBox struct {
	*QStyleOptionComplex
}
type QStyleOptionGroupBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionGroupBox_PTR() *QStyleOptionGroupBox
}

func (ptr *QStyleOptionGroupBox) QStyleOptionGroupBox_PTR() *QStyleOptionGroupBox { return ptr }

func (this *QStyleOptionGroupBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionGroupBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionGroupBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionGroupBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionGroupBox{bcthis0}
}
func (*QStyleOptionGroupBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:639
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionGroupBox()

/*

 */
func (*QStyleOptionGroupBox) NewForInherit() *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBox()
}
func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGroupBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:642
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionGroupBox(int)

/*

 */
func (*QStyleOptionGroupBox) NewForInherit1(version int) *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBox1(version)
}
func NewQStyleOptionGroupBox1(version int) *QStyleOptionGroupBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGroupBox)
	return gothis
}

func DeleteQStyleOptionGroupBox(this *QStyleOptionGroupBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionGroupBox__StyleOptionType = int

//
const QStyleOptionGroupBox__Type QStyleOptionGroupBox__StyleOptionType = 983046

func (this *QStyleOptionGroupBox) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionGroupBox__Type: // 983046
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionGroupBox_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionGroupBox
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionGroupBox__StyleOptionVersion = int

// 1
const QStyleOptionGroupBox__Version QStyleOptionGroupBox__StyleOptionVersion = 1

func (this *QStyleOptionGroupBox) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionGroupBox__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionGroupBox_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionGroupBox
	return nilthis.StyleOptionVersionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11045() {
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
