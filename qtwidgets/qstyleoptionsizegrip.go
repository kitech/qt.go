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
type QStyleOptionSizeGrip struct {
	*QStyleOptionComplex
}
type QStyleOptionSizeGrip_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSizeGrip_PTR() *QStyleOptionSizeGrip
}

func (ptr *QStyleOptionSizeGrip) QStyleOptionSizeGrip_PTR() *QStyleOptionSizeGrip { return ptr }

func (this *QStyleOptionSizeGrip) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSizeGrip) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSizeGripFromPointer(cthis unsafe.Pointer) *QStyleOptionSizeGrip {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSizeGrip{bcthis0}
}
func (*QStyleOptionSizeGrip) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGripFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:642
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionSizeGrip()

/*

 */
func (*QStyleOptionSizeGrip) NewForInherit() *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGrip()
}
func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSizeGrip)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:645
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionSizeGrip(int)

/*

 */
func (*QStyleOptionSizeGrip) NewForInherit1(version int) *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGrip1(version)
}
func NewQStyleOptionSizeGrip1(version int) *QStyleOptionSizeGrip {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSizeGrip)
	return gothis
}

func DeleteQStyleOptionSizeGrip(this *QStyleOptionSizeGrip) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionSizeGrip__StyleOptionType = int

//
const QStyleOptionSizeGrip__Type QStyleOptionSizeGrip__StyleOptionType = 983047

func (this *QStyleOptionSizeGrip) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionSizeGrip__Type: // 983047
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSizeGrip_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionSizeGrip
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionSizeGrip__StyleOptionVersion = int

// 1
const QStyleOptionSizeGrip__Version QStyleOptionSizeGrip__StyleOptionVersion = 1

func (this *QStyleOptionSizeGrip) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionSizeGrip__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSizeGrip_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionSizeGrip
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
