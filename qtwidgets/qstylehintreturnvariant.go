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
type QStyleHintReturnVariant struct {
	*QStyleHintReturn
}
type QStyleHintReturnVariant_ITF interface {
	QStyleHintReturn_ITF
	QStyleHintReturnVariant_PTR() *QStyleHintReturnVariant
}

func (ptr *QStyleHintReturnVariant) QStyleHintReturnVariant_PTR() *QStyleHintReturnVariant { return ptr }

func (this *QStyleHintReturnVariant) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleHintReturn.GetCthis()
	}
}
func (this *QStyleHintReturnVariant) SetCthis(cthis unsafe.Pointer) {
	this.QStyleHintReturn = NewQStyleHintReturnFromPointer(cthis)
}
func NewQStyleHintReturnVariantFromPointer(cthis unsafe.Pointer) *QStyleHintReturnVariant {
	bcthis0 := NewQStyleHintReturnFromPointer(cthis)
	return &QStyleHintReturnVariant{bcthis0}
}
func (*QStyleHintReturnVariant) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturnVariant {
	return NewQStyleHintReturnVariantFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:733
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturnVariant()

/*

 */
func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QStyleHintReturnVariantC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturnVariant)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:734
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleHintReturnVariant()

/*

 */
func DeleteQStyleHintReturnVariant(this *QStyleHintReturnVariant) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QStyleHintReturnVariantD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleHintReturnVariant__StyleOptionType = int

//
const QStyleHintReturnVariant__Type QStyleHintReturnVariant__StyleOptionType = 61442

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleHintReturnVariant__StyleOptionVersion = int

// 1
const QStyleHintReturnVariant__Version QStyleHintReturnVariant__StyleOptionVersion = 1

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
