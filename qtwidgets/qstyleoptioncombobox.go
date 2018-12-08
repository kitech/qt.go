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
type QStyleOptionComboBox struct {
	*QStyleOptionComplex
}
type QStyleOptionComboBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionComboBox_PTR() *QStyleOptionComboBox
}

func (ptr *QStyleOptionComboBox) QStyleOptionComboBox_PTR() *QStyleOptionComboBox { return ptr }

func (this *QStyleOptionComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionComboBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionComboBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionComboBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionComboBox{bcthis0}
}
func (*QStyleOptionComboBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionComboBox {
	return NewQStyleOptionComboBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:601
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionComboBox()

/*

 */
func (*QStyleOptionComboBox) NewForInherit() *QStyleOptionComboBox {
	return NewQStyleOptionComboBox()
}
func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComboBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:605
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionComboBox(int)

/*

 */
func (*QStyleOptionComboBox) NewForInherit1(version int) *QStyleOptionComboBox {
	return NewQStyleOptionComboBox1(version)
}
func NewQStyleOptionComboBox1(version int) *QStyleOptionComboBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComboBox)
	return gothis
}

func DeleteQStyleOptionComboBox(this *QStyleOptionComboBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionComboBox__StyleOptionType = int

//
const QStyleOptionComboBox__Type QStyleOptionComboBox__StyleOptionType = 983044

func (this *QStyleOptionComboBox) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionComboBox__Type: // 983044
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionComboBox_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionComboBox
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionComboBox__StyleOptionVersion = int

// 1
const QStyleOptionComboBox__Version QStyleOptionComboBox__StyleOptionVersion = 1

func (this *QStyleOptionComboBox) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionComboBox__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionComboBox_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionComboBox
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
