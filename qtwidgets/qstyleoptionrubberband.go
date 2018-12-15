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
type QStyleOptionRubberBand struct {
	*QStyleOption
}
type QStyleOptionRubberBand_ITF interface {
	QStyleOption_ITF
	QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand
}

func (ptr *QStyleOptionRubberBand) QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand { return ptr }

func (this *QStyleOptionRubberBand) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionRubberBand) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionRubberBandFromPointer(cthis unsafe.Pointer) *QStyleOptionRubberBand {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionRubberBand{bcthis0}
}
func (*QStyleOptionRubberBand) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBandFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:480
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand()

/*

 */
func (*QStyleOptionRubberBand) NewForInherit() *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBand()
}
func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:484
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand(int)

/*

 */
func (*QStyleOptionRubberBand) NewForInherit1(version int) *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBand1(version)
}
func NewQStyleOptionRubberBand1(version int) *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

func DeleteQStyleOptionRubberBand(this *QStyleOptionRubberBand) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionRubberBand__StyleOptionType = int

//
const QStyleOptionRubberBand__Type QStyleOptionRubberBand__StyleOptionType = 13

func (this *QStyleOptionRubberBand) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionRubberBand__Type: // 13
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionRubberBand_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionRubberBand
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionRubberBand__StyleOptionVersion = int

// 1
const QStyleOptionRubberBand__Version QStyleOptionRubberBand__StyleOptionVersion = 1

func (this *QStyleOptionRubberBand) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionRubberBand__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionRubberBand_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionRubberBand
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
