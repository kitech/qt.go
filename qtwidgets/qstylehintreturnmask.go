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
type QStyleHintReturnMask struct {
	*QStyleHintReturn
}
type QStyleHintReturnMask_ITF interface {
	QStyleHintReturn_ITF
	QStyleHintReturnMask_PTR() *QStyleHintReturnMask
}

func (ptr *QStyleHintReturnMask) QStyleHintReturnMask_PTR() *QStyleHintReturnMask { return ptr }

func (this *QStyleHintReturnMask) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleHintReturn.GetCthis()
	}
}
func (this *QStyleHintReturnMask) SetCthis(cthis unsafe.Pointer) {
	this.QStyleHintReturn = NewQStyleHintReturnFromPointer(cthis)
}
func NewQStyleHintReturnMaskFromPointer(cthis unsafe.Pointer) *QStyleHintReturnMask {
	bcthis0 := NewQStyleHintReturnFromPointer(cthis)
	return &QStyleHintReturnMask{bcthis0}
}
func (*QStyleHintReturnMask) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturnMask {
	return NewQStyleHintReturnMaskFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:722
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturnMask()

/*

 */
func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnMaskFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturnMask)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:723
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleHintReturnMask()

/*

 */
func DeleteQStyleHintReturnMask(this *QStyleHintReturnMask) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleHintReturnMask__StyleOptionType = int

//
const QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = 61441

func (this *QStyleHintReturnMask) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleHintReturnMask__Type: // 61441
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleHintReturnMask_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleHintReturnMask
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleHintReturnMask__StyleOptionVersion = int

// 1
const QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = 1

func (this *QStyleHintReturnMask) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleHintReturnMask__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleHintReturnMask_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleHintReturnMask
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
