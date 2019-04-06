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
// extern C begin: 3
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
type QStyleHintReturn struct {
	*qtrt.CObject
}
type QStyleHintReturn_ITF interface {
	QStyleHintReturn_PTR() *QStyleHintReturn
}

func (ptr *QStyleHintReturn) QStyleHintReturn_PTR() *QStyleHintReturn { return ptr }

func (this *QStyleHintReturn) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleHintReturn) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStyleHintReturnFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return &QStyleHintReturn{&qtrt.CObject{cthis}}
}
func (*QStyleHintReturn) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return NewQStyleHintReturnFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturn(int, int)

/*

 */
func (*QStyleHintReturn) NewForInherit(version int, type_ int) *QStyleHintReturn {
	return NewQStyleHintReturn(version, type_)
}
func NewQStyleHintReturn(version int, type_ int) *QStyleHintReturn {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturn)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturn(int, int)

/*

 */
func (*QStyleHintReturn) NewForInheritp() *QStyleHintReturn {
	return NewQStyleHintReturnp()
}
func NewQStyleHintReturnp() *QStyleHintReturn {
	// arg: 0, int=Int, =Invalid, , Invalid
	version := 0 /*QStyleOption::Version*/
	// arg: 1, int=Int, =Invalid, , Invalid
	type_ := QStyleHintReturn__SH_Default
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturn)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturn(int, int)

/*

 */
func (*QStyleHintReturn) NewForInheritp1(version int) *QStyleHintReturn {
	return NewQStyleHintReturnp1(version)
}
func NewQStyleHintReturnp1(version int) *QStyleHintReturn {
	// arg: 1, int=Int, =Invalid, , Invalid
	type_ := QStyleHintReturn__SH_Default
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturn)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:711
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleHintReturn()

/*

 */
func DeleteQStyleHintReturn(this *QStyleHintReturn) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QStyleHintReturn__HintReturnType = int

//
const QStyleHintReturn__SH_Default QStyleHintReturn__HintReturnType = 61440

//
const QStyleHintReturn__SH_Mask QStyleHintReturn__HintReturnType = 61441

//
const QStyleHintReturn__SH_Variant QStyleHintReturn__HintReturnType = 61442

func (this *QStyleHintReturn) HintReturnTypeItemName(val int) string {
	switch val {
	case QStyleHintReturn__SH_Default: // 61440
		return "SH_Default"
	case QStyleHintReturn__SH_Mask: // 61441
		return "SH_Mask"
	case QStyleHintReturn__SH_Variant: // 61442
		return "SH_Variant"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleHintReturn_HintReturnTypeItemName(val int) string {
	var nilthis *QStyleHintReturn
	return nilthis.HintReturnTypeItemName(val)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleHintReturn__StyleOptionType = int

//
const QStyleHintReturn__Type QStyleHintReturn__StyleOptionType = 61440

func (this *QStyleHintReturn) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleHintReturn__Type: // 61440
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleHintReturn_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleHintReturn
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleHintReturn__StyleOptionVersion = int

// 1
const QStyleHintReturn__Version QStyleHintReturn__StyleOptionVersion = 1

func (this *QStyleHintReturn) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleHintReturn__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleHintReturn_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleHintReturn
	return nilthis.StyleOptionVersionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11051() {
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
